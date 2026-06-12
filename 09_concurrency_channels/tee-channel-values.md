# tee channel values

## Live interview task
Duplicate each input value to two output channels.

## Concepts covered
- select
- nil channels
- generics

## Candidate solution

```go
package main

import "fmt"

func tee[T any](in <-chan T) (<-chan T, <-chan T) {
    a, b := make(chan T), make(chan T)
    go func() {
        defer close(a); defer close(b)
        for v := range in {
            out1, out2 := a, b
            for i := 0; i < 2; i++ {
                select {
                case out1 <- v: out1 = nil
                case out2 <- v: out2 = nil
                }
            }
        }
    }()
    return a, b
}

func main() { in := make(chan int, 1); in <- 7; close(in); a, b := tee(in); fmt.Println(<-a, <-b) }
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
