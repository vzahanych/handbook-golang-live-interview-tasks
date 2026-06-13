# graceful shutdown server

## Live interview task
Run an HTTP server and gracefully shut it down on SIGINT/SIGTERM.

## Concepts covered
- http.Server
- Shutdown
- signal.NotifyContext

## Candidate solution

```go
package main

import (
    "context"
    "log"
    "net/http"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("ok"))
    })

    srv := &http.Server{
        Addr:    ":8080",
        Handler: mux,
    }

    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatal(err)
        }
    }()

    ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
    defer stop()
    <-ctx.Done()

    shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := srv.Shutdown(shutdownCtx); err != nil {
        log.Println("shutdown:", err)
    }
}
```

## Run

```bash
go run .
# Ctrl+C to trigger shutdown
```

## Interview notes / pitfalls
- `Shutdown` stops accepting new conns, waits for in-flight requests (up to ctx timeout).
- `ListenAndServe` returns `http.ErrServerClosed` after Shutdown — not an error.
- Set `ReadHeaderTimeout`, `IdleTimeout` on Server — slowloris protection.
- `Close()` force-closes — use only when Shutdown times out.

## Q&A

**Q: In-flight request longer than 5s?**  
A: Shutdown ctx expires — may force kill or call `Close`.

**Q: Kubernetes?**  
A: SIGTERM → graceful drain before pod kill.

**Q: Background jobs?**  
A: Separate wg — shutdown hooks after HTTP drain.

**Q: vs `ListenAndServe`?**  
A: `Server` struct required for controlled lifecycle.

**Q: Test?**  
A: Start server in goroutine, call Shutdown from test.
