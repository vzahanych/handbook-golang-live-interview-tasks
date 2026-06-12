# json api handler

## Live interview task
Decode JSON request body and encode JSON response.

## Concepts covered
- json
- http handlers
- struct tags

## Candidate solution

```go
package main

import (
    "encoding/json"
    "net/http"
)

type Request struct{ Name string `json:"name"` }
type Response struct{ Message string `json:"message"` }

func greet(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()
    var req Request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil { http.Error(w, err.Error(), http.StatusBadRequest); return }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(Response{Message: "hello " + req.Name})
}

func main() { http.HandleFunc("/greet", greet); http.ListenAndServe(":8080", nil) }
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
