# parallel processing with context

## Live interview task
Map items in parallel but stop early when context is canceled.

## Concepts covered
- context
- parallel map
- cancellation

## Candidate solution

```go
package main

import (
    "context"
    "fmt"
    "sync"
)

func Process(ctx context.Context, in []int, workers int, f func(int) int) ([]int, error) {
    out := make([]int, len(in)); jobs := make(chan int); var wg sync.WaitGroup
    for w:=0; w<workers; w++ { wg.Add(1); go func(){ defer wg.Done(); for i := range jobs { select { case <-ctx.Done(): return; default: out[i] = f(in[i]) } } }() }
    for i := range in { select { case <-ctx.Done(): close(jobs); wg.Wait(); return nil, ctx.Err(); case jobs <- i: } }
    close(jobs); wg.Wait(); return out, ctx.Err()
}

func main() { out, err := Process(context.Background(), []int{1,2,3}, 2, func(x int) int { return x*x }); fmt.Println(out, err) }
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
