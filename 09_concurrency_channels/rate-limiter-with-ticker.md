# rate limiter with ticker

## Live interview task
Throttle work using `time.Ticker` — fixed interval between operations.

## Concepts covered
- time.Ticker
- rate limiting
- defer Stop

## Candidate solution

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(100 * time.Millisecond)
    defer ticker.Stop()

    for i := 0; i < 3; i++ {
        <-ticker.C
        fmt.Println("request", i)
    }
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Always `defer ticker.Stop()` — leaks goroutine if not stopped.
- Ticker sends on `C` at fixed intervals — first tick after first duration (not immediate unless drain).
- `time.Sleep` between requests drifts under variable work — ticker maintains pace.
- Token bucket / `golang.org/x/time/rate` for burst + average rate — interview follow-up.

## Q&A

**Q: First request immediate?**  
A: Do work first, then wait on ticker — or use `time.After` once.

**Q: vs `time.Tick`?**  
A: `Tick` cannot stop — leaks; use `NewTicker` + `Stop`.

**Q: Context cancel?**  
A: `select { case <-ticker.C: ... case <-ctx.Done(): return }`.

**Q: Distributed rate limit?**  
A: Redis sliding window — out of scope for stdlib demo.

**Q: Complexity?**  
A: O(1) per tick wait.
