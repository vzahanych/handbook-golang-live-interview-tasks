# functional options pattern

## Live interview task
Build a server config with functional options.

## Concepts covered
- variadic functions
- function types
- options pattern

## Candidate solution

```go
package main

import (
    "fmt"
    "time"
)

type Server struct {
    Addr    string
    Timeout time.Duration
}

type Option func(*Server)

func WithAddr(addr string) Option {
    return func(s *Server) { s.Addr = addr }
}

func WithTimeout(d time.Duration) Option {
    return func(s *Server) { s.Timeout = d }
}

func NewServer(opts ...Option) *Server {
    s := &Server{Addr: ":8080", Timeout: 5 * time.Second}
    for _, opt := range opts {
        opt(s)
    }
    return s
}

func main() {
    fmt.Printf("%+v\n", NewServer(WithAddr(":9090")))
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Defaults set in `NewServer` before applying options — options override.
- `Option` is `func(*T)` — easy to add fields without breaking callers.
- Validate in `NewServer` after options applied — return `(*Server, error)` if invalid.
- Alternative: builder struct with chained methods — more verbose, clearer for many required fields.

## Q&A

**Q: Why not many constructor params?**  
A: Avoids `NewServer(addr, timeout, tls, ...)` explosion; optional params stay optional.

**Q: Order of options?**  
A: Last wins if two options set same field — document or error on conflict.

**Q: Testing?**  
A: Pass `WithTimeout(0)` in tests; inject fake options.

**Q: Used in stdlib?**  
A: `grpc.DialOption`, `client.Option` patterns in many Go libraries.

**Q: Complexity?**  
A: O(k) for k options applied.
