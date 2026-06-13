# atomic counter

## Live interview task
Use typed atomics for a concurrent counter.

## Concepts covered
- sync/atomic
- atomic.Int64
- lock-free counter

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    var n atomic.Int64
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            n.Add(1)
        }()
    }
    wg.Wait()
    fmt.Println(n.Load()) // 1000
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Go 1.19+ typed atomics (`atomic.Int64`) — prefer over `atomic.AddInt64(&n, 1)`.
- Atomics for **single word** ops — no mutex for simple counters/metrics.
- `CompareAndSwap`, `Swap`, `Load`, `Store` — building blocks for lock-free structures.
- Not a substitute for mutex when updating multiple related fields.

## Q&A

**Q: Memory ordering?**  
A: Atomics provide happens-before with other atomics — don't mix unsynchronized non-atomic reads.

**Q: `Add` return value?**  
A: New value after add — useful for unique IDs.

**Q: Float atomics?**  
A: `atomic.Uint64` + `math.Float64bits` encoding — or mutex.

**Q: vs mutex counter?**  
A: Atomics faster under contention for single int; mutex simpler for compound ops.

**Q: Race detector?**  
A: Atomics suppress races on that variable — correct use is race-free.
