# pipeline with context cancellation

## Live interview task
Cancel a pipeline early so upstream goroutines exit and don't leak.

## Concepts covered
- context
- pipeline cancellation
- goroutine leak prevention

## Candidate solution

```go
package main

import (
    "context"
    "fmt"
)

func gen(ctx context.Context) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for i := 0; ; i++ {
            select {
            case <-ctx.Done():
                return
            case out <- i:
            }
        }
    }()
    return out
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    for v := range gen(ctx) {
        fmt.Println(v)
        if v == 3 {
            cancel()
            break
        }
    }
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `select` with `ctx.Done()` on **send** — unblocks blocked producer when consumer stops.
- Consumer breaks `for range` but producer may block on send without ctx — always wire cancel into stages.
- `defer cancel()` in main ensures cleanup even without explicit break.
- Each pipeline stage should accept `context.Context` and exit on `Done()`.

## Q&A

**Q: Goroutine leak symptom?**  
A: Growing goroutine count — producer blocked on send to nobody reading.

**Q: Buffered channel help?**  
A: Delays leak — doesn't fix missing cancel.

**Q: `errgroup` + pipeline?**  
A: Cancel group on first error; all stages share `ctx`.

**Q: Close vs cancel?**  
A: Cancel signals stop; close channel signals end of data — often both.

**Q: Complexity?**  
A: O(stages) goroutines; cancel O(1) propagation.
