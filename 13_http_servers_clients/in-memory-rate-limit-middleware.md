# in memory rate limit middleware

## Live interview task
Implement a simple per-process token-bucket-like rate limiter middleware.

## Concepts covered
- middleware
- mutex
- rate limiting

## Candidate solution

```go
package main

import (
    "net/http"
    "sync"
    "time"
)

type limiter struct { mu sync.Mutex; next time.Time; delay time.Duration }
func (l *limiter) allow() bool { l.mu.Lock(); defer l.mu.Unlock(); now := time.Now(); if now.Before(l.next) { return false }; l.next = now.Add(l.delay); return true }

func limit(l *limiter, next http.Handler) http.Handler { return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){ if !l.allow() { http.Error(w,"too many requests",429); return }; next.ServeHTTP(w,r) }) }

func main() { l := &limiter{delay: time.Second}; http.ListenAndServe(":8080", limit(l, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){ w.Write([]byte("ok")) }))) }
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
