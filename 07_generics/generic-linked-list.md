# generic linked list

## Live interview task
Build a minimal generic singly linked list.

## Concepts covered
- generic types
- recursive struct types
- pointers

## Candidate solution

```go
package main

import "fmt"

type Node[T any] struct {
    Value T
    Next  *Node[T]
}

type List[T any] struct {
    head *Node[T]
}

func (l *List[T]) PushFront(v T) {
    l.head = &Node[T]{Value: v, Next: l.head}
}

func (l *List[T]) Values() []T {
    out := make([]T, 0)
    for n := l.head; n != nil; n = n.Next {
        out = append(out, n.Value)
    }
    return out
}

func main() {
    var l List[string]
    l.PushFront("b")
    l.PushFront("a")
    fmt.Println(l.Values()) // [a b]
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Go allows generic struct with `*Node[T]` — recursive type param resolved at compile time.
- Linked list poor cache locality vs slice — mention when interviewer asks "why slice instead?"
- `PushFront` O(1); index access O(n).
- Clear list: set `head = nil` — GC collects nodes.

## Q&A

**Q: Doubly linked list?**  
A: Add `Prev *Node[T]` — `container/list` is non-generic doubly linked.

**Q: Thread-safe?**  
A: Mutex around mutations.

**Q: PopBack?**  
A: Need tail pointer or O(n) scan from head.

**Q: Complexity PushFront?**  
A: O(1) time, O(1) alloc per node.

**Q: Interview preference?**  
A: Often slice + append unless O(1) middle insert required (rare).
