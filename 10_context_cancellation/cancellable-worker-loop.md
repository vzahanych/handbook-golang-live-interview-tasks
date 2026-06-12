# cancellable worker loop

## Live interview task
Stop a worker loop when context is canceled.

## Concepts covered
- context cancellation
- select

## Candidate solution

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("stop:", ctx.Err())
            return
        default:
            fmt.Println("work")
            time.Sleep(50 * time.Millisecond)
        }
    }
}

func main() { ctx, cancel := context.WithCancel(context.Background()); go worker(ctx); time.Sleep(120*time.Millisecond); cancel(); time.Sleep(20*time.Millisecond) }
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
