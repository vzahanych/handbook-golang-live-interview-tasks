# json api handler

## Live interview task
Decode JSON request body and encode JSON response.

## Concepts covered
- encoding/json
- http handlers
- Content-Type

## Candidate solution

```go
package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Request struct {
    Name string `json:"name"`
}

type Response struct {
    Message string `json:"message"`
}

func greet(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }
    defer r.Body.Close()

    var req Request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(Response{Message: "hello " + req.Name})
}

func main() {
    http.HandleFunc("/greet", greet)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Run

```bash
go run .
curl -X POST localhost:8080/greet -d '{"name":"Ada"}' -H 'Content-Type: application/json'
```

## Interview notes / pitfalls
- Set `Content-Type` before `Write` — headers sent on first write.
- `json.Decoder` streams — good for large bodies; `json.Unmarshal` for small `[]byte`.
- Unknown JSON fields ignored by default — `DisallowUnknownFields()` for strict APIs.
- Limit body size: `http.MaxBytesReader(w, r.Body, max)`.

## Q&A

**Q: Encoder error after partial write?**  
A: Hard to change status — encode to buffer first or set status before write.

**Q: Pretty JSON?**  
A: `json.MarshalIndent` — not for hot APIs.

**Q: Validation?**  
A: Check `req.Name != ""` — return 400 with problem JSON.

**Q: `omitempty`?**  
A: Skip zero values in output structs.

**Q: Security?**  
A: Size limits, don't reflect raw errors to clients in prod.
