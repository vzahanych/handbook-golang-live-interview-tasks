# preallocate slice before append

## Live interview task
Preallocate slice capacity before append to avoid repeated growth allocations.

## Concepts covered
- make with cap
- append growth
- allocations

## Candidate solution

```go
package main

import "fmt"

func build(n int) []int {
    s := make([]int, 0, n)
    for i := 0; i < n; i++ {
        s = append(s, i)
    }
    return s
}

func main() {
    fmt.Println(len(build(1000)), cap(build(1000)))
}
```

## Run

```bash
go run .
go test -bench=. -benchmem # if benchmark added
```

## Interview notes / pitfalls
- `append` on zero cap doubles backing array — O(n log n) copies total without prealloc.
- `make([]T, 0, n)` — len 0, cap n — one allocation.
- Over-estimating cap wastes memory — estimate upper bound when known.
- `len` vs `cap`: prealloc cap does not set len — still use append or `make([]T, n)` if fixed size.

## Q&A

**Q: When `make([]T, n)` instead?**  
A: When you will fill all slots by index — no append needed.

**Q: Verify savings?**  
A: `go test -benchmem` — compare allocs/op.

**Q: Complexity?**  
A: O(n) with one alloc vs O(n) amortized with multiple growth copies.

**Q: Edge cases?**  
A: n=0 → empty slice; n unknown — grow dynamically or use conservative estimate.

**Q: Production?**  
A: Pre-size from DB `COUNT` or prior run metrics.
