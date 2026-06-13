# middleware logging status recorder

## Live interview task
Write logging middleware that records the response status code.

## Concepts covered
- middleware pattern
- wrapping ResponseWriter
- embedding

## Candidate solution

```go
package main

import (
    "log"
    "net/http"
)

type recorder struct {
    http.ResponseWriter
    status int
}

func (r *recorder) WriteHeader(code int) {
    r.status = code
    r.ResponseWriter.WriteHeader(code)
}

func logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        rec := &recorder{ResponseWriter: w, status: http.StatusOK}
        next.ServeHTTP(rec, r)
        log.Println(r.Method, r.URL.Path, rec.status)
    })
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("ok"))
    })
    log.Fatal(http.ListenAndServe(":8080", logging(mux)))
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Default status **200** if `WriteHeader` never called — recorder assumes 200.
- Embedding `http.ResponseWriter` — must forward `Write`, `WriteHeader`, `Header`.
- Optional: implement `http.Flusher`, `http.Hijacker` if downstream needs them — type assert embed.
- Middleware signature: `func(http.Handler) http.Handler`.

## Q&A

**Q: Capture bytes written?**  
A: Override `Write` to count len(p).

**Q: Order of middleware?**  
A: `logging(auth(mux))` — outer runs first on way in, last on way out.

**Q: `chi`/`gorilla`?**  
A: Same pattern — `Use(middleware)`.

**Q: Panic in handler?**  
A: Recover middleware logs 500 — separate concern.

**Q: Complexity?**  
A: O(1) per request.
