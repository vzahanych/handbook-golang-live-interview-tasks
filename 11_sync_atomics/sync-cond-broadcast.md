# sync cond broadcast

## Live interview task
Coordinate goroutines with `sync.Cond` — wait until condition true, then broadcast.

## Concepts covered
- sync.Cond
- spurious wakeup
- mutex pairing

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var mu sync.Mutex
    cond := sync.NewCond(&mu)
    ready := false

    go func() {
        mu.Lock()
        for !ready {
            cond.Wait() // releases mu, waits, re-acquires mu on wake
        }
        fmt.Println("go")
        mu.Unlock()
    }()

    mu.Lock()
    ready = true
    cond.Broadcast()
    mu.Unlock()
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- **Always** `for !condition { cond.Wait() }` — not `if` — spurious wakeups and races.
- `Wait` must be called with `mu` held — releases lock while sleeping.
- `Signal` wakes one waiter; `Broadcast` wakes all — choose based on work distribution.
- Prefer channels/context for new code — Cond for legacy patterns (queues, barriers).

## Q&A

**Q: Why loop not if?**  
A: Multiple waiters, condition may change before you run — re-check after wake.

**Q: vs channel?**  
A: Cond for complex condition on shared state under mutex; channel for events.

**Q: Missed signal?**  
A: Set `ready` before `Broadcast` while holding mutex — waiter checks condition.

**Q: Deadlock?**  
A: Wait without holding mutex — panic; Signal without mutex — race.

**Q: Complexity?**  
A: O(waiters) for Broadcast wakeups.
