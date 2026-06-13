# bounds check elimination hint

## Live interview task
Write tight loops where the compiler can eliminate redundant bounds checks.

## Concepts covered
- bounds check elimination (BCE)
- compiler optimization
- slice length hoisting

## Candidate solution

```go
package main

import "fmt"

func sum4(s []int) int {
    if len(s) < 4 {
        panic("need 4 elements")
    }
    // Compiler may eliminate per-index checks after len guard
    return s[0] + s[1] + s[2] + s[3]
}

func sumLoop(s []int) int {
    total := 0
    for i := 0; i < len(s); i++ {
        total += s[i]
    }
    return total
}

func main() {
    fmt.Println(sum4([]int{1, 2, 3, 4}))
}
```

## Run

```bash
go run .
go build -gcflags=-m=2 . 2>&1 | head -30
```

## Interview notes / pitfalls
- Go inserts bounds checks on slice access — BCE removes redundant ones in loop.
- Hoist `n := len(s)` and loop `i < n` — classic pattern.
- Early `if len(s) < k { return }` proves safety for fixed indexes.
- Micro-optimization — profile first; clarity over BCE tricks in app code.

## Q&A

**Q: Inspect escapes/BCE?**  
A: `go build -gcflags=-m` shows compiler decisions.

**Q: `for range` vs index?**  
A: Both optimized similarly in modern Go — benchmark if critical.

**Q: When it matters?**  
A: Tight inner loops on millions of iterations.

**Q: Unsafe skip checks?**  
A: `unsafe` — avoid unless expert; not interview default.

**Q: Complexity?**  
A: Same O(n); constant factor may shrink.
