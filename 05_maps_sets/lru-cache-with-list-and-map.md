# lru cache with list and map

## Live interview task
Implement a fixed-capacity **LRU (Least Recently Used) cache**: `Get`/`Put` in O(1) using a **map** (key → list node) plus **`container/list`** (recency order). On `Put` when full, evict the least recently used entry. `Get` and `Put` both mark an entry as most recently used (move to front). Example: capacity `2`, put `a`, put `b`, put `c` → `a` is evicted; `Get("a")` misses, `Get("c")` hits.

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

// entry is stored inside each list node; key is duplicated so eviction from
// the list tail can delete the matching map entry in O(1).
type entry struct {
    key   string
    value int
}

// LRU: map for O(1) lookup, doubly linked list for recency order.
// Front = most recently used (MRU); back = least recently used (LRU).
type LRU struct {
    cap   int
    ll    *list.List
    items map[string]*list.Element // key → node in ll
}

func NewLRU(cap int) *LRU {
    return &LRU{cap: cap, ll: list.New(), items: make(map[string]*list.Element)}
}

func (c *LRU) Get(k string) (int, bool) {
    if e, ok := c.items[k]; ok {
        c.ll.MoveToFront(e) // read counts as use — promote to MRU
        return e.Value.(entry).value, true
    }
    return 0, false
}

func (c *LRU) Put(k string, v int) {
    if e, ok := c.items[k]; ok {
        e.Value = entry{k, v}
        c.ll.MoveToFront(e) // update existing key — no size change
        return
    }
    e := c.ll.PushFront(entry{k, v}) // new key at MRU
    c.items[k] = e
    if c.ll.Len() > c.cap {
        old := c.ll.Back() // LRU node — least recently touched
        c.ll.Remove(old)
        delete(c.items, old.Value.(entry).key)
    }
}

func main() {
    c := NewLRU(2)
    c.Put("a", 1) // list: [a]
    c.Put("b", 2) // list: [b, a]  — b is MRU
    c.Put("c", 3) // list: [c, b]  — a evicted (was LRU at back)
    fmt.Println(c.Get("a")) // 0 false — evicted
    fmt.Println(c.Get("c")) // 3 true  — c still present
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
