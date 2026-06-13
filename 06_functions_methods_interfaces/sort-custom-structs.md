# sort custom structs

## Live interview task
Sort structs by multiple fields using `sort.Slice`.

## Concepts covered
- sort.Slice
- less function
- stable vs unstable sort

## Candidate solution

```go
package main

import (
    "fmt"
    "sort"
)

type User struct {
    Name string
    Age  int
}

func main() {
    users := []User{{"Bob", 30}, {"Ann", 30}, {"Cat", 20}}
    sort.Slice(users, func(i, j int) bool {
        if users[i].Age != users[j].Age {
            return users[i].Age < users[j].Age
        }
        return users[i].Name < users[j].Name
    })
    fmt.Println(users) // [{Cat 20} {Ann 30} {Bob 30}]
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `sort.Slice` is **not stable** — equal keys may reorder; use `sort.SliceStable` if needed.
- `cmp.Or` + `cmp.Compare` (Go 1.21+) cleaner for multi-field compare.
- Implement `sort.Interface` on a named slice type for reusable sorts — `Len/Less/Swap`.
- `slices.SortFunc` (Go 1.21+) replaces `sort.Slice` for many cases.

## Q&A

**Q: Complexity?**  
A: O(n log n) comparisons.

**Q: Sort descending?**  
A: Invert comparison: `return users[i].Age > users[j].Age`.

**Q: `sort.Slice` panics?**  
A: If `less` is not a strict weak ordering — e.g. not transitive.

**Q: Generic alternative?**  
A: `slices.SortFunc(users, func(a, b User) int { ... })` returning -1/0/1.

**Q: Edge cases?**  
A: Empty slice, single element, all equal keys.
