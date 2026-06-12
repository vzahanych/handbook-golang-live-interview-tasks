# generic comparable map key cache

## Live interview task
Implement a cache where the key type must be comparable.

## Concepts covered
- type parameters
- comparable

## Candidate solution

```go
package main

import "fmt"

type Cache[K comparable, V any] struct { data map[K]V }

func NewCache[K comparable, V any]() *Cache[K,V] { return &Cache[K,V]{data: make(map[K]V)} }
func (c *Cache[K,V]) Get(k K) (V, bool) { v, ok := c.data[k]; return v, ok }
func (c *Cache[K,V]) Set(k K, v V) { c.data[k] = v }

func main() { c := NewCache[string,int](); c.Set("a", 1); fmt.Println(c.Get("a")) }
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
