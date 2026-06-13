# responsewriter writeheader once

## Live interview task
Explain `WriteHeader` may be called only once and default status 200 on first `Write`.

## Concepts covered
- http.ResponseWriter
- status code rules
- middleware traps

## Candidate solution

```go
package main

import (
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.WriteHeader(http.StatusCreated) // 201 — must be before body
    w.Write([]byte("created\n"))

    // w.WriteHeader(http.StatusTeapot) // ignored — already sent
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Run

```bash
go run .
curl -i localhost:8080/
```

## Interview notes / pitfalls
- First `Write` triggers implicit `WriteHeader(200)` if not called — can't set 404 after `Write`.
- `http.Error` calls `WriteHeader` then writes body — convenience helper.
- Middleware calling `WriteHeader` twice — bug; use recorder pattern.
- Header map can be modified until `WriteHeader` called.

## Q&A

**Q: Change status after write?**  
A: Impossible — headers already sent.

**Q: JSON error mid-handler?**  
A: Encode to buffer, then single Write with correct status.

**Q: `Flusher` after WriteHeader?**  
A: OK — flush body chunks.

**Q: httptest `rec.Code`?**  
A: Defaults 200 if Write without WriteHeader.

**Q: Interview one-liner?**  
A: "Set status before any body write."
