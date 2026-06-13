# parallel dot product

## Live interview task
Compute dot product of two equal-length slices in parallel.

## Concepts covered
- parallel reduction
- paired slice access

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

func Dot(a, b []int, workers int) int {
    if len(a) != len(b) {
        panic("length mismatch")
    }
    if len(a) == 0 {
        return 0
    }
    if workers > len(a) {
        workers = len(a)
    }

    partial := make([]int, workers)
    var wg sync.WaitGroup
    for w := 0; w < workers; w++ {
        lo := w * len(a) / workers
        hi := (w + 1) * len(a) / workers
        wg.Add(1)
        go func(w, lo, hi int) {
            defer wg.Done()
            for i := lo; i < hi; i++ {
                partial[w] += a[i] * b[i]
            }
        }(w, lo, hi)
    }
    wg.Wait()

    total := 0
    for _, v := range partial {
        total += v
    }
    return total
}

func main() {
    fmt.Println(Dot([]int{1, 2, 3}, []int{4, 5, 6}, 2)) // 32
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Same chunking pattern as parallel sum — product per index then sum partials.
- Validate equal lengths up front — return error instead of panic in libraries.
- SIMD/BLAS beats goroutines for numeric heavy — mention when asked.

## Q&A

**Q: Result for example?**  
A: 1*4 + 2*5 + 3*6 = 32.

**Q: Complexity?**  
A: O(n) multiplies and adds.

**Q: Float vectors?**  
A: Same pattern with `float64` and Kahan summation if precision matters.

**Q: Matrix multiply?**  
A: Row-column dot products — extend pattern.

**Q: Edge cases?**  
A: Mismatched len — error; len 0 → 0.
