# in memory rate limit middleware

## Live interview task
Implement simple per-process rate limiting middleware.

## Concepts covered
- middleware
- mutex
- 429 Too Many Requests

## Candidate solution

```go
package main

import (
    "log"
    "net/http"
    "sync"
    "time"
)

type limiter struct {
    mu    sync.Mutex
    next  time.Time
    delay time.Duration
}

func (l *limiter) allow() bool {
    l.mu.Lock()
    defer l.mu.Unlock()
    now := time.Now()
    if now.Before(l.next) {
        return false
    }
    l.next = now.Add(l.delay)
    return true
}

func rateLimit(l *limiter, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !l.allow() {
            http.Error(w, "too many requests", http.StatusTooManyRequests)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func main() {
    l := &limiter{delay: time.Second}
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("ok"))
    })
    log.Fatal(http.ListenAndServe(":8080", rateLimit(l, mux)))
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Demo is **global** one-request-per-delay — not per IP.
- Production: token bucket per client key (IP, API key), Redis for distributed limit.
- Return `Retry-After` header with 429.
- Mutex serializes `allow` — fine for demo; shard limiters by key for scale.

## Q&A

**Q: Per-IP limiter?**  
A: `map[string]*limiter` with mutex or sync.Map + cleanup.

**Q: `golang.org/x/time/rate`?**  
A: Token bucket — `rate.Limiter` with `Allow()`.

**Q: Distributed?**  
A: Redis INCR + TTL or sliding window Lua script.

**Q: Bypass for health?**  
A: Skip limit in middleware if `r.URL.Path == "/healthz"`.

**Q: Complexity?**  
A: O(1) per request check.
