# defer lifo resource cleanup

## Live interview task
Show that defers run in last-in-first-out order when a function returns.

## Concepts covered
- defer
- LIFO order
- resource cleanup

## Candidate solution

```go
package main

import "fmt"

func work() {
    defer fmt.Println("print third")
    defer fmt.Println("print second")
    defer fmt.Println("print first")
    fmt.Println("do work")
}

func main() { work() }
```

## Run

```bash
go run .
```

## Expected output

```
do work
print first
print second
print third
```

Defers run in LIFO order: the last `defer` registered (`print first`) runs first; the first registered (`print third`) runs last.

## Interview notes / pitfalls
- Defer runs when the **surrounding function** returns (normal return or panic), not when the defer statement is reached.
- Deferred function **arguments are evaluated immediately**; only the call is delayed: `defer fmt.Println(x); x = 2` prints the old `x`.
- `defer` in a loop registers N defers — all run at function exit (often a performance/memory trap). Prefer one defer with a slice of closers or a scoped function.
- `defer recover()` does **not** work — `recover` must be called directly inside a deferred **function literal**.

## Q&A

**Q: Why LIFO for cleanup?**  
A: Mirrors acquisition order: open A, open B → close B, then A (nested resources).

**Q: What is the cost of defer?**  
A: Small overhead (was ~50ns; improved over releases). Hot paths may inline manual cleanup; most code should use defer for clarity.

**Q: Do defers run after `return` evaluates?**  
A: Yes — return values are set, then defers run, then the function actually returns. Named return + defer can modify the named result.

**Q: Edge cases?**  
A: `os.Exit` skips defers; defers in goroutines only run when **that** goroutine's function returns.

**Q: Production pattern?**  
A: `defer f.Close()` immediately after error check on `Open`; combine with `defer func() { _ = f.Close() }()` when Close returns error.
