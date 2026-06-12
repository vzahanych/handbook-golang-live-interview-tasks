# method set pointer vs value receiver

## Live interview task
Show which methods belong to T and *T method sets.

## Concepts covered
- method sets
- pointer receiver
- addressability

## Candidate solution

```go
package main

import "fmt"

type Counter int
func (c Counter) Value() int { return int(c) }
func (c *Counter) Inc() { *c++ }

type Valuer interface{ Value() int }
type Incer interface{ Inc() }

func main() {
    var c Counter
    var _ Valuer = c
    var _ Valuer = &c
    // var _ Incer = c  // compile error: Inc has pointer receiver
    var _ Incer = &c
    c.Inc() // ok because c is addressable; compiler uses (&c).Inc()
    fmt.Println(c.Value())
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
