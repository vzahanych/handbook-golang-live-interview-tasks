# rwmutex cache

## Live interview task
Use `sync.RWMutex` for a read-heavy in-memory cache.

## Concepts covered
- sync.RWMutex
- read/write lock
- map concurrency

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

type Cache struct {
    mu   sync.RWMutex
    data map[string]string
}

func NewCache() *Cache {
    return &Cache{data: make(map[string]string)}
}

func (c *Cache) Get(k string) (string, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    v, ok := c.data[k]
    return v, ok
}

func (c *Cache) Set(k, v string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[k] = v
}

func main() {
    c := NewCache()
    c.Set("a", "1")
    fmt.Println(c.Get("a"))
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Multiple `RLock` concurrent; `Lock` exclusive — writers block all readers.
- `RLock` cannot upgrade to `Lock` — unlock R, acquire Lock (race window) or use single Mutex.
- Map still needs lock — raw map concurrent access panics.
- `sync.Map` alternative for read-heavy cache with stable keys — profile first.

## Q&A

**Q: When RWMutex slower than Mutex?**  
A: Write-heavy or very short critical sections — RWMutex has more overhead.

**Q: Get-or-create?**  
A: `Lock` full path: double-check after `RLock` upgrade pattern or single `Lock`.

**Q: Copy map under RLock?**  
A: Snapshot for iteration — hold RLock during copy or copy keys slice.

**Q: Complexity?**  
A: Get O(1) under RLock; Set O(1) under Lock.

**Q: Starvation?**  
A: Go RWMutex favors writers in recent versions — mention fairness if asked deeply.
