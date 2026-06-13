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
    if len(a) == 0 {
        return 0
    }
    workers := runtime.GOMAXPROCS(0)
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
            for _, v := range a[lo:hi] {
                partial[w] += v
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
    fmt.Println(ParallelSum([]int{1, 2, 3, 4, 5, 6})) // 21
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Each worker writes **own** `partial[w]` slot — avoids false sharing on single shared counter.
- Chunk boundaries: `lo = w*n/workers`, `hi = (w+1)*n/workers` — covers all elements, no gaps.
- Parallel sum only wins on **large** slices — goroutine overhead dominates small inputs.
- Integer overflow not checked — mention if interviewer cares.

## Q&A

**Q: Complexity?**  
A: O(n) work, O(workers) extra space; wall time ~O(n/workers) with enough CPUs.

**Q: False sharing?**  
A: Adjacent `partial` ints may share cache line — pad or use per-goroutine local var then merge once.

**Q: `GOMAXPROCS`?**  
A: Default worker count; cap at `len(a)` to avoid idle workers.

**Q: vs sequential?**  
A: Benchmark crossover often thousands+ elements — state that in interview.

**Q: Edge cases?**  
A: Empty slice → 0; single element; negative numbers.
