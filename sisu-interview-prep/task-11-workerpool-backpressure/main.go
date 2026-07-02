package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
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

func ingestAndProcess(ctx context.Context, n int) map[string]any {
	// BUG: unbounded goroutines and slice growth.
	errs := make([]string, 0, 10)
	var mu sync.Mutex

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			tx := Transaction{
				ID:       int64(i),
				Currency: []string{"EUR", "USD", "GBP"}[i%3],
				Amount:   int64(100 + i%1000),
			}
			if err := process(tx); err != nil {
				mu.Lock()
				if len(errs) < 10 {
					errs = append(errs, err.Error())
				}
				mu.Unlock()
			}
		}(i)
	}
	wg.Wait()

	return map[string]any{
		"processed": n,
		"errors":    errs,
	}
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

		// TODO: use r.Context() and a per-request timeout.
		_ = workers
		_ = queue
		_ = mode

		out := ingestAndProcess(context.Background(), n)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(out)
	})

	log.Printf("listening on :8080")
	_ = http.ListenAndServe(":8080", nil)
}
