# select default nonblocking receive

## Live interview task
Implement non-blocking receive with select and default.

## Concepts covered
- select
- default
- non-blocking receive
- generics

## Candidate solution

```go
package main

import "fmt"

func tryRecv[T any](ch <-chan T) (T, bool) {
    select {
    case v := <-ch:
        return v, true
    default:
        var zero T
        return zero, false
    }
}

func main() {
    ch := make(chan int, 1)
    fmt.Println(tryRecv(ch))
    ch <- 10
    fmt.Println(tryRecv(ch))
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
