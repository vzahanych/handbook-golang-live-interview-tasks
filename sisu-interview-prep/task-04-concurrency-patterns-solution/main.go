package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 1. Run `go run -race .`. There is a data race on the shared `totals` map
//    (concurrent writes). Fix it — sync.Mutex, a sharded set of counters, or
//    another approach — and be ready to explain the trade-off you chose.
//    The race is caused by multiple goroutines writing to the same map without
//    synchronization. The simplest safe fix is a sync.Mutex around the map
//    update: correct and easy to explain, but it serializes all updates on one
//    lock. A more scalable option is sharding (multiple maps/locks by hash of
//    currency) or aggregating through a single goroutine, but the mutex is the
//    clean baseline.
// 2. The WaitGroup is used incorrectly: there is no wg.Add before the goroutines
//    call wg.Done, so this panics with "negative WaitGroup counter" (and, once
//    fixed the wrong way, main can print before the workers finish). Fix the
//    Add / Done / Wait placement.
//    wg.Add must happen before starting worker goroutines (or at least before
//    they can call Done). With a worker pool, we Add(workers), each worker
//    defers Done once, and main calls Wait after closing the jobs channel.
// 3. This spawns one goroutine PER transaction. With millions of transactions
//    that exhausts memory and the scheduler. Refactor to a bounded worker pool:
//    a fixed number of workers reading transactions from a channel.
//    Use a buffered jobs channel and a fixed number of workers. This bounds
//    goroutines to O(workers) while keeping throughput high.
// 4. Add a context so the aggregation can be cancelled mid-batch (an upstream
//    timeout or a shutdown) without leaking the workers.
//    Thread a context through: workers select on ctx.Done() as well as jobs,
//    and the producer stops sending when cancelled. Closing the jobs channel
//    lets workers exit cleanly.

// A back-end aggregates incoming payment transactions by currency. This starter
// "works" for a handful of transactions but breaks under load — it is the code
// behind the "very large number of simultaneous transactions" discussion the
// briefing says accompanies every session.
//
// 1. Run `go run -race .`. There is a data race on the shared `totals` map
//    (concurrent writes). Fix it — sync.Mutex, a sharded set of counters, or
//    another approach — and be ready to explain the trade-off you chose.
//
// 2. The WaitGroup is used incorrectly: there is no wg.Add before the goroutines
//    call wg.Done, so this panics with "negative WaitGroup counter" (and, once
//    fixed the wrong way, main can print before the workers finish). Fix the
//    Add / Done / Wait placement.
//
// 3. This spawns one goroutine PER transaction. With millions of transactions
//    that exhausts memory and the scheduler. Refactor to a bounded worker pool:
//    a fixed number of workers reading transactions from a channel.
//
// 4. Add a context so the aggregation can be cancelled mid-batch (an upstream
//    timeout or a shutdown) without leaking the workers.

type Transaction struct {
	Currency string
	Amount   int64 // minor units (e.g. cents)
}

func aggregate(ctx context.Context, txs []Transaction, workers int, queue int) (map[string]int64, error) {
	if workers <= 0 {
		workers = 1
	}
	if queue <= 0 {
		queue = 1
	}

	totals := map[string]int64{}
	var mu sync.Mutex // answer 1: protect totals map

	jobs := make(chan Transaction, queue) // answer 3: bounded queue
	var wg sync.WaitGroup

	wg.Add(workers) // answer 2: Add before goroutines can call Done
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done(): // answer 4: cancellation
					return
				case tx, ok := <-jobs:
					if !ok {
						return
					}
					mu.Lock()
					totals[tx.Currency] += tx.Amount
					mu.Unlock()
				}
			}
		}()
	}

	for _, tx := range txs {
		select {
		case <-ctx.Done(): // answer 4: stop producing on cancel
			close(jobs)
			wg.Wait()
			return totals, ctx.Err()
		case jobs <- tx:
		}
	}
	close(jobs)
	wg.Wait()
	return totals, nil
}

func main() {
	// Sample data (paste from the meeting chat). Grow this to stress it:
	txs := []Transaction{
		{"EUR", 1200}, {"USD", 999}, {"EUR", 350},
		{"GBP", 5000}, {"USD", 1}, {"EUR", 8000},
	}

	// totals := map[string]int64{} // old: unsafely shared map across goroutines
	// var wg sync.WaitGroup        // old: wg.Add missing before wg.Done

	// answer 4: cancellation example (deadline).
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	totals, err := aggregate(ctx, txs, 4, 16)
	if err != nil {
		fmt.Printf("aggregation error: %v\n", err)
	}

	for cur, sum := range totals {
		fmt.Printf("%s: %d\n", cur, sum)
	}
}

// How to run and test:
//   go run .
//   go run -race .
//
//   go test ./sisu-interview-prep/task-04-concurrency-patterns-solution -v
//   go test ./sisu-interview-prep/task-04-concurrency-patterns-solution -bench=. -benchmem
