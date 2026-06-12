# generic linked list

## Live interview task
Build a minimal generic singly linked list.

## Concepts covered
- recursive generic types
- pointers

## Candidate solution

```go
package main

import "fmt"

type Node[T any] struct { Value T; Next *Node[T] }
type List[T any] struct { head *Node[T] }

func (l *List[T]) PushFront(v T) { l.head = &Node[T]{Value: v, Next: l.head} }
func (l *List[T]) Values() []T {
    var out []T
    for n := l.head; n != nil; n = n.Next { out = append(out, n.Value) }
    return out
}

func main() { var l List[string]; l.PushFront("b"); l.PushFront("a"); fmt.Println(l.Values()) }
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
