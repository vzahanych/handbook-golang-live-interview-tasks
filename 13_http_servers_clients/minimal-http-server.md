# minimal http server

## Live interview task
Build a minimal HTTP server with one handler.

## Concepts covered
- net/http
- ListenAndServe
- HandlerFunc

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
    log.Println("listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Run

```bash
go run .
curl localhost:8080
```

## Interview notes / pitfalls
- `nil` handler uses `DefaultServeMux` — `http.HandleFunc` registers globally.
- Each request runs in its own goroutine — handlers must be concurrent-safe.
- `ListenAndServe` blocks; errors only on shutdown or fatal bind failure.
- Production: use explicit `http.Server` with timeouts (see graceful-shutdown).

## Q&A

**Q: `Handle` vs `HandleFunc`?**  
A: `HandleFunc` wraps function as `Handler`; `Handle` takes `http.Handler`.

**Q: Default mux thread-safe?**  
A: Registration at init/main once; concurrent Serve OK.

**Q: Path matching?**  
A: `ServeMux` longest prefix match; Go 1.22+ method-aware patterns.

**Q: Complexity?**  
A: O(1) per request dispatch overhead.

**Q: Next step?**  
A: `mux := http.NewServeMux()` — isolated router, no global state.
