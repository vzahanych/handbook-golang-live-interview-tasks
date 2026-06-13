# lru cache with list and map

## Live interview task
Implement a small LRU cache with `container/list` and a map.

## Concepts covered
- maps
- doubly linked list
- cache eviction
- O(1) get/put

## Candidate solution

```go
package main

import (
    "container/list"
    "fmt"
)

type entry struct {
    key   string
    value int
}

type LRU struct {
    cap   int
    ll    *list.List
    items map[string]*list.Element
}

func NewLRU(cap int) *LRU {
    return &LRU{cap: cap, ll: list.New(), items: make(map[string]*list.Element)}
}

func (c *LRU) Get(k string) (int, bool) {
    if e, ok := c.items[k]; ok {
        c.ll.MoveToFront(e)
        return e.Value.(entry).value, true
    }
    return 0, false
}

func (c *LRU) Put(k string, v int) {
    if e, ok := c.items[k]; ok {
        e.Value = entry{k, v}
        c.ll.MoveToFront(e)
        return
    }
    e := c.ll.PushFront(entry{k, v})
    c.items[k] = e
    if c.ll.Len() > c.cap {
        old := c.ll.Back()
        c.ll.Remove(old)
        delete(c.items, old.Value.(entry).key)
    }
}

func main() {
    c := NewLRU(2)
    c.Put("a", 1)
    c.Put("b", 2)
    c.Put("c", 3) // evicts a
    fmt.Println(c.Get("a")) // 0 false
    fmt.Println(c.Get("c")) // 3 true
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Map: key → list element pointer for O(1) lookup.
- List: front = MRU, back = LRU — evict from back.
- `Get` must **move to front** — touch updates recency.
- `Put` existing key: update value + move front, no size change.
- Not thread-safe — wrap with mutex for concurrent use.

## Q&A

**Q: Complexity?**  
A: Get/Put O(1) average; list + map operations constant time.

**Q: `cap == 0`?**  
A: Define behavior — reject puts or evict every put; clarify in API.

**Q: Why `container/list`?**  
A: Doubly linked list with O(1) move/remove; Go 1.21+ you might use custom struct nodes.

**Q: vs `sync.Map`?**  
A: Different problem — LRU needs ordering; `sync.Map` is concurrent map without LRU.

**Q: LeetCode 146?**  
A: Classic — same data structure interview question.

**Q: Edge cases?**  
A: cap 1, update existing key, get missing key.
