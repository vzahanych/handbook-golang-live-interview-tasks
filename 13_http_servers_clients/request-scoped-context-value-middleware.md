# request scoped context value middleware

## Live interview task
Attach a request ID to context in middleware and read it in handlers.

## Concepts covered
- middleware
- context.WithValue
- r.WithContext

## Candidate solution

```go
package main

import (
    "context"
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "log"
    "net/http"
)

type ctxKey int

const requestIDKey ctxKey = iota

func withRequestID(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        id := newRequestID()
        w.Header().Set("X-Request-ID", id)
        ctx := context.WithValue(r.Context(), requestIDKey, id)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func newRequestID() string {
    var b [8]byte
    _, _ = rand.Read(b[:])
    return hex.EncodeToString(b[:])
}

func handler(w http.ResponseWriter, r *http.Request) {
    id, _ := r.Context().Value(requestIDKey).(string)
    log.Println("request", id)
    fmt.Fprintln(w, id)
}

func main() {
    http.Handle("/", withRequestID(http.HandlerFunc(handler)))
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Run

```bash
go run .
curl -i localhost:8080/
```

## Interview notes / pitfalls
- Use **unexported key type** — avoid context key collisions.
- Pass `r.WithContext(ctx)` to next handler — not the old request.
- `r.Context()` canceled when client disconnects — respect in long handlers.
- Prefer structured logging with request ID in every log line.

## Q&A

**Q: Accept client X-Request-ID?**  
A: Optional — validate format or generate if missing.

**Q: vs global variable?**  
A: Context is per-request — safe concurrent.

**Q: OpenTelemetry?**  
A: Trace ID in context — industry standard extension.

**Q: Middleware order?**  
A: Request ID early — logging/auth can use it.

**Q: Complexity?**  
A: O(1) per request.
