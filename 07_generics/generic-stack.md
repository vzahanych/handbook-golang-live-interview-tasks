# generic stack

## Live interview task
Implement a generic stack with Push, Pop and Len.

## Concepts covered
- generic type
- methods on generic types
- zero value

## Candidate solution

```go
package main

import "fmt"

type Stack[T any] struct { data []T }
func (s *Stack[T]) Push(v T) { s.data = append(s.data, v) }
func (s *Stack[T]) Pop() (T, bool) {
    if len(s.data) == 0 { var zero T; return zero, false }
    i := len(s.data)-1
    v := s.data[i]
    var zero T
    s.data[i] = zero
    s.data = s.data[:i]
    return v, true
}
func (s *Stack[T]) Len() int { return len(s.data) }

func main() { var s Stack[int]; s.Push(1); s.Push(2); fmt.Println(s.Pop()) }
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
