package main

// 1. The current implementation retries POST blindly. Explain why this can
//    double-charge in real systems, and implement IDPOTENCY: generate a request ID
//    per payment and send it as `Idempotency-Key` header on every retry to the
//    SAME backend.
//    Blind retries are unsafe because a timeout can happen after the downstream
//    has already committed the side effect (charged the card) but before we got
//    the response, so retrying creates a second charge. The fix is to make each
//    attempt idempotent: generate a stable idempotency key per payment+backend
//    and send it on every retry so the backend can de-duplicate and return the
//    same result for the same key instead of charging twice.
// 2. Add a per-request deadline (context) so one slow backend doesn't hang the
//    entire batch. Each downstream call must be created with
//    http.NewRequestWithContext.
//    Put a deadline on each downstream call (and/or a global deadline for the
//    whole fanout). Use http.NewRequestWithContext(ctx, ...) so cancellation
//    interrupts the request and unblocks goroutines.
// 3. Fix the resource handling: close response bodies on all paths, and return
//    a per-backend result including HTTP status and error. Ensure nil-pointer
//    panics are impossible when a request fails.
//    Always check err before reading resp.Body and always close resp.Body (even
//    on 5xx). Return results in a slice of the same length as backends so the
//    caller gets partial successes plus per-backend errors.
// 4. Bound the concurrency of downstream calls (e.g. max 10 in-flight). Assume
//    `len(backends)` can be hundreds.
//    Use a semaphore channel to cap in-flight requests; this protects both this
//    process (FDs/scheduler) and downstream services.
// 5. Right now results are appended to a shared slice without synchronization.
//    Fix the race and preserve input order (result[i] corresponds to backends[i]).
//    Preallocate `results := make([]payResult, len(backends))` and have each
//    goroutine write exactly one slot by index; no mutex is needed and order is
//    preserved naturally.

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

// In the CIL, a single API call often triggers multiple downstream calls (partners,
// Allianz, internal services). Under load, the common "easy bugs" are:
// - retries that accidentally double-charge (no idempotency),
// - no timeouts (hung requests),
// - leaking response bodies (FD/memory leaks),
// - unbounded concurrency (host/downstreams collapse),
// - shared mutable state without synchronization (races).
//
// You are given a toy "payment" fan-out to multiple backends.
//
// 1. The current implementation retries POST blindly. Explain why this can
//    double-charge in real systems, and implement IDPOTENCY: generate a request ID
//    per payment and send it as `Idempotency-Key` header on every retry to the
//    SAME backend.
//
// 2. Add a per-request deadline (context) so one slow backend doesn't hang the
//    entire batch. Each downstream call must be created with
//    http.NewRequestWithContext.
//
// 3. Fix the resource handling: close response bodies on all paths, and return
//    a per-backend result including HTTP status and error. Ensure nil-pointer
//    panics are impossible when a request fails.
//
// 4. Bound the concurrency of downstream calls (e.g. max 10 in-flight). Assume
//    `len(backends)` can be hundreds.
//
// 5. Right now results are appended to a shared slice without synchronization.
//    Fix the race and preserve input order (result[i] corresponds to backends[i]).
//
// Run hints:
// - Start the fake backends: `go run .`
// - In another terminal: `curl -sS localhost:8080/pay -d '{"amount":123,"backends":20,"fail_rate":0.3}' | jq`

type payRequest struct {
	Amount   int64   `json:"amount"`
	Backends int     `json:"backends"`
	FailRate float64 `json:"fail_rate"`
}

type payResult struct {
	Backend string `json:"backend"`
	Status  int    `json:"status"`
	Charge  int64  `json:"charge"`
	Err     string `json:"err,omitempty"`
}

// doPaymentFanout intentionally contains multiple issues for the interview.
// func doPaymentFanout(amount int64, backends []string, failRate float64) []payResult {
func doPaymentFanout(ctx context.Context, client *http.Client, amount int64, backends []string, failRate float64, perRequestTimeout time.Duration, maxInFlight int, failFast bool) ([]payResult, error) {
	// answer 5: fixed-size slice, one writer per index => no race, preserves order.
	results := make([]payResult, len(backends))

	// answer 4: bound concurrency with a semaphore.
	sem := make(chan struct{}, maxInFlight)
	var wg sync.WaitGroup

	// answer 1: stable request id per payment; backend key derives from it.
	requestID := fmt.Sprintf("%d", time.Now().UnixNano())

	// answer 3/extra: allow fail-fast vs partial results behind a switch.
	errCh := make(chan error, 1)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for i, b := range backends {
		wg.Add(1)
		go func(i int, backendURL string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			// answer 1: stable idempotency key for retries to THIS backend.
			idempotencyKey := requestID + ":" + fmt.Sprintf("%d", i)

			var lastErr error
			for attempt := 0; attempt < 3; attempt++ {
				// answer 2: per-request deadline and cancellation propagation.
				reqCtx, cancelReq := context.WithTimeout(ctx, perRequestTimeout)
				reqBody, _ := json.Marshal(map[string]any{
					"amount":    amount,
					"fail_rate": failRate,
				})
				// resp, err := http.Post(backendURL+"/charge", "application/json", bytes.NewReader(reqBody))
				httpReq, _ := http.NewRequestWithContext(reqCtx, http.MethodPost, backendURL+"/charge", bytes.NewReader(reqBody)) // answer 2
				httpReq.Header.Set("Content-Type", "application/json")
				httpReq.Header.Set("Idempotency-Key", idempotencyKey) // answer 1

				resp, err := client.Do(httpReq)
				cancelReq()
				if err != nil {
					lastErr = err
					continue
				}

				// answer 3: always close bodies.
				// body, readErr := io.ReadAll(resp.Body) // old: buffers whole response
				// _ = resp.Body.Close()
				defer func() { _ = resp.Body.Close() }()

				if resp.StatusCode >= 400 {
					// answer 3: stream a bounded error snippet, not the whole body.
					var errBuf strings.Builder
					if _, readErr := io.Copy(&errBuf, io.LimitReader(resp.Body, 4<<10)); readErr != nil {
						lastErr = readErr
						continue
					}
					lastErr = errors.New(errBuf.String())
					if resp.StatusCode >= 500 {
						continue
					}
					break
				}

				// answer 3: drain success bodies without buffering them.
				if _, readErr := io.Copy(io.Discard, resp.Body); readErr != nil {
					lastErr = readErr
					continue
				}

				results[i] = payResult{
					Backend: backendURL,
					Status:  resp.StatusCode,
					Charge:  amount,
				}
				return
			}

			if lastErr != nil {
				results[i] = payResult{
					Backend: backendURL,
					Status:  0,
					Err:     lastErr.Error(),
				}
				if failFast {
					select {
					case errCh <- lastErr:
						cancel()
					default:
					}
				}
			}
		}(i, b)
	}

	wg.Wait()
	close(errCh)

	if failFast {
		if err, ok := <-errCh; ok && err != nil {
			return results, err
		}
	}
	return results, nil
}

func makeBackendURLs(n int) []string {
	out := make([]string, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, fmt.Sprintf("http://localhost:%d", 9000+i))
	}
	return out
}

func startFakeBackends(n int) {
	for i := 0; i < n; i++ {
		port := 9000 + i
		mux := http.NewServeMux()

		// BUG: shared state without lock; and no idempotency handling
		charges := map[string]int64{}
		var mu sync.Mutex

		mux.HandleFunc("/charge", func(w http.ResponseWriter, r *http.Request) {
			var payload struct {
				Amount   int64   `json:"amount"`
				FailRate float64 `json:"fail_rate"`
			}
			_ = json.NewDecoder(r.Body).Decode(&payload)

			// Random 5xx to trigger retries.
			if rand.Float64() < payload.FailRate {
				http.Error(w, "temporary backend error", http.StatusServiceUnavailable)
				return
			}

			key := r.Header.Get("Idempotency-Key")
			if key == "" {
				// In real systems this is often required for POST to avoid duplicates.
				http.Error(w, "missing Idempotency-Key", http.StatusBadRequest)
				return
			}

			// BUG: map race under concurrent calls.
			mu.Lock()
			// answer 1: idempotency semantics in the fake backend - first write wins.
			if _, ok := charges[key]; !ok {
				charges[key] = payload.Amount
			}
			charged := charges[key]
			mu.Unlock()

			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"charged_total_for_key": charged,
				"port":                  port,
			})
		})

		srv := &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		}

		go func() {
			_ = srv.ListenAndServe()
		}()
	}
}

func payHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var req payRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Backends <= 0 {
		http.Error(w, "backends must be > 0", http.StatusBadRequest)
		return
	}

	failFast := r.URL.Query().Get("fail_fast") == "true"
	backends := makeBackendURLs(req.Backends)

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        200,
			MaxIdleConnsPerHost: 50, // answer 4: avoid connect storms under load
			IdleConnTimeout:     30 * time.Second,
		},
	}

	results, err := doPaymentFanout(r.Context(), client, req.Amount, backends, req.FailRate, 800*time.Millisecond, 10, failFast)

	status := http.StatusOK
	if failFast && err != nil {
		status = http.StatusBadGateway
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"backends": len(backends),
		"failFast": failFast,
		"results":  results,
		"error":    errString(err),
	})
}

func errString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	startFakeBackends(30)

	http.HandleFunc("/pay", payHandler)

	srv := &http.Server{Addr: ":8080"} // nil Handler => DefaultServeMux (prompt rule)
	log.Printf("listening on %s (backends: 9000+)", srv.Addr)

	// answer: simplest graceful shutdown (prompt rule).
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("server error: %v", err)
		}
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_ = srv.Shutdown(shutdownCtx)
}

// How to run and test:
//   go run .
//
//   # "partial results" mode (default). POST only; GET returns 405.
//   curl -sS -X POST "localhost:8080/pay" --data-binary @sisu-interview-prep/task-07-idempotent-retries-solution/sample.json | jq
//
//   # fail-fast mode (returns 502 if any backend fails):
//   curl -sS -X POST "localhost:8080/pay?fail_fast=true" --data-binary @sisu-interview-prep/task-07-idempotent-retries-solution/sample.json | jq
//
//   go test ./sisu-interview-prep/task-07-idempotent-retries-solution -v
