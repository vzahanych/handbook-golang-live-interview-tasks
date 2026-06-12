# httptest server client

## Live interview task
Use httptest.Server for testing client code without external network calls.

## Concepts covered
- httptest
- http client

## Candidate solution

```go
package main

import (
    "fmt"
    "io"
    "net/http"
    "net/http/httptest"
)

func main() {
    srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){ w.Write([]byte("ok")) }))
    defer srv.Close()
    resp, _ := http.Get(srv.URL)
    defer resp.Body.Close()
    b, _ := io.ReadAll(resp.Body)
    fmt.Println(string(b))
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
