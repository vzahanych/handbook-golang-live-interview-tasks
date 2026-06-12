# parallel sum int slice

## Live interview task
Split a large int slice into chunks and sum it in parallel.

## Concepts covered
- parallel reduction
- chunking
- false sharing awareness

## Candidate solution

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
)

func ParallelSum(a []int) int {
    if len(a) == 0 { return 0 }
    workers := runtime.GOMAXPROCS(0)
    if workers > len(a) { workers = len(a) }
    partial := make([]int, workers)
    var wg sync.WaitGroup
    for w := 0; w < workers; w++ {
        lo := w * len(a) / workers
        hi := (w+1) * len(a) / workers
        wg.Add(1)
        go func(w, lo, hi int) { defer wg.Done(); for _, v := range a[lo:hi] { partial[w] += v } }(w, lo, hi)
    }
    wg.Wait()
    total := 0
    for _, v := range partial { total += v }
    return total
}

func main() { fmt.Println(ParallelSum([]int{1,2,3,4,5,6})) }
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
