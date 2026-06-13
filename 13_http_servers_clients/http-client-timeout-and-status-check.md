# http client timeout and status check

## Live interview task
Create an HTTP client with timeout and verify status codes.

## Concepts covered
- http.Client
- timeouts
- status codes

## Candidate solution

```go
package main

import (
    "fmt"
    "io"
    "net/http"
    "time"
)

func fetch(url string) ([]byte, error) {
    client := &http.Client{Timeout: 3 * time.Second}
    resp, err := client.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode < 200 || resp.StatusCode >= 300 {
        return nil, fmt.Errorf("bad status: %s", resp.Status)
    }
    return io.ReadAll(resp.Body)
}

func main() {
    b, err := fetch("https://example.com")
    fmt.Println(len(b), err)
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `Client.Timeout` covers entire exchange — dial, TLS, headers, body read.
- Finer control: `Transport` with `DialContext`, `TLSHandshakeTimeout`, `ResponseHeaderTimeout`.
- Always close body — leak connections if not.
- Check status before read — 500 may still have body to drain.

## Q&A

**Q: Reuse client?**  
A: Yes — `var client = &http.Client{...}` — connection pooling via Transport.

**Q: 404 vs error?**  
A: 404 returns nil transport err — must check `StatusCode`.

**Q: Context?**  
A: `NewRequestWithContext` preferred over Client.Timeout alone.

**Q: Redirects?**  
A: Default follows 3xx — cap with custom CheckRedirect.

**Q: Production?**  
A: Retry idempotent GET with backoff; circuit breaker for deps.
