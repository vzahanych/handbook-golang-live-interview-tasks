# errgroup like cancel on error

## Live interview task
Implement a small errgroup-like helper using context cancellation.

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
        fn := fn
        wg.Add(1)
        go func(){ defer wg.Done(); if err := fn(ctx); err != nil { cancel(); errCh <- err } }()
    }
    wg.Wait(); close(errCh)
    for err := range errCh { return err }
    return nil
}

func main() { fmt.Println(run(context.Background(), func(context.Context) error { return errors.New("fail") })) }
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- None specific; discuss edge cases and complexity.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
