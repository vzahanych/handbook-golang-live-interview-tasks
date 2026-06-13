# context timeout http request

## Live interview task
Make an HTTP request with a context timeout.

## Concepts covered
- context
- http.Client
- timeouts

## Candidate solution

```go
package main

import (
    "context"
    "fmt"
    "io"
    "net/http"
    "time"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://example.com", nil)
    if err != nil {
        panic(err)
    }

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println("error:", err) // context.DeadlineExceeded if timeout
        return
    }
    defer resp.Body.Close()
    io.Copy(io.Discard, resp.Body) // drain body for connection reuse
    fmt.Println(resp.StatusCode)
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Always `defer cancel()` — releases timer resources even on success.
- `Client.Timeout` vs request context — both can apply; prefer context per request in servers.
- `DefaultClient` has no timeout — dangerous in production; use custom `http.Client`.
- Timeout cancels request; error is often `context.DeadlineExceeded`.

## Q&A

**Q: `WithTimeout` vs `WithDeadline`?**  
A: Timeout = duration from now; deadline = absolute `time.Time`.

**Q: Body not read?**  
A: Connection may not reuse — drain or close body.

**Q: Retry on timeout?**  
A: New context per attempt; exponential backoff; idempotent methods only.

**Q: Server side?**  
A: `r.Context()` from `http.Request` — cancel when client disconnects.

**Q: Complexity?**  
A: O(1) setup; I/O bound for request.
