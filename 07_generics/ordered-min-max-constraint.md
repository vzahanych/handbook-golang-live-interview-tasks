# ordered min max constraint

## Live interview task
Define an `Ordered` constraint and implement generic Min and Max.

## Concepts covered
- type constraints
- underlying type terms
- `~T`

## Candidate solution

```go
package main

import (
    "cmp"
    "fmt"
)

// Prefer stdlib in production:
// import "cmp" — Min, Max, Compare for Ordered types (Go 1.21+)

type Ordered interface {
    ~int | ~int64 | ~float64 | ~string
}

func Min[T Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

type Age int

func main() {
    fmt.Println(Min(Age(10), Age(20))) // 10
    fmt.Println(cmp.Min(3, 1))         // 1
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `~int` allows defined types with underlying `int` like `type Age int`.
- Union in constraint is OR — type must match one branch.
- `cmp.Ordered` in stdlib replaces hand-rolled Ordered for Min/Max/Compare.
- Floats: `NaN` breaks `<` ordering — mention if interviewer asks.

## Q&A

**Q: Why `~`?**  
A: Includes defined types whose underlying type is `int`, not just bare `int`.

**Q: Custom type not in union?**  
A: Compile error — extend constraint or use `cmp.Compare` with ordered types.

**Q: Max of slice?**  
A: Loop or `slices.Max` (Go 1.21+) with `cmp.Ordered`.

**Q: Complexity?**  
A: O(1) for two values.

**Q: vs `constraints.Ordered`?**  
A: `golang.org/x/exp/constraints` — largely superseded by `cmp` package.
