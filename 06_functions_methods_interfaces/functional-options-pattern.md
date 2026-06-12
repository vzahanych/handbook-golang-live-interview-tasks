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

type Server struct { Addr string; Timeout time.Duration }
type Option func(*Server)

func WithAddr(addr string) Option { return func(s *Server) { s.Addr = addr } }
func WithTimeout(d time.Duration) Option { return func(s *Server) { s.Timeout = d } }

func NewServer(opts ...Option) *Server {
    s := &Server{Addr: ":8080", Timeout: 5 * time.Second}
    for _, opt := range opts { opt(s) }
    return s
}

func main() { fmt.Printf("%+v\n", NewServer(WithAddr(":9090"))) }
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
