# parallel prefix sum two pass

## Live interview task
Compute prefix sums in parallel: local inclusive scan per chunk, then add chunk offsets.

## Concepts covered
- prefix sum (scan)
- two-pass parallel algorithm

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

func PrefixSum(a []int, workers int) []int {
    if len(a) == 0 || workers <= 0 {
        return nil
    }
    if workers > len(a) {
        workers = len(a)
    }

    out := make([]int, len(a))
    sums := make([]int, workers)
    var wg sync.WaitGroup

    // Pass 1: local inclusive prefix per chunk + chunk total
    for w := 0; w < workers; w++ {
        lo := w * len(a) / workers
        hi := (w + 1) * len(a) / workers
        wg.Add(1)
        go func(w, lo, hi int) {
            defer wg.Done()
            sum := 0
            for i := lo; i < hi; i++ {
                sum += a[i]
                out[i] = sum
            }
            sums[w] = sum
        }(w, lo, hi)
    }
    wg.Wait()

    // Pass 2: offsets from previous chunk totals
    offsets := make([]int, workers)
    for i := 1; i < workers; i++ {
        offsets[i] = offsets[i-1] + sums[i-1]
    }
    for w := 1; w < workers; w++ {
        lo := w * len(a) / workers
        hi := (w + 1) * len(a) / workers
        off := offsets[w]
        for i := lo; i < hi; i++ {
            out[i] += off
        }
    }
    return out
}

func main() {
    fmt.Println(PrefixSum([]int{1, 2, 3, 4, 5}, 2))
}
```

## Run

```bash
go run .
```

## Expected output

```
[1 3 6 10 15]
```

## Interview notes / pitfalls
- Pass 1: inclusive scan within each chunk; store chunk sum in `sums[w]`.
- Pass 2: add prefix of prior chunk sums to each element — global inclusive prefix.
- Second pass can run sequentially (small workers) or parallel per chunk.
- Blelloch scan is O(log n) depth — overkill for interviews; two-pass is enough.

## Q&A

**Q: Complexity?**  
A: O(n) work, O(workers) auxiliary.

**Q: Exclusive prefix?**  
A: Shift right or use `out[i-1]` when emitting.

**Q: Use case?**  
A: Parallel allocation, compaction, GPU-style algorithms.

**Q: Overflow?**  
A: int sum growth — use int64.

**Q: Edge cases?**  
A: workers=1 → sequential prefix; len=0 → nil.