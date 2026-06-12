# request scoped context value middleware

## Live interview task
Attach a request ID to context in middleware.

## Concepts covered
- http middleware
- context values

## Candidate solution

```go
package main

import (
    "context"
    "fmt"
    "net/http"
    "time"
)

type key string
const requestID key = "requestID"

func withRequestID(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        id := fmt.Sprintf("%d", time.Now().UnixNano())
        ctx := context.WithValue(r.Context(), requestID, id)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func main() { http.ListenAndServe(":8080", withRequestID(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){ fmt.Fprintln(w, r.Context().Value(requestID)) }))) }
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
