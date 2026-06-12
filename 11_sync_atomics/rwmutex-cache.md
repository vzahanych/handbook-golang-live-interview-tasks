# rwmutex cache

## Live interview task
Use sync.RWMutex for a read-heavy cache.

## Concepts covered
- sync.RWMutex
- maps are not concurrency-safe

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

type Cache struct { mu sync.RWMutex; data map[string]string }
func New() *Cache { return &Cache{data: make(map[string]string)} }
func (c *Cache) Get(k string) (string, bool) { c.mu.RLock(); defer c.mu.RUnlock(); v, ok := c.data[k]; return v, ok }
func (c *Cache) Set(k, v string) { c.mu.Lock(); defer c.mu.Unlock(); c.data[k] = v }

func main() { c := New(); c.Set("a", "1"); fmt.Println(c.Get("a")) }
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
