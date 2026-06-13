# generic comparable map key cache

## Live interview task
Implement a cache where the key type must be comparable.

## Concepts covered
- type parameters
- comparable constraint
- map keys

## Candidate solution

```go
package main

import "fmt"

type Cache[K comparable, V any] struct {
    data map[K]V
}

func NewCache[K comparable, V any]() *Cache[K, V] {
    return &Cache[K, V]{data: make(map[K]V)}
}

func (c *Cache[K, V]) Get(k K) (V, bool) {
    v, ok := c.data[k]
    return v, ok
}

func (c *Cache[K, V]) Set(k K, v V) {
    c.data[k] = v
}

func main() {
    c := NewCache[string, int]()
    c.Set("a", 1)
    fmt.Println(c.Get("a")) // 1 true
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `K comparable` — all map key rules apply (no slice/map/func keys).
- `V any` — value can be anything including pointers.
- Not concurrent — add `sync.RWMutex` for shared cache.
- LRU/TTL are extensions — separate fields and eviction logic.

## Q&A

**Q: Why `comparable` not `any` for K?**  
A: Map keys must support `==`.

**Q: Struct key?**  
A: OK if all fields comparable — no slices inside.

**Q: Get-or-compute?**  
A: `sync.Singleflight` or `if !ok { v = compute(); Set }`.

**Q: Complexity?**  
A: Get/Set O(1) average.

**Q: vs `sync.Map`?**  
A: `sync.Map` is `any` keys, concurrent, no generics — different use case.
