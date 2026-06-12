# graceful shutdown server

## Live interview task
Run an HTTP server and gracefully shut it down on SIGINT/SIGTERM.

## Concepts covered
- http.Server
- graceful shutdown
- signals

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
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){ w.Write([]byte("ok")) })
    srv := &http.Server{Addr: ":8080", Handler: mux}
    go func(){ if err := srv.ListenAndServe(); err != http.ErrServerClosed { log.Fatal(err) } }()
    ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
    defer stop()
    <-ctx.Done()
    shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    log.Println(srv.Shutdown(shutdownCtx))
}
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
