# parallel filter stable

## Live interview task
Filter a slice in parallel and keep the original relative order.

## Concepts covered
- parallel filter
- stable order
- chunk-local buffers

## Candidate solution

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
)

func ParallelFilter[T any](in []T, keep func(T) bool) []T {
    workers := runtime.GOMAXPROCS(0)
    if workers > len(in) { workers = len(in) }
    chunks := make([][]T, workers)
    var wg sync.WaitGroup
    for w := 0; w < workers; w++ {
        lo, hi := w*len(in)/workers, (w+1)*len(in)/workers
        wg.Add(1)
        go func(w, lo, hi int){ defer wg.Done(); for _, v := range in[lo:hi] { if keep(v) { chunks[w] = append(chunks[w], v) } } }(w, lo, hi)
    }
    wg.Wait()
    var out []T
    for _, c := range chunks { out = append(out, c...) }
    return out
}

func main() { fmt.Println(ParallelFilter([]int{1,2,3,4}, func(x int) bool { return x%2 == 0 })) }
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
