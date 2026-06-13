# cancellable worker loop

## Live interview task
Stop a worker loop when context is canceled.

## Concepts covered
- context cancellation
- select
- cooperative shutdown

## Candidate solution

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context) {
    ticker := time.NewTicker(50 * time.Millisecond)
    defer ticker.Stop()

    for {
        select {
        case <-ctx.Done():
            fmt.Println("stop:", ctx.Err())
            return
        case <-ticker.C:
            fmt.Println("work")
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    go worker(ctx)
    time.Sleep(120 * time.Millisecond)
    cancel()
    time.Sleep(20 * time.Millisecond)
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Check `ctx.Done()` in **every** blocking loop — not just once at start.
- `default` + sleep spin burns CPU — prefer `ticker` or blocking op with `select`.
- `ctx.Err()` returns `context.Canceled` or `context.DeadlineExceeded`.
- Pass `ctx` as **first parameter** — convention: `func Do(ctx context.Context, ...)`.

## Q&A

**Q: Non-cancelable blocking call?**  
A: Run in goroutine + select on `ctx.Done()` — or use APIs with context (HTTP, DB).

**Q: Cleanup on cancel?**  
A: `defer` in worker after `<-ctx.Done()` or use `context.AfterFunc` (Go 1.21+).

**Q: Parent cancel?**  
A: Child contexts inherit cancel — cancel parent cancels all descendants.

**Q: `context.Background()`?**  
A: Root context — never canceled, no values, for main/tests.

**Q: Complexity?**  
A: O(1) per cancel check.
