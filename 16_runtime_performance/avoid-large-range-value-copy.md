# avoid large range value copy

## Live interview task
Avoid copying large structs in `for range` by iterating indexes or pointers.

## Concepts covered
- range copy semantics
- large structs
- index iteration

## Candidate solution

```go
package main

import "fmt"

type Big struct {
    Data [1024]byte
    N    int
}

func sumIndex(xs []Big) int {
    total := 0
    for i := range xs {
        total += xs[i].N // no copy of Big
    }
    return total
}

func main() {
    fmt.Println(sumIndex([]Big{{N: 1}, {N: 2}}))
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `for _, v := range xs` copies each element — costly for large `T`.
- `for i := range xs` uses index — accesses slice without copying element.
- Alternative: `[]*Big` — range copies pointer (word size), not struct.
- Small ints/strings — copy cost negligible; optimize when struct is large or hot loop.

## Q&A

**Q: `for i, v := range`?**  
A: Still copies `v` each iteration — use index only for large types.

**Q: Modify during range?**  
A: `xs[i].N++` works; `v.N++` does not update slice.

**Q: Complexity?**  
A: O(n); constant factor differs by struct size.

**Q: Benchmark?**  
A: Compare value-range vs index-range on 1KB struct.

**Q: Interview one-liner?**  
A: "Range copies values; index or pointer slice for large elements."
