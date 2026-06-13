# set with map struct

## Live interview task
Implement a generic set with `map[T]struct{}`.

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

func (s Set[T]) Add(v T)       { s[v] = struct{}{} }
func (s Set[T]) Has(v T) bool  { _, ok := s[v]; return ok }
func (s Set[T]) Delete(v T)    { delete(s, v) }
func (s Set[T]) Len() int      { return len(s) }

func main() {
    s := make(Set[string])
    s.Add("go")
    fmt.Println(s.Has("go"), s.Has("rust")) // true false
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `struct{}` uses zero bytes per value — set marker only; `map[T]bool` wastes one byte per key.
- Keys must be **comparable** — slices, maps, funcs cannot be keys.
- `Add` is idempotent; no error on duplicate.
- Go 1.23+ has `map[T]struct{}` patterns; no stdlib Set type until you use custom or third-party.

## Q&A

**Q: Why not `map[T]bool`?**  
A: Works but `true` values cost memory; `struct{}` is idiomatic empty value.

**Q: Iterate set?**  
A: `for k := range s { ... }` — order random.

**Q: Union/intersection?**  
A: Union: add all from other; intersection: range self, delete if not in other.

**Q: Complexity?**  
A: Add/Has/Delete average O(1).

**Q: `[]T` as key?**  
A: Invalid — use `string` hash of serialized form or `map[string]struct{}` with canonical encoding.
