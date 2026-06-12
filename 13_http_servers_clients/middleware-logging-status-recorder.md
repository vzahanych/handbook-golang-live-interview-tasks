# middleware logging status recorder

## Live interview task
Write logging middleware that records response status code.

## Concepts covered
- middleware
- embedding
- http.ResponseWriter

## Candidate solution

```go
package main

import (
    "log"
    "net/http"
)

type recorder struct { http.ResponseWriter; status int }
func (r *recorder) WriteHeader(code int) { r.status = code; r.ResponseWriter.WriteHeader(code) }

func logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        rec := &recorder{ResponseWriter: w, status: http.StatusOK}
        next.ServeHTTP(rec, r)
        log.Println(r.Method, r.URL.Path, rec.status)
    })
}

func main() { http.ListenAndServe(":8080", logging(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){ w.Write([]byte("ok")) }))) }
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
