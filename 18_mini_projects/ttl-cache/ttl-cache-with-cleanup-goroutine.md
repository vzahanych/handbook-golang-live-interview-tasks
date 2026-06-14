# ttl cache with cleanup goroutine

## Live interview task
Implement in-memory TTL cache with lazy expiry on Get and optional cleanup goroutine.

## Concepts covered
- TTL cache
- RWMutex
- generics
- time

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type item[V any] struct {
    value   V
    expires time.Time
}

type Cache[K comparable, V any] struct {
    mu   sync.RWMutex
    data map[K]item[V]
}

func New[K comparable, V any]() *Cache[K, V] {
    return &Cache[K, V]{data: make(map[K]item[V])}
}

func (c *Cache[K, V]) Set(k K, v V, ttl time.Duration) {
    c.mu.Lock()
    c.data[k] = item[V]{v, time.Now().Add(ttl)}
    c.mu.Unlock()
}

func (c *Cache[K, V]) Get(k K) (V, bool) {
    c.mu.RLock()
    it, ok := c.data[k]
    c.mu.RUnlock()
    var zero V
    if !ok || time.Now().After(it.expires) {
        return zero, false
    }
    return it.value, true
}

func (c *Cache[K, V]) cleanupLoop(interval time.Duration, stop <-chan struct{}) {
    ticker := time.NewTicker(interval)
    defer ticker.Stop()
    for {
        select {
        case <-stop:
            return
        case <-ticker.C:
            c.mu.Lock()
            now := time.Now()
            for k, it := range c.data {
                if now.After(it.expires) {
                    delete(c.data, k)
                }
            }
            c.mu.Unlock()
        }
    }
}

func main() {
    c := New[string, int]()
    c.Set("a", 1, time.Second)
    fmt.Println(c.Get("a"))
}
```

## Run

Runnable version lives in [ttl-cache/](ttl-cache/main.go). It adds `Delete`/`Len`
and a `StartCleanup` that returns a stop func, and `main` demonstrates lazy
expiry vs. the background sweep.

```bash
go run ./18_mini_projects/ttl-cache
# prove the locking is correct:
go run -race ./18_mini_projects/ttl-cache
```

## Interview notes / pitfalls
- Lazy expiry on `Get` — stale keys remain in map until cleanup or access.
- Background cleanup prevents unbounded map growth from expired keys never read.
- `RWMutex` on Get — write on Set/delete; consider sharding for hot paths.
- Not distributed — single process; use Redis TTL for cluster.

## Q&A

**Q: `sync.Map`?**  
A: Possible but TTL logic still needed per entry.

**Q: Clock skew?**  
A: Use monotonic `time.Now()` relative TTL — OK on one machine.

**Q: Complexity?**  
A: Get/Set O(1); cleanup O(entries) per tick.

**Q: Extend?**  
A: LRU + TTL, max size eviction.

**Q: Interview variant?**  
A: Implement `Delete` and `Len` for completeness.
