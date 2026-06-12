# lru cache with list and map

## Live interview task
Implement a small LRU cache with container/list and a map.

## Concepts covered
- maps
- linked list
- cache eviction

## Candidate solution

```go
package main

import (
    "container/list"
    "fmt"
)

type entry struct{ key string; value int }
type LRU struct { cap int; ll *list.List; items map[string]*list.Element }

func NewLRU(cap int) *LRU { return &LRU{cap: cap, ll: list.New(), items: make(map[string]*list.Element)} }

func (c *LRU) Get(k string) (int, bool) {
    if e, ok := c.items[k]; ok { c.ll.MoveToFront(e); return e.Value.(entry).value, true }
    return 0, false
}

func (c *LRU) Put(k string, v int) {
    if e, ok := c.items[k]; ok { e.Value = entry{k,v}; c.ll.MoveToFront(e); return }
    e := c.ll.PushFront(entry{k,v}); c.items[k] = e
    if c.ll.Len() > c.cap { old := c.ll.Back(); c.ll.Remove(old); delete(c.items, old.Value.(entry).key) }
}

func main() { c := NewLRU(2); c.Put("a",1); c.Put("b",2); c.Put("c",3); fmt.Println(c.Get("a"), c.Get("c")) }
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
