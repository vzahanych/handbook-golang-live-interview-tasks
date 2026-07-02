package main

// 1. Refactor to a bounded worker pool with a fixed number of workers.
//    Use a fixed number of goroutines reading transactions from a channel; this
//    bounds goroutine count and memory under large N.
// 2. Add backpressure: use a bounded channel for the queue. Decide on behavior
//    when the queue is full:
//    - block (slow producer),
//    - drop (with a counter),
//    - or fail-fast (return 429).
//    Implement at least two behaviors behind ?mode=block|drop|fail_fast.
//    Block is best for internal batch APIs where you must process everything;
//    drop is best for best-effort telemetry; fail-fast is best for synchronous
//    user APIs where you want the caller to retry later.
// 3. Add context cancellation: if the client disconnects or a deadline is hit,
//    stop producing new work and let workers exit without leaks.
//    Use r.Context() with a timeout and have both producer and workers select on
//    ctx.Done().
// 4. Collect per-transaction errors (simulate random failure) without data races,
//    and return a summary JSON response: processed count, failed count, dropped
//    count, and first K errors.
//    Protect counters/slices with a mutex and cap stored errors at K.

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Another common CIL scenario: a single inbound request enqueues a LOT of work
// (e.g. "create policies for 1M devices" or "reconcile transactions"). The
// interviewers typically push you to discuss:
// - bounded concurrency (worker pools),
// - backpressure (what happens when producers outpace consumers),
// - cancellation/shutdown (stop cleanly, no leaks),
// - memory usage (avoid unbounded queues),
// - partial failure handling (collect errors without stopping everything).
//
// This starter exposes /ingest which generates N "transactions" and processes
// them. It currently spawns 1 goroutine per transaction and keeps an unbounded
// in-memory list.
//
// 1. Refactor to a bounded worker pool with a fixed number of workers.
//
// 2. Add backpressure: use a bounded channel for the queue. Decide on behavior
//    when the queue is full:
//    - block (slow producer),
//    - drop (with a counter),
//    - or fail-fast (return 429).
//    Implement at least two behaviors behind a query param like
//    ?mode=block|drop|fail_fast and explain which is best for which API.
//
// 3. Add context cancellation: if the client disconnects or a deadline is hit,
//    stop producing new work and let workers exit without leaks.
//
// 4. Collect per-transaction errors (simulate random failure) without data races,
//    and return a summary JSON response: processed count, failed count, dropped
//    count, and first K errors.
//
// Run hints:
// - Start: `go run .`
// - Example: `curl -sS "localhost:8080/ingest?n=50000&workers=20&queue=1000&mode=drop" | jq`

type Transaction struct {
	ID       int64  `json:"id"`
	Currency string `json:"currency"`
	Amount   int64  `json:"amount"`
}

func process(tx Transaction) error {
	// Simulate work.
	time.Sleep(time.Duration(1+rand.Intn(3)) * time.Millisecond)
	// Random failure.
	if rand.Intn(200) == 0 {
		return fmt.Errorf("tx %d failed", tx.ID)
	}
	return nil
}

type summary struct {
	Processed int      `json:"processed"`
	Failed    int      `json:"failed"`
	Dropped   int      `json:"dropped"`
	Errors    []string `json:"errors"`
}

func ingestAndProcess(ctx context.Context, n int, workers int, queue int, mode string) (summary, error) {
	const maxErrs = 10
	if workers <= 0 {
		workers = 1
	}
	if queue <= 0 {
		queue = 1
	}

	jobs := make(chan Transaction, queue)

	var mu sync.Mutex
	out := summary{Errors: make([]string, 0, maxErrs)}

	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case tx, ok := <-jobs:
					if !ok {
						return
					}
					if err := process(tx); err != nil {
						mu.Lock()
						out.Failed++
						if len(out.Errors) < maxErrs {
							out.Errors = append(out.Errors, err.Error())
						}
						mu.Unlock()
						continue
					}
					mu.Lock()
					out.Processed++
					mu.Unlock()
				}
			}
		}()
	}

	// Producer with backpressure behavior.
	for i := 0; i < n; i++ {
		tx := Transaction{
			ID:       int64(i),
			Currency: []string{"EUR", "USD", "GBP"}[i%3],
			Amount:   int64(100 + i%1000),
		}

		switch mode {
		case "drop":
			select {
			case <-ctx.Done():
				close(jobs)
				wg.Wait()
				return out, ctx.Err()
			case jobs <- tx:
			default:
				mu.Lock()
				out.Dropped++
				mu.Unlock()
			}
		case "fail_fast":
			select {
			case <-ctx.Done():
				close(jobs)
				wg.Wait()
				return out, ctx.Err()
			case jobs <- tx:
			default:
				close(jobs)
				wg.Wait()
				return out, fmt.Errorf("queue full")
			}
		default: // "block"
			select {
			case <-ctx.Done():
				close(jobs)
				wg.Wait()
				return out, ctx.Err()
			case jobs <- tx:
			}
		}
	}

	close(jobs)
	wg.Wait()
	return out, nil
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	http.HandleFunc("/ingest", func(w http.ResponseWriter, r *http.Request) {
		n := 1000
		workers := 20
		queue := 1000
		mode := r.URL.Query().Get("mode")
		if mode == "" {
			mode = "block"
		}

		_, _ = fmt.Sscanf(r.URL.Query().Get("n"), "%d", &n)
		_, _ = fmt.Sscanf(r.URL.Query().Get("workers"), "%d", &workers)
		_, _ = fmt.Sscanf(r.URL.Query().Get("queue"), "%d", &queue)

		// answer 3: request-scoped cancellation + timeout.
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		out, err := ingestAndProcess(ctx, n, workers, queue, mode)
		if err != nil && mode == "fail_fast" && err.Error() == "queue full" {
			http.Error(w, err.Error(), http.StatusTooManyRequests)
			return
		}
		if err != nil && err != context.Canceled && err != context.DeadlineExceeded {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(out)
	})

	srv := &http.Server{Addr: ":8080"} // nil handler => DefaultServeMux
	log.Printf("listening on %s", srv.Addr)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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
//   # block (default) - producer slows down:
//   curl -sS "localhost:8080/ingest?n=50000&workers=20&queue=1000&mode=block" | jq
//
//   # drop - queue stays bounded, but some work is dropped:
//   curl -sS "localhost:8080/ingest?n=50000&workers=20&queue=1000&mode=drop" | jq
//
//   # fail-fast - return 429 when queue is full:
//   curl -sS -i "localhost:8080/ingest?n=50000&workers=1&queue=1&mode=fail_fast"
//
//   go test ./sisu-interview-prep/task-11-workerpool-backpressure-solution -v
