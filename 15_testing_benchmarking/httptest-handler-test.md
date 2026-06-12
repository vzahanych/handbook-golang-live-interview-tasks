# httptest handler test

## Live interview task
Test an HTTP handler with httptest.ResponseRecorder.

## Concepts covered
- httptest
- HTTP handler tests

## Candidate solution

```go
package api

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func health(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK); w.Write([]byte("ok")) }

func TestHealth(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/health", nil)
    rec := httptest.NewRecorder()
    health(rec, req)
    if rec.Code != http.StatusOK { t.Fatalf("status=%d", rec.Code) }
    if rec.Body.String() != "ok" { t.Fatalf("body=%q", rec.Body.String()) }
}
```

## Run

```bash
go test
```

## Interview notes / pitfalls
- None specific; discuss edge cases and complexity.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
