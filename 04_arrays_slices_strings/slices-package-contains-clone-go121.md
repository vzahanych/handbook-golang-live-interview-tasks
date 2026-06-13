# slices package contains clone go121

## Live interview task
Use the Go 1.21+ `slices` package for common slice operations instead of hand-rolled loops.

## Concepts covered
- slices package
- generics
- stdlib idioms

## Candidate solution

```go
package main

import (
    "fmt"
    "slices"
)

func main() {
    nums := []int{3, 1, 4, 1, 5}
    cloned := slices.Clone(nums)
    cloned[0] = 99
    fmt.Println("orig", nums)   // unchanged
    fmt.Println("clone", cloned)

    fmt.Println(slices.Contains(nums, 4))       // true
    fmt.Println(slices.Index(nums, 1))          // 1 — first index
    slices.Sort(nums)
    fmt.Println(nums)                           // [1 1 3 4 5]
    fmt.Println(slices.BinarySearch(nums, 4))  // 3, true
}
```

## Run

```bash
go run . # Go 1.21+
```

## Interview notes / pitfalls
- `slices` works on slices of **comparable** or **ordered** types via constraints — not all ops on all types.
- `slices.Clone` allocates new backing array — unlike sub-slice share.
- `slices.Sort` is pdqsort — O(n log n); not stable; use `slices.SortStableFunc` if needed.
- Prefer stdlib in production; implement manually when interviewer asks "without slices package".

## Q&A

**Q: `slices.Equal` vs `reflect.DeepEqual`?**  
A: `slices.Equal` is typed, faster, element `==` only — no nested struct deep compare.

**Q: Delete element?**  
A: `slices.Delete(s, i, i+1)` — may not zero tail; `DeleteFunc` for filter.

**Q: Compact?**  
A: `slices.Compact` removes consecutive duplicates — slice must be sorted first.

**Q: Why know both manual and stdlib?**  
A: Interviews test understanding; production uses stdlib.

**Q: Go version?**  
A: `go 1.21` in mod — `slices` and `maps` packages landed together.
