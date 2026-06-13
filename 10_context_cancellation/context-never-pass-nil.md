# context never pass nil

## Live interview task
Explain why `context` must not be nil and what to use instead.

## Concepts covered
- context contract
- context.TODO
- context.Background

## Candidate solution

```go
package main

import (
    "context"
    "fmt"
)

func work(ctx context.Context) error {
    if ctx == nil {
        return fmt.Errorf("nil context")
    }
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
        return nil
    }
}

func main() {
    _ = work(context.Background()) // root for production top-level
    _ = work(context.TODO())       // placeholder when unsure of parent
    // _ = work(nil) // undefined behavior — many APIs panic or hang
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Passing `nil` context: `ctx.Done()` panics — "nil context" in many stdlib paths.
- `context.Background()` — non-nil root, never canceled.
- `context.TODO()` — non-nil placeholder during refactor; same behavior as Background for cancel.
- Libraries should document `ctx` required; defensively check nil at public API if needed.

## Q&A

**Q: Background vs TODO?**  
A: Semantics for readers — TODO means temporary/unclear parent.

**Q: Test context?**  
A: `context.WithCancel` in test, defer cancel.

**Q: HTTP handler?**  
A: Always `r.Context()` — non-nil when request active.

**Q: First param rule?**  
A: `func Foo(ctx context.Context, ...)` — Go convention since context package.

**Q: One-liner?**  
A: "If you don't have a context, use Background — never nil."
