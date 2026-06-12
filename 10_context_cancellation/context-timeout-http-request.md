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
    "net/http"
    "time"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://example.com", nil)
    resp, err := http.DefaultClient.Do(req)
    if err != nil { fmt.Println("error:", err); return }
    defer resp.Body.Close()
    fmt.Println(resp.StatusCode)
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- None specific; discuss edge cases and complexity.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
