# defer modifies named result

## Live interview task
Show how a deferred function can observe and modify named return values.

## Concepts covered
- defer
- named results
- naked return

## Candidate solution

```go
package main

import "fmt"

func compute() (n int) {
    defer func() { n *= 2 }()
    return 21
}

func main() {
    fmt.Println(compute()) // 42
}
```

## Run

```bash
go run .
```

## Expected output

```
42
```

## Interview notes / pitfalls
- `return 21` on named result `n` sets `n=21`, **then** defers run, then function returns `n`.
- Common pattern: `defer func() { if err != nil { ... } }()` with named `(err error)`.
- Defer can change named return; cannot change unnamed return value the same way.
- `return` with expression on named results still assigns before defers.

## Q&A

**Q: Why named returns?**  
A: Defer cleanup modifying `err`; documentation of return values in signature.

**Q: Downside?**  
A: Easy to shadow `err` with `:=` — named return trap in nested blocks.

**Q: `recover` + named `err`?**  
A: `defer func() { if r := recover(); r != nil { err = fmt.Errorf(...) } }()`.

**Q: Performance?**  
A: Named results may escape to heap — minor; clarity trade-off.

**Q: Interview output?**  
A: Always 42 — defers multiply after `return 21`.
