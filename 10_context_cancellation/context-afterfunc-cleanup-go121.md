# context afterfunc cleanup go121

## Live interview task
Register cleanup to run when a context is canceled using `context.AfterFunc` (Go 1.21+).

## Concepts covered
- context.AfterFunc
- cancel cleanup
- goroutine lifecycle

## Candidate solution

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())

    stop := context.AfterFunc(ctx, func() {
        fmt.Println("cleanup: resources released")
    })

    go func() {
        time.Sleep(50 * time.Millisecond)
        cancel()
    }()

    time.Sleep(100 * time.Millisecond)
    _ = stop // call stop() to prevent AfterFunc from running if cancel not fired
}
```

## Run

```bash
go run . # Go 1.21+
```

## Interview notes / pitfalls
- `AfterFunc` runs in **separate goroutine** when ctx canceled — not synchronous with `cancel()`.
- `stop()` prevents callback if called before cancel — returns whether func already started.
- Alternative: `defer cleanup()` in scope that owns cancel — simpler for linear code.
- Don't rely on AfterFunc order across multiple registrations — not guaranteed.

## Q&A

**Q: vs defer cancel()?**  
A: `defer cancel()` releases timer; AfterFunc is for custom cleanup on **any** cancel path.

**Q: Panic in AfterFunc?**  
A: Crashes that goroutine — keep cleanup minimal.

**Q: Parent canceled?**  
A: Child AfterFunc runs when child ctx done — includes parent cancel.

**Q: Use case?**  
A: Close idle connections, stop background poller tied to request ctx.

**Q: Production?**  
A: Often explicit `defer` in handler; AfterFunc for library hooks.
