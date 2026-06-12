# defer argument evaluation time

## Live interview task
Show that deferred call arguments are evaluated when defer is executed.

## Concepts covered
- defer
- argument evaluation

## Candidate solution

```go
package main

import "fmt"

func main() {
    x := 1
    defer fmt.Println("deferred", x)
    x = 2
    fmt.Println("normal", x)
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
