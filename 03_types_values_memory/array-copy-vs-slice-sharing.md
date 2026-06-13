# array copy vs slice sharing

## Live interview task
Demonstrate that arrays copy their elements while slices share an underlying array.

## Concepts covered
- value representation
- arrays
- slices
- copy semantics

## Candidate solution

```go
package main

import "fmt"

func main() {
    a := [3]int{1, 2, 3}
    b := a
    b[0] = 99
    fmt.Println("array", a, b) // [1 2 3] [99 2 3]

    s := []int{1, 2, 3}
    t := s
    t[0] = 99
    fmt.Println("slice", s, t) // [99 2 3] [99 2 3]
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Array assignment copies the **entire value** — size is part of the type `[3]int` vs `[4]int` are incompatible.
- Slice assignment copies the **header** (pointer, len, cap) — both slices may alias the same backing array.
- Sub-slices share memory: `sub := s[:2]; sub[0] = 0` mutates `s[0]`.
- `copy(dst, src)` copies elements; `append` may allocate a new backing array when capacity exceeded.

## Q&A

**Q: When to use array vs slice?**  
A: Almost always slice. Arrays appear in crypto sizes, stack buffers, or `[N]T` as map keys when size is fixed.

**Q: How to get an independent slice?**  
A: `append([]T(nil), s...)` or `copy` into `make([]T, len(s))`.

**Q: Passing array to function?**  
A: Passed by value — full copy. Passing `*[N]int` or slice avoids copy.

**Q: Complexity?**  
A: Array assign O(n) in element count; slice assign O(1).

**Q: Full slice expression?**  
A: `s[low:high:max]` sets cap to `max-low`, limiting visibility of shared tail — useful before `append` on a sub-slice.
