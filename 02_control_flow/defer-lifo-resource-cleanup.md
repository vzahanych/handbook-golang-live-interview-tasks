# defer lifo resource cleanup

## Live interview task
Show that defers run in last-in-first-out order.

## Concepts covered
- defer
- LIFO order

## Candidate solution

```go
package main

import "fmt"

func work() {
    defer fmt.Println("close file")
    defer fmt.Println("flush buffer")
    defer fmt.Println("unlock mutex")
    fmt.Println("do work")
}

func main() { work() }
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
