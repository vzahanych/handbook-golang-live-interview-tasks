# mutex protected counter

## Live interview task
Protect shared state with `sync.Mutex`.

## Concepts covered
- sync.Mutex
- data races

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

type Counter struct {
    mu sync.Mutex
    n  int
}

func (c *Counter) Inc() {
    c.mu.Lock()
    c.n++
    c.mu.Unlock()
}

func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.n
}

func main() {
    var c Counter
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            c.Inc()
        }()
    }
    wg.Wait()
    fmt.Println(c.Value()) // 1000
}
```

## Run

```bash
go run .
go test -race . # in test package
```

## Interview notes / pitfalls
- Mutex protects **invariants** across multiple fields — not just single int (atomics may suffice for one word).
- Always unlock — use `defer mu.Unlock()` after Lock in functions with multiple returns.
- **Don't** copy struct containing `sync.Mutex` — mutex must not be copied after first use.
- Lock ordering: A then B everywhere — prevents deadlock.

## Q&A

**Q: Mutex vs atomic?**  
A: Atomic for single counter/flag; mutex for compound state or multiple variables.

**Q: `RWMutex` when?**  
A: Read-heavy, rare writes — see rwmutex-cache.

**Q: Deadlock?**  
A: Lock twice same goroutine without `RWMutex` — panic; or circular wait across mutexes.

**Q: Test races?**  
A: `go test -race` — essential.

**Q: Complexity?**  
A: Inc O(1); contention serializes goroutines.
