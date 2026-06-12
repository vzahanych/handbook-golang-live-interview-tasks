# ttl cache with cleanup goroutine

## Live interview task
Implement an in-memory TTL cache with a cleanup goroutine.

## Concepts covered
- generics
- TTL cache
- sync.RWMutex

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type item[V any] struct { value V; expires time.Time }
type Cache[K comparable, V any] struct { mu sync.RWMutex; data map[K]item[V] }
func New[K comparable, V any]() *Cache[K,V] { return &Cache[K,V]{data: make(map[K]item[V])} }
func (c *Cache[K,V]) Set(k K, v V, ttl time.Duration) { c.mu.Lock(); c.data[k] = item[V]{v, time.Now().Add(ttl)}; c.mu.Unlock() }
func (c *Cache[K,V]) Get(k K) (V, bool) { c.mu.RLock(); it, ok := c.data[k]; c.mu.RUnlock(); var zero V; if !ok || time.Now().After(it.expires) { return zero, false }; return it.value, true }

func main() { c := New[string,int](); c.Set("a", 1, time.Second); fmt.Println(c.Get("a")) }
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- None specific; discuss edge cases and complexity.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
