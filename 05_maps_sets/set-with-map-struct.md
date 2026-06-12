# set with map struct

## Live interview task
Implement a generic set with map[T]struct{}.

## Concepts covered
- maps
- struct{}
- generics
- comparable

## Candidate solution

```go
package main

import "fmt"

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(v T) { s[v] = struct{}{} }
func (s Set[T]) Has(v T) bool { _, ok := s[v]; return ok }
func (s Set[T]) Delete(v T) { delete(s, v) }

func main() {
    s := make(Set[string])
    s.Add("go")
    fmt.Println(s.Has("go"), s.Has("rust"))
}
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
