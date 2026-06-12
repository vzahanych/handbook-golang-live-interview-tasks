# parallel histogram by buckets

## Live interview task
Build a histogram in parallel using worker-local buckets.

## Concepts covered
- histogram
- parallel aggregation

## Candidate solution

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
)

func Histogram(a []int, bucketCount int) []int {
    workers := runtime.GOMAXPROCS(0); if workers > len(a) { workers = len(a) }
    local := make([][]int, workers)
    var wg sync.WaitGroup
    for w:=0; w<workers; w++ { lo, hi := w*len(a)/workers, (w+1)*len(a)/workers; local[w] = make([]int, bucketCount); wg.Add(1); go func(w,lo,hi int){ defer wg.Done(); for _, v := range a[lo:hi] { if v >= 0 && v < bucketCount { local[w][v]++ } } }(w,lo,hi) }
    wg.Wait()
    out := make([]int, bucketCount)
    for _, part := range local { for i, v := range part { out[i] += v } }
    return out
}

func main() { fmt.Println(Histogram([]int{1,1,2,3,3,3}, 5)) }
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
