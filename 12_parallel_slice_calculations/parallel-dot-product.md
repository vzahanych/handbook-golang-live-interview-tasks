# parallel dot product

## Live interview task
Compute dot product of two equal-length slices in parallel.

## Concepts covered
- parallel math
- slice indexing

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

func Dot(a, b []int, workers int) int {
    if len(a) != len(b) { panic("length mismatch") }
    partial := make([]int, workers)
    var wg sync.WaitGroup
    for w:=0; w<workers; w++ { lo, hi := w*len(a)/workers, (w+1)*len(a)/workers; wg.Add(1); go func(w,lo,hi int){ defer wg.Done(); for i:=lo; i<hi; i++ { partial[w] += a[i] * b[i] } }(w,lo,hi) }
    wg.Wait()
    total := 0; for _, v := range partial { total += v }; return total
}

func main() { fmt.Println(Dot([]int{1,2,3}, []int{4,5,6}, 2)) }
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
