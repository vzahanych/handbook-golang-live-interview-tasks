# parallel prefix sum two pass

## Live interview task
Compute prefix sums in parallel with local prefix pass plus offsets.

## Concepts covered
- prefix sum
- parallel scan
- two-pass algorithm

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

func PrefixSum(a []int, workers int) []int {
    out := make([]int, len(a)); sums := make([]int, workers)
    var wg sync.WaitGroup
    for w:=0; w<workers; w++ { lo, hi := w*len(a)/workers, (w+1)*len(a)/workers; wg.Add(1); go func(w,lo,hi int){ defer wg.Done(); sum := 0; for i:=lo; i<hi; i++ { sum += a[i]; out[i] = sum }; sums[w] = sum }(w,lo,hi) }
    wg.Wait()
    offsets := make([]int, workers); for i:=1; i<workers; i++ { offsets[i] = offsets[i-1] + sums[i-1] }
    for w:=1; w<workers; w++ { lo, hi := w*len(a)/workers, (w+1)*len(a)/workers; off := offsets[w]; for i:=lo; i<hi; i++ { out[i] += off } }
    return out
}

func main() { fmt.Println(PrefixSum([]int{1,2,3,4,5}, 2)) }
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
