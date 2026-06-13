# httptest server client

## Live interview task
Use `httptest.Server` to test HTTP client code without real network.

## Concepts covered
- httptest
- integration-style unit test
- table-driven tests

## Candidate solution

```go
package main

import (
    "fmt"
    "io"
    "net/http"
    "net/http/httptest"
)

func fetch(client *http.Client, url string) (string, error) {
    resp, err := client.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    b, err := io.ReadAll(resp.Body)
    return string(b), err
}

func main() {
    srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "ok")
    }))
    defer srv.Close()

    body, err := fetch(srv.Client(), srv.URL)
    fmt.Println(body, err)
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `srv.URL` is full URL with random port — use `srv.Client()` for correct transport.
- `defer srv.Close()` — shuts down test server.
- `httptest.NewRequest` + `httptest.ResponseRecorder` — test handlers without network.
- `srv.Listener` for custom client if needed.

## Q&A

**Q: Test handler only?**  
A: `req := httptest.NewRequest("GET", "/", nil); rec := httptest.NewRecorder(); handler(rec, req)`.

**Q: TLS test?**  
A: `httptest.NewTLSServer` + `client.Transport` from `srv.Client()`.

**Q: Assert status?**  
A: `rec.Code == http.StatusOK`.

**Q: vs mocking Transport?**  
A: httptest is integration-light — real HTTP stack, local only.

**Q: Parallel tests?**  
A: Each test own `NewServer` — avoid shared global.
