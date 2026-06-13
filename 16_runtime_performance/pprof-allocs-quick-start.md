# pprof allocs quick start

## Live interview task
Profile heap allocations with `runtime/pprof` or test benchmark.

## Concepts covered
- pprof
- alloc profiling
- performance workflow

## Candidate solution

```go
package main

import (
    "fmt"
    "os"
    "runtime/pprof"
)

func leaky(n int) {
    _ = make([]byte, n)
}

func main() {
    f, _ := os.Create("allocs.prof")
    defer f.Close()

    _ = pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()

    for i := 0; i < 1000; i++ {
        leaky(1024 * i)
    }
    fmt.Println("wrote allocs.prof — use: go tool pprof allocs.prof")
}
```

## Benchmark alternative

```bash
go test -bench=. -benchmem -memprofile=mem.prof
go tool pprof -top mem.prof
```

## Interview notes / pitfalls
- **Measure first** — profile before micro-optimizing slices/maps.
- `-benchmem` shows allocs/op in benchmarks — quick interview demo.
- `go tool pprof -http=:0 cpu.prof` — flame graph UI.
- CPU profile vs heap profile — different questions.

## Q&A

**Q: Workflow?**  
A: Benchmark → identify hot path → profile → fix → re-benchmark.

**Q: `testing.AllocsPerRun`?**  
A: Assert max allocs in unit tests for critical paths.

**Q: Production?**  
A: `net/http/pprof` on debug port — protect with auth/network policy.

**Q: Trace vs profile?**  
A: `runtime/trace` for latency; pprof for CPU/heap.

**Q: Read flame graph?**  
A: Wide box = more time/allocs in that frame.
