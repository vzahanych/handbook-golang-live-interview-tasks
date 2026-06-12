# generic set operations

## Live interview task
Implement union, intersection and difference for generic sets.

## Concepts covered
- comparable
- generic maps
- sets

## Candidate solution

```go
package main

import "fmt"

type Set[T comparable] map[T]struct{}

func Union[T comparable](a, b Set[T]) Set[T] {
    out := make(Set[T], len(a)+len(b))
    for v := range a { out[v] = struct{}{} }
    for v := range b { out[v] = struct{}{} }
    return out
}

func Intersect[T comparable](a, b Set[T]) Set[T] {
    out := make(Set[T])
    if len(a) > len(b) { a, b = b, a }
    for v := range a { if _, ok := b[v]; ok { out[v] = struct{}{} } }
    return out
}

func main() { fmt.Println(Union(Set[int]{1:{}}, Set[int]{2:{}}), Intersect(Set[int]{1:{},2:{}}, Set[int]{2:{}})) }
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
