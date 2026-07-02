package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"sync"
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
func doPaymentFanout(amount int64, backends []string, failRate float64) []payResult {
	results := make([]payResult, 0, len(backends)) // race: appended from goroutines

	for _, b := range backends {
		go func(backendURL string) {
			// BUG: no timeout; blind retries; no idempotency; response bodies not closed; data race
			var lastErr error
			for attempt := 0; attempt < 3; attempt++ {
				reqBody, _ := json.Marshal(map[string]any{
					"amount":    amount,
					"fail_rate": failRate,
				})
				resp, err := http.Post(backendURL+"/charge", "application/json", bytes.NewReader(reqBody))
				if err != nil {
					lastErr = err
					continue
				}
				body, _ := io.ReadAll(resp.Body)
				if resp.StatusCode >= 500 {
					lastErr = errors.New(string(body))
					continue
				}
				results = append(results, payResult{
					Backend: backendURL,
					Status:  resp.StatusCode,
					Charge:  amount,
				})
				return
			}

			results = append(results, payResult{
				Backend: backendURL,
				Status:  0,
				Err:     lastErr.Error(),
			})
		}(b)
	}

	// BUG: no wait - returns before goroutines finish
	return results
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
			charges[key] += payload.Amount

			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"charged_total_for_key": charges[key],
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

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	startFakeBackends(30)

	http.HandleFunc("/pay", func(w http.ResponseWriter, r *http.Request) {
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

		backends := makeBackendURLs(req.Backends)

		// BUG: ignores request context; returns empty most of the time (no wait).
		results := doPaymentFanout(req.Amount, backends, req.FailRate)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"backends": len(backends),
			"results":  results,
		})
	})

	srv := &http.Server{
		Addr: ":8080",
	}
	log.Printf("listening on %s (backends: 9000+)", srv.Addr)

	// Keep the program alive long enough to exercise it manually.
	// In an interview, you'd be asked to add graceful shutdown.
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		_ = srv.ListenAndServe()
	}()
	wg.Wait()
}
