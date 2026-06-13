# loop variable closure capture

## Live interview task
Explain and fix the classic loop-variable capture bug when spawning goroutines or storing pointers in a slice.

## Concepts covered
- for loops
- closures
- goroutines
- Go 1.22 per-iteration semantics

## Valid version (Go 1.22+, current Go)

On any modern Go (1.22 and later, including today's 1.26.x) loop variables are per-iteration, so the straightforward code is correct as written — it prints `1 2 3` (order nondeterministic):

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    vals := []int{1, 2, 3}
    var wg sync.WaitGroup
    for _, v := range vals {
        wg.Add(1)
        go func() {
            defer wg.Done()
            fmt.Println(v) // per-iteration v — prints 1, 2, 3 in some order
        }()
    }
    wg.Wait()
}
```

Note: the semantics is selected by the `go` directive in `go.mod`, not the toolchain — a module declaring `go 1.21` or lower still gets the old shared-variable behavior even when built with Go 1.26.

## Legacy bug (pre-Go 1.22 semantics)

Before Go 1.22 (or in a module with `go < 1.22` in `go.mod`), all iterations shared one `v`, so the same code often printed `3, 3, 3`. The classic fix was a per-iteration copy:

```go
for _, v := range vals {
    v := v // per-iteration copy — required before Go 1.22
    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println(v)
    }()
}
```

Interviewers may still ask about legacy code and `go vet` / `loopclosure` analyzer.

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Same bug without goroutines (pre-1.22): `funcs := []func(){ func() { fmt.Println(i) } }` — all closures share one `i`.
- Pointer trap (pre-1.22): `items = append(items, &item)` in a range loop — all pointers point at the same variable.
- Passing loop var as a parameter avoids shadowing: `go func(v int) { ... }(v)`.
- `t.Parallel()` in table tests had the same bug — fixed in Go 1.22.

## Q&A

**Q: Why did Go reuse one variable?**  
A: Historical implementation simplicity; fixing it broke some code that accidentally relied on the old behavior.

**Q: How does `go vet` help?**  
A: The `loopclosure` analyzer reports captures of loop variables by closures/goroutines.

**Q: Is `for i := 0; i < n; i++` affected?**  
A: Yes — `i` was per-loop before 1.22; now per-iteration in 1.22+.

**Q: Complexity?**  
A: O(n) goroutines; the fix is O(1) extra per iteration (one int copy).

**Q: Reference?**  
A: [Go 1.22 loop variable change](https://go.dev/blog/loopvar-preview)
