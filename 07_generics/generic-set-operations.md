# generic set operations

## Live interview task
Implement union, intersection and difference for generic sets.

## Concepts covered
- comparable constraint
- generic maps as sets

## Candidate solution

```go
package main

import "fmt"

type Set[T comparable] map[T]struct{}

func Union[T comparable](a, b Set[T]) Set[T] {
    out := make(Set[T], len(a)+len(b))
    for v := range a {
        out[v] = struct{}{}
    }
    for v := range b {
        out[v] = struct{}{}
    }
    return out
}

func Intersect[T comparable](a, b Set[T]) Set[T] {
    out := make(Set[T])
    if len(a) > len(b) {
        a, b = b, a // iterate smaller set
    }
    for v := range a {
        if _, ok := b[v]; ok {
            out[v] = struct{}{}
        }
    }
    return out
}

func Difference[T comparable](a, b Set[T]) Set[T] {
    out := make(Set[T])
    for v := range a {
        if _, ok := b[v]; !ok {
            out[v] = struct{}{}
        }
    }
    return out
}

func main() {
    a := Set[int]{1: {}, 2: {}}
    b := Set[int]{2: {}, 3: {}}
    fmt.Println(Union(a, b))       // 1 2 3
    fmt.Println(Intersect(a, b))   // 2
    fmt.Println(Difference(a, b))  // 1
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `T comparable` required for map keys.
- Intersect iterates smaller set — O(min(|a|,|b|)).
- Union allocates new set — does not mutate inputs.
- `struct{}` values — idiomatic empty set member.

## Q&A

**Q: Union complexity?**  
A: O(|a|+|b|).

**Q: Subset check?**  
A: `len(Intersect(a,b)) == len(a)`.

**Q: Slice not comparable?**  
A: Cannot be map key — hash serialized form or use custom set struct.

**Q: Immutable sets?**  
A: Return new sets from ops (as here) — functional style.

**Q: Production?**  
A: `map[T]struct{}` or third-party; generics avoid copy-paste per type.
