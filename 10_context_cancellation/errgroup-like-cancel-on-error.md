# errgroup like cancel on error

## Live interview task
Implement a small errgroup-style helper: run tasks in parallel, cancel siblings on first error.

## Concepts covered
- context cancellation
- WaitGroup
- error propagation

## Candidate solution

```go
package main

import (
    "context"
    "errors"
    "fmt"
    "sync"
)

func run(ctx context.Context, fns ...func(context.Context) error) error {
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    errCh := make(chan error, len(fns))
    var wg sync.WaitGroup

    for _, fn := range fns {
        wg.Add(1)
        go func(f func(context.Context) error) {
            defer wg.Done()
            if err := f(ctx); err != nil {
                cancel()
                errCh <- err
            }
        }(fn)
    }

    wg.Wait()
    close(errCh)

    for err := range errCh {
        return err // first collected error
    }
    return ctx.Err()
}

func main() {
    err := run(context.Background(),
        func(ctx context.Context) error {
            <-ctx.Done()
            return ctx.Err()
        },
        func(ctx context.Context) error {
            return errors.New("fail")
        },
    )
    fmt.Println(err)
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `cancel()` on first error — other goroutines should respect `ctx.Done()`.
- Buffered `errCh` — sender doesn't block if main already moved on (size = task count).
- Production: `golang.org/x/sync/errgroup` — `Group`, `Go`, `Wait`, optional `SetLimit`.
- Return first error only — join multiple with `errors.Join` if needed.

## Q&A

**Q: Wait for all after error?**  
A: Yes — `Wait` drains goroutines that exit on canceled ctx; avoid leak.

**Q: Limit concurrency?**  
A: `errgroup` with semaphore or worker pool channel.

**Q: Panic in worker?**  
A: errgroup doesn't recover — wrap with safeGo or defer recover.

**Q: vs WaitGroup alone?**  
A: WaitGroup no error propagation or cancel.

**Q: Complexity?**  
A: O(n) goroutines for n tasks.
