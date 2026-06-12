# exported identifiers and package api

## Live interview task
Demonstrate exported versus unexported identifiers in a package API.

## Concepts covered
- exported identifiers
- packages
- methods

## Candidate solution

```go
// file: counter/counter.go
package counter

type Counter struct { n int } // Counter is exported; n is internal.

func New() *Counter { return &Counter{} }
func (c *Counter) Inc() { c.n++ }
func (c *Counter) Value() int { return c.n }

// file: main.go
package main

import (
    "fmt"
    "example/counter"
)

func main() {
    c := counter.New()
    c.Inc()
    fmt.Println(c.Value())
}
```

## Run

```bash
go mod init example && go run .
```

## Interview notes / pitfalls
- None specific; discuss edge cases and complexity.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
