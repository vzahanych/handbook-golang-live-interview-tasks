# slices package contains clone go121

## Live interview task
Use the Go 1.21+ `slices` package for common slice operations instead of hand-rolled loops.

## Concepts covered
- slices package
- generics
- stdlib idioms

## `slices` package reference (Go 1.26.x)

All functions are generic; constraints vary (`comparable`, `cmp.Ordered`, or custom `func`).

| Category | Functions |
|----------|-----------|
| **Search** | `Contains`, `ContainsFunc`, `Index`, `IndexFunc`, `BinarySearch`, `BinarySearchFunc` |
| **Sort** | `Sort`, `SortFunc`, `SortStableFunc`, `IsSorted`, `IsSortedFunc` |
| **Compare** | `Equal`, `EqualFunc`, `Compare`, `CompareFunc` |
| **Min / max** | `Min`, `MinFunc`, `Max`, `MaxFunc` |
| **Copy / grow** | `Clone`, `Clip`, `Grow`, `Concat`, `Repeat` |
| **In-place edit** | `Reverse`, `Compact`, `CompactFunc`, `Delete`, `DeleteFunc`, `Insert`, `Replace` |
| **Iterators** (Go 1.23+) | `All`, `Backward`, `Values`, `Chunk`, `Collect`, `AppendSeq`, `Sorted`, `SortedFunc`, `SortedStableFunc` |

Full list (40 functions):

`All`, `AppendSeq`, `Backward`, `BinarySearch`, `BinarySearchFunc`, `Chunk`, `Clip`, `Clone`, `Collect`, `Compact`, `CompactFunc`, `Compare`, `CompareFunc`, `Concat`, `Contains`, `ContainsFunc`, `Delete`, `DeleteFunc`, `Equal`, `EqualFunc`, `Grow`, `Index`, `IndexFunc`, `Insert`, `IsSorted`, `IsSortedFunc`, `Max`, `MaxFunc`, `Min`, `MinFunc`, `Repeat`, `Replace`, `Reverse`, `Sort`, `SortFunc`, `SortStableFunc`, `Sorted`, `SortedFunc`, `SortedStableFunc`, `Values`

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
A: Core API in **Go 1.21** (`Clone`, `Contains`, `Sort`, `Delete`, …). Iterator helpers (`All`, `Chunk`, `Collect`, …) need **Go 1.23+**. Confirm `go` line in `go.mod`.
