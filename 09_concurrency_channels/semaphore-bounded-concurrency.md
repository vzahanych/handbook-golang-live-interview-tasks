# semaphore bounded concurrency

## Live interview task
Limit concurrent work with a buffered channel used as a semaphore.

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
    sem := make(chan struct{}, 2) // max 2 concurrent
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            sem <- struct{}{}        // acquire
            defer func() { <-sem }() // release
            time.Sleep(10 * time.Millisecond)
            fmt.Println("done", id)
        }(i)
    }
    wg.Wait()
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
