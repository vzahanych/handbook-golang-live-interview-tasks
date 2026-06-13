# map concurrent write panic

## Live interview task
Explain why concurrent map writes panic and show the mutex-protected fix.

## Concepts covered
- maps
- concurrency
- sync.Mutex
- fatal error

## Buggy version

```go
// DO NOT RUN with -race expecting clean exit — concurrent write panics:
// fatal error: concurrent map writes

func broken() {
    m := make(map[int]int)
    go func() { m[1] = 1 }()
    go func() { m[2] = 2 }()
}
```

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

type SafeMap struct {
    mu sync.Mutex
    m  map[string]int
}

func (s *SafeMap) Inc(key string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.m[key]++
}

func main() {
    sm := &SafeMap{m: make(map[string]int)}
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            sm.Inc("x")
        }()
    }
    wg.Wait()
    fmt.Println(sm.m["x"]) // 100
}
```

## Run

```bash
go run .
go test -race ./... # when in a test package
```

## Interview notes / pitfalls
- Runtime detects concurrent map read + write or write + write — **panic**, not data race report only.
- Concurrent **read-only** maps are safe if no writer exists.
- `sync.Map` for read-heavy cache; `map + RWMutex` for general case; never raw map shared across goroutines.
- `maps.Clone` then process copy — fork-on-read pattern for snapshots.

## Q&A

**Q: Read while write?**  
A: Also undefined — can panic; protect all access or use `RWMutex`.

**Q: `sync.Map` when?**  
A: Many goroutines, keys stable, mostly Load/Store, few deletes.

**Q: Channel instead of mutex?**  
A: Single goroutine owns map, others send ops — actor pattern.

**Q: Detect in tests?**  
A: `go test -race` — essential for concurrent code.

**Q: One-liner for interview?**  
A: "Maps are not safe for concurrent use without synchronization."
