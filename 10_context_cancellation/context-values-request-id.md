# context values request id

## Live interview task
Pass request-scoped values with `context.WithValue` and retrieve with `ctx.Value`.

## Concepts covered
- context values
- unexported key type
- request-scoped data

## Candidate solution

```go
package main

import (
    "context"
    "fmt"
)

type ctxKey int

const requestIDKey ctxKey = iota

func withRequestID(ctx context.Context, id string) context.Context {
    return context.WithValue(ctx, requestIDKey, id)
}

func requestID(ctx context.Context) (string, bool) {
    v, ok := ctx.Value(requestIDKey).(string)
    return v, ok
}

func log(ctx context.Context, msg string) {
    id, _ := requestID(ctx)
    fmt.Println(id, msg)
}

func main() {
    ctx := withRequestID(context.Background(), "req-123")
    log(ctx, "started")
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Use **unexported custom type** for keys — prevents collisions with other packages' `WithValue`.
- Context values are for **request-scoped** metadata (trace ID, auth), not function parameters.
- Go blog: treat context as read-only propagation — don't pass optional args or DB handles (debated; keep minimal).
- `Value` returns `any` — type assert with comma-ok.

## Q&A

**Q: String key `context.WithValue(ctx, "id", v)`?**  
A: Collision risk across packages — use private type.

**Q: Nil context?**  
A: `context.TODO()` for unclear parent; never pass nil to `WithValue`.

**Q: Chain of middleware?**  
A: Each layer `WithValue` child context — lookup walks up tree.

**Q: Performance?**  
A: Linked list lookup O(depth) — keep chain shallow.

**Q: Alternative?**  
A: Explicit struct param, OpenTelemetry baggage — for cross-cutting observability.
