# parallel processing with context

## Live interview task
Map items in parallel but stop when context is canceled.

## Concepts covered
- context cancellation
- parallel map
- cooperative exit

## Candidate solution

```go
package main

import (
    "context"
    "fmt"
    "sync"
)

func Process(ctx context.Context, in []int, workers int, f func(int) int) ([]int, error) {
    if len(in) == 0 {
        return nil, ctx.Err()
    }
    out := make([]int, len(in))
    jobs := make(chan int)
    var wg sync.WaitGroup

    for w := 0; w < workers; w++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for i := range jobs {
                select {
                case <-ctx.Done():
                    return
                default:
                    out[i] = f(in[i])
                }
            }
        }()
    }

    for i := range in {
        select {
        case <-ctx.Done():
            close(jobs)
            wg.Wait()
            return nil, ctx.Err()
        case jobs <- i:
        }
    }
    close(jobs)
    wg.Wait()
    return out, ctx.Err()
}

func main() {
    out, err := Process(context.Background(), []int{1, 2, 3}, 2, func(x int) int { return x * x })
    fmt.Println(out, err)
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Producer checks `ctx.Done()` when enqueueing — stops feeding on cancel.
- Workers check ctx in loop — exit without processing remaining jobs.
- Partial `out` may be incomplete on cancel — document or zero / return error only.
- Use `errgroup` with shared ctx for production.

## Q&A

**Q: Jobs left in channel on cancel?**  
A: Workers may exit early; acceptable if returning error — or drain with cancel.

**Q: `f` respects ctx?**  
A: Pass ctx into `f` if it does I/O: `func(context.Context, int) (int, error)`.

**Q: Complexity?**  
A: O(n) if completes; O(1) exit latency on cancel after in-flight work.

**Q: vs `context.WithCancel` parent?**  
A: Child inherits deadline/cancel from HTTP request.

**Q: Test cancel?**  
A: Slow `f` + early `cancel()` — assert `ctx.Err() != nil`.
