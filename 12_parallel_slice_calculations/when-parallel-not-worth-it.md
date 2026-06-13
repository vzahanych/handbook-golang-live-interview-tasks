# when parallel not worth it

## Live interview task
Explain when parallelizing slice work hurts performance and show a simple threshold rule.

## Concepts covered
- Amdahl's law
- goroutine overhead
- granularity

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

const parallelThreshold = 10_000

func SumSequential(a []int) int {
    s := 0
    for _, v := range a {
        s += v
    }
    return s
}

func SumParallel(a []int, workers int) int {
    if len(a) < parallelThreshold {
        return SumSequential(a)
    }
    // ... use ParallelSum pattern for large slices
    return SumSequential(a) // placeholder in demo
}

func main() {
    small := make([]int, 100)
    for i := range small {
        small[i] = i
    }
    fmt.Println(SumParallel(small, 4)) // sequential path
}
```

## Interview notes / pitfalls
- Goroutine creation + scheduling ~ microseconds — tiny work per item loses.
- Sync overhead (WaitGroup, channel) dominates small n.
- False sharing and cache thrashing can make parallel **slower** than sequential.
- Rule of thumb: parallelize when work per element is costly OR n > ~10k (benchmark your hardware).

## Q&A

**Q: Amdahl's law?**  
A: Speedup limited by serial fraction — merge phase, coordination.

**Q: `GOMAXPROCS`?**  
A: More workers than CPUs → contention, not always faster.

**Q: When always parallel?**  
A: Independent blocking I/O per item — not CPU slice scan.

**Q: Interview answer?**  
A: "Measure first; default sequential until proven slow."

**Q: `testing.B`?**  
A: Benchmark sequential vs parallel with `ReportAllocs`.
