# atomic vs mutex when

## Live interview task
Choose between atomics and mutex for concurrent counter + explain trade-offs.

## Concepts covered
- atomic vs mutex
- when to use which
- compound state

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

// Simple counter — atomic is enough
type AtomicCounter struct {
    n atomic.Int64
}

func (c *AtomicCounter) Inc() { c.n.Add(1) }
func (c *AtomicCounter) Val() int64 { return c.n.Load() }

// Counter + last update time — need mutex (two fields)
type Stats struct {
    mu   sync.Mutex
    n    int64
    last int64
}

func (s *Stats) Inc(ts int64) {
    s.mu.Lock()
    s.n++
    s.last = ts
    s.mu.Unlock()
}

func main() {
    var a AtomicCounter
    a.Inc()
    fmt.Println(a.Val())

    var st Stats
    st.Inc(100)
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- **Atomic**: single word, counters, flags, publish pointer via `atomic.Value` / `atomic.Pointer`.
- **Mutex**: multiple fields, invariants spanning data, maps, slices, complex logic.
- **RWMutex**: many readers, few writers.
- Wrong: mutex for every int increment in hot path; atomic for map update.

## Q&A

**Q: `atomic.Pointer[T]`?**  
A: Lock-free swap of config snapshot — readers Load without mutex.

**Q: Performance?**  
A: Atomics scale better for single hot counter; profile before optimizing.

**Q: `sync/atomic` + mutex together?**  
A: Common — atomic for stats, mutex for structure.

**Q: Interview answer template?**  
A: "One number → atomic; struct/map or multiple fields → mutex."

**Q: Race on atomic?**  
A: Incorrect mixing atomic and non-atomic access to same var — still a race.
