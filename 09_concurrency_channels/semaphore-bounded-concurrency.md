# semaphore bounded concurrency

## Live interview task
Run many tasks concurrently but cap how many execute **at the same time** using a buffered channel as a **counting semaphore** — capacity = max in-flight workers. Example: launch 5 goroutines with `sem := make(chan struct{}, 2)` — at most 2 sleep/work at once; each task **acquires** with `sem <- struct{}{}` before work and **releases** with `<-sem` in `defer` when done.

## Concepts covered
- buffered channels
- semaphore
- bounded concurrency

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    // Buffered channel as counting semaphore: len(sem) = slots in use, cap = max slots.
    // Pre-filled buffer would mean fewer acquires available — start empty (0/2 in use).
    sem := make(chan struct{}, 2)
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            sem <- struct{}{} // acquire: send fills a slot; blocks when 2 workers already hold tokens
            defer func() { <-sem }() // release: receive frees a slot for the next waiter
            time.Sleep(10 * time.Millisecond) // simulated work — only 2 ids run this at once
            fmt.Println("done", id)
        }(i) // pass id — avoid loop variable capture
    }
    wg.Wait() // wait for all 5 tasks (each acquired, worked, released)
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Empty struct `struct{}{}` — zero-size token, count = buffer capacity.
- Acquire: send to sem; release: receive — inverted from mutex mental model.
- `defer` release ensures unlock on panic — pair with worker defer Done.
- `golang.org/x/sync/semaphore` — weighted semaphore for acquire N units.

## Q&A

**Q: vs worker pool?**  
A: Semaphore limits concurrency; pool also queues work — often combined.

**Q: Deadlock?**  
A: Acquire without release — always defer release; don't block forever on full sem without timeout.

**Q: Try-acquire?**  
A: `select { case sem<-struct{}{}: ...; default: ... }`.

**Q: Capacity 0?**  
A: Blocks every acquire until release — synchronous handoff.

**Q: Production?**  
A: Rate-limit DB connections, API calls, file descriptors.
