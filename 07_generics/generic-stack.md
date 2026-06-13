# generic stack

## Live interview task
Implement a generic stack with Push, Pop and Len.

## Concepts covered
- generic types
- methods on generic types
- zero value clearing on Pop

## Candidate solution

```go
package main

import "fmt"

type Stack[T any] struct {
    data []T
}

func (s *Stack[T]) Push(v T) { s.data = append(s.data, v) }

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.data) == 0 {
        var zero T
        return zero, false
    }
    i := len(s.data) - 1
    v := s.data[i]
    var zero T
    s.data[i] = zero // clear for GC if T contains pointers
    s.data = s.data[:i]
    return v, true
}

func (s *Stack[T]) Len() int { return len(s.data) }

func main() {
    var s Stack[int]
    s.Push(1)
    s.Push(2)
    fmt.Println(s.Pop()) // 2 true
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Methods must use matching type params: `(s *Stack[T])`.
- `Pop` returns `(zero, false)` on empty — distinguish from legit zero value with `ok`.
- Pointer receiver for mutating methods; value receiver only for read-only `Len` if stack is small copy (usually use pointer).
- Pre-1.18: code generation or `interface{}` — mention generics replaced boilerplate.

## Q&A

**Q: Complexity?**  
A: Push/Pop amortized O(1).

**Q: Thread-safe?**  
A: No — wrap with mutex or channel.

**Q: Peek?**  
A: Return `s.data[len-1]` without shrink if non-empty.

**Q: Compare to `container/list`?**  
A: Slice stack is faster, better cache locality for interviews.

**Q: `Stack[int]` zero value?**  
A: Empty stack — `Pop` returns false.
