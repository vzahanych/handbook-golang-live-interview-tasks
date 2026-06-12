# or done channel combinator

## Live interview task
Close an output channel when any input done channel closes.

## Concepts covered
- select
- channel combinators
- cancellation

## Candidate solution

```go
package main

import "fmt"

func or(chs ...<-chan struct{}) <-chan struct{} {
    done := make(chan struct{})
    go func() {
        defer close(done)
        switch len(chs) {
        case 0: return
        case 1: <-chs[0]
        default:
            select { case <-chs[0]: case <-chs[1]: case <-or(chs[2:]...): }
        }
    }()
    return done
}

func main() { a := make(chan struct{}); close(a); <-or(a); fmt.Println("closed") }
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
