# parallel histogram by buckets

## Live interview task
Build a histogram in parallel using worker-local bucket arrays.

## Concepts covered
- histogram
- parallel aggregation
- bucket indexing

## Candidate solution

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
)

func Histogram(a []int, bucketCount int) []int {
    workers := runtime.GOMAXPROCS(0)
    if workers > len(a) {
        workers = len(a)
    }
    if workers == 0 || bucketCount == 0 {
        return nil
    }

    local := make([][]int, workers)
    var wg sync.WaitGroup
    for w := 0; w < workers; w++ {
        lo := w * len(a) / workers
        hi := (w + 1) * len(a) / workers
        local[w] = make([]int, bucketCount)
        wg.Add(1)
        go func(w, lo, hi int) {
            defer wg.Done()
            for _, v := range a[lo:hi] {
                if v >= 0 && v < bucketCount {
                    local[w][v]++
                }
            }
        }(w, lo, hi)
    }
    wg.Wait()

    out := make([]int, bucketCount)
    for _, part := range local {
        for i, v := range part {
            out[i] += v
        }
    }
    return out
}

func main() {
    fmt.Println(Histogram([]int{1, 1, 2, 3, 3, 3}, 5))
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Values as direct bucket indices — generalize with `bucket = v / width` for ranges.
- Each worker has full `bucketCount` slice — memory O(workers * buckets); OK for small buckets.
- Out-of-range values skipped — document or use clamp.

## Q&A

**Q: Complexity?**  
A: O(n + workers*buckets) merge.

**Q: Large bucket range?**  
A: Use map[int]int per worker instead of dense slice.

**Q: vs atomic increment on shared buckets?**  
A: Atomic per bucket works but false sharing on hot buckets — local arrays better.

**Q: Real use?**  
A: Latency histograms, score distributions.

**Q: Edge cases?**  
A: Empty input → zero buckets.
