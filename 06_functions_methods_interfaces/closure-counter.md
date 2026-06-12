# closure counter

## Live interview task
Return a closure that captures and mutates local state.

## Concepts covered
- function literals
- closures

## Candidate solution

```go
package main

import "fmt"

func counter() func() int {
    n := 0
    return func() int { n++; return n }
}

func main() {
    next := counter()
    fmt.Println(next(), next(), next())
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
