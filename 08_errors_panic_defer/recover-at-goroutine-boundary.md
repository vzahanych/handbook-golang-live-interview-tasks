# recover at goroutine boundary

## Live interview task
Recover from a panic at a goroutine boundary and report it as an error.

## Concepts covered
- panic
- recover
- goroutines

## Candidate solution

```go
package main

import "fmt"

func safeGo(fn func()) <-chan error {
    done := make(chan error, 1)
    go func() {
        defer func() {
            if r := recover(); r != nil {
                done <- fmt.Errorf("panic: %v", r)
            } else {
                done <- nil
            }
        }()
        fn()
    }()
    return done
}

func main() {
    fmt.Println(<-safeGo(func() { panic("boom") }))
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `recover` only works inside **deferred** function in the **same goroutine** that panicked.
- Panic in child goroutine does **not** crash parent — unless you propagate via channel.
- `defer recover()` alone does **not** work — must be `defer func() { recover() }()`.
- Production: log stack with `debug.Stack()`; consider re-panic after log in `main` only.

## Q&A

**Q: Why recover at goroutine boundary?**  
A: One bad task should not kill entire server — HTTP handler, worker pool pattern.

**Q: Return `error` vs `panic`?**  
A: Libraries return errors; `panic` for programmer bugs or `Must` helpers.

**Q: `recover()` return value?**  
A: Value passed to `panic()` — often `string` or `error`.

**Q: Test panics?**  
A: `defer func() { if r := recover(); r == nil { t.Fatal } }()` in test.

**Q: Complexity?**  
A: O(1) — control flow, not algorithmic.
