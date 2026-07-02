package main

import (
	"fmt"
	"sync"
)

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

func main() {
	// Sample data (paste from the meeting chat). Grow this to stress it:
	txs := []Transaction{
		{"EUR", 1200}, {"USD", 999}, {"EUR", 350},
		{"GBP", 5000}, {"USD", 1}, {"EUR", 8000},
	}

	totals := map[string]int64{}
	var wg sync.WaitGroup

	for _, tx := range txs {
		go func() {
			totals[tx.Currency] += tx.Amount // concurrent map write ⇒ data race
			wg.Done()                        // no matching wg.Add ⇒ panic
		}()
	}
	wg.Wait()

	for cur, sum := range totals {
		fmt.Printf("%s: %d\n", cur, sum)
	}
}
