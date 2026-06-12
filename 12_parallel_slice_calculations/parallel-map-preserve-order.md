# parallel map preserve order

## Live interview task
Apply a function to each slice element in parallel while preserving output order.

## Concepts covered
- parallel map
- order preservation
- generics

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

func ParallelMap[A, B any](in []A, workers int, f func(A) B) []B {
    out := make([]B, len(in))
    jobs := make(chan int)
    var wg sync.WaitGroup
    for w := 0; w < workers; w++ {
        wg.Add(1)
        go func(){ defer wg.Done(); for i := range jobs { out[i] = f(in[i]) } }()
    }
    for i := range in { jobs <- i }
    close(jobs); wg.Wait()
    return out
}

func main() { fmt.Println(ParallelMap([]int{1,2,3}, 2, func(x int) int { return x*x })) }
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
