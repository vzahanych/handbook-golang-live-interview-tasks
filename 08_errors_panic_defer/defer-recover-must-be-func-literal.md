# defer recover must be func literal

## Live interview task
Show that `defer recover()` does not catch panics — only `defer func() { recover() }()`.

## Concepts covered
- recover rules
- defer
- spec semantics

## Broken version (does not recover)

```go
func broken() {
    defer recover() // recover NOT called by deferred function — useless
    panic("boom")
}
```

## Candidate solution

```go
package main

import "fmt"

func safe() (caught bool) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered:", r)
            caught = true
        }
    }()
    panic("boom")
}

func main() {
    fmt.Println("safe returned:", safe()) // true
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Spec: `recover` effective only when called **directly** by a deferred function.
- `defer recover()` defers the **result** of calling `recover()` at defer time — returns nil, no-op at panic.
- Same class of bug as `go recover()` — must be inside goroutine's deferred func.
- golang/go#14273 — working as intended.

## Q&A

**Q: Why spec strict?**  
A: Clear stack frame for panic handling; discourage sloppy `defer recover()`.

**Q: Check recover result?**  
A: `if r := recover(); r != nil` — nil if no panic.

**Q: Re-panic after log?**  
A: `defer func() { if r := recover(); r != nil { log(); panic(r) } }()` — middleware pattern.

**Q: Production handler?**  
A: HTTP server recovers per request in middleware deferred func.

**Q: One-liner?**  
A: "recover must live inside defer func(), never defer recover()."
