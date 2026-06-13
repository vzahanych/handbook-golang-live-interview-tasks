# slices sortfunc generic go121

## Live interview task
Sort a slice of custom structs using `slices.SortFunc` and `cmp.Compare`.

## Concepts covered
- slices package
- cmp.Compare
- generic sorting

## Candidate solution

```go
package main

import (
    "cmp"
    "fmt"
    "slices"
)

type User struct {
    Name string
    Age  int
}

func main() {
    users := []User{{"Bob", 30}, {"Ann", 30}, {"Cat", 20}}
    slices.SortFunc(users, func(a, b User) int {
        if c := cmp.Compare(a.Age, b.Age); c != 0 {
            return c
        }
        return cmp.Compare(a.Name, b.Name)
    })
    fmt.Println(users)
}
```

## Run

```bash
go run . # Go 1.21+
```

## Interview notes / pitfalls
- `SortFunc` less function returns `int` — negative/zero/positive like `strcmp`.
- `cmp.Compare` works on `cmp.Ordered` types — cleaner than manual `<`.
- `slices.Sort` for simple ordered element types only.
- Stable sort: `slices.SortStableFunc`.

## Q&A

**Q: vs `sort.Slice`?**  
A: `slices.SortFunc` is generic, type-safe, preferred in new code.

**Q: Complexity?**  
A: O(n log n).

**Q: Reverse sort?**  
A: Negate compare: `return -cmp.Compare(a.Age, b.Age)`.

**Q: Nil slice?**  
A: Sort is no-op on nil/empty.

**Q: Multi-field?**  
A: Chain `cmp.Compare` calls — first nonzero wins.
