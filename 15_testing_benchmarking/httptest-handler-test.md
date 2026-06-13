# httptest handler test

## Live interview task
Test an HTTP handler with `httptest.ResponseRecorder`.

## Concepts covered
- httptest
- handler unit tests
- status and body assertions

## Candidate solution

```go
package api

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func health(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }
    w.WriteHeader(http.StatusOK)
    _, _ = w.Write([]byte("ok"))
}

func TestHealth(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/health", nil)
    rec := httptest.NewRecorder()

    health(rec, req)

    if rec.Code != http.StatusOK {
        t.Fatalf("status=%d", rec.Code)
    }
    if got := rec.Body.String(); got != "ok" {
        t.Fatalf("body=%q", got)
    }
}
```

## Run

```bash
go test
```

## Interview notes / pitfalls
- `NewRequest` + `Recorder` — no network, fast unit test.
- Check `rec.Header()` for `Content-Type`, cookies.
- Table-drive method/path/body/status like other tests.
- Middleware: wrap handler with middleware under test, same recorder pattern.

## Q&A

**Q: POST with body?**  
A: `httptest.NewRequest("POST", "/", strings.NewReader(body))`.

**Q: Context on request?**  
A: `req = req.WithContext(ctx)`.

**Q: vs `NewServer`?**  
A: Recorder = handler only; Server = full stack + client.

**Q: HandlerFunc adapter?**  
A: `http.HandlerFunc(health).ServeHTTP(rec, req)`.

**Q: Complexity?**  
A: O(1) per test case.
