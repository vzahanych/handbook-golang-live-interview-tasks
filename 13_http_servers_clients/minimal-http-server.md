# minimal http server

## Live interview task
Build a minimal HTTP server with one handler.

## Concepts covered
- net/http
- handlers

## Candidate solution

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "hello")
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Run

```bash
go run . && curl localhost:8080
```

## Interview notes / pitfalls
- None specific; discuss edge cases and complexity.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
