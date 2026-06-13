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
    if workers > len(in) {
        workers = len(in)
    }
    if workers == 0 {
        return nil
    }

    chunks := make([][]T, workers)
    var wg sync.WaitGroup
    for w := 0; w < workers; w++ {
        lo := w * len(in) / workers
        hi := (w + 1) * len(in) / workers
        wg.Add(1)
        go func(w, lo, hi int) {
            defer wg.Done()
            for _, v := range in[lo:hi] {
                if keep(v) {
                    chunks[w] = append(chunks[w], v)
                }
            }
        }(w, lo, hi)
    }
    wg.Wait()

    var out []T
    for _, c := range chunks {
        out = append(out, c...)
    }
    return out
}

func main() {
    fmt.Println(ParallelFilter([]int{1, 2, 3, 4}, func(x int) bool { return x%2 == 0 }))
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- **Stable**: process contiguous chunks in order, concatenate chunks `[0],[1],...` — global order preserved.
- Each chunk filtered independently — `keep` must not depend on global state unless synchronized.
- Two-phase: parallel filter → sequential merge (cheap concat).
- Unstable parallel filter: shared output channel — faster but wrong order.

## Q&A

**Q: Complexity?**  
A: O(n) time, O(k) output size; merge O(workers) chunk appends.

**Q: Preallocate output?**  
A: Estimate density of `keep` — hard; optional second pass count.

**Q: vs `slices.DeleteFunc`?**  
A: Sequential in-place — parallel only for large n.

**Q: Predicate expensive?**  
A: Parallel wins when `keep` dominates CPU.

**Q: Edge cases?**  
A: None match → empty slice; all match → copy of input order.
