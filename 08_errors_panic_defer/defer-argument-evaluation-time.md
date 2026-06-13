# defer argument evaluation time

## Live interview task
Show that deferred call arguments are evaluated when `defer` executes, not when the deferred call runs.

## Concepts covered
- defer
- argument evaluation

## Candidate solution

```go
package main

import "fmt"

func main() {
    x := 1
    defer fmt.Println("deferred", x)
    x = 2
    fmt.Println("normal", x)
}
```

## Run

```bash
go run .
```

## Expected output

```
normal 2
deferred 1
```

## Interview notes / pitfalls
- `defer fmt.Println(x)` captures **value of x at defer line** (1), not at function exit.
- To defer with latest value: `defer func() { fmt.Println(x) }()` — closure reads `x` at run time (2).
- Same for `defer file.Close()` — file variable must be non-nil; often `defer func() { _ = f.Close() }()` after err check.
- `defer mu.Unlock()` — function call deferred, not lock state — correct pattern.

## Q&A

**Q: Closure defer cost?**  
A: Tiny extra alloc — worth it for dynamic values.

**Q: `defer os.Remove(name)` with changing name?**  
A: `name` evaluated at defer — if `name` reassigned, wrong file removed — use closure.

**Q: Interview trick?**  
A: Stack multiple defers with loop variable — see defer-in-loop trap.

**Q: LIFO order still applies?**  
A: Yes — argument eval at registration, calls at exit in reverse order.

**Q: Related?**  
A: `go fmt.Println(x)` — `x` evaluated at `go` statement too.
