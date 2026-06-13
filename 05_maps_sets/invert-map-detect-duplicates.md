# invert map detect duplicates

## Live interview task
Invert a map `K→V` to `V→K` while detecting duplicate values.

## Concepts covered
- maps
- generics
- comparable
- errors

## Candidate solution

```go
package main

import (
    "errors"
    "fmt"
)

func invert[K comparable, V comparable](m map[K]V) (map[V]K, error) {
    out := make(map[V]K, len(m))
    for k, v := range m {
        if _, exists := out[v]; exists {
            return nil, errors.New("duplicate value")
        }
        out[v] = k
    }
    return out, nil
}

func main() {
    fmt.Println(invert(map[string]int{"a": 1, "b": 2}))
    fmt.Println(invert(map[string]int{"a": 1, "b": 1})) // error
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Bijection required for clean invert — duplicate values make inverse ambiguous.
- `V` must be comparable to be map key.
- Non-comparable `K` cannot be used in `map[K]V` anyway.
- Use case: reverse lookup tables, enum name ↔ id.

## Q&A

**Q: Allow duplicate values → `map[V][]K`?**  
A: Yes — collect all keys per value for one-to-many.

**Q: Complexity?**  
A: O(n) time and space.

**Q: Empty map?**  
A: Returns empty inverted map, nil error.

**Q: Nil input map?**  
A: Range over nil is no-op — returns empty map.

**Q: Production?**  
A: Return typed error `ErrDuplicateValue` for `errors.Is` checks.
