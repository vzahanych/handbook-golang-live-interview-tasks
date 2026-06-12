# generic channel fan in

## Live interview task
Write a generic fan-in function for multiple receive-only channels.

## Concepts covered
- generics
- channels
- fan-in

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

func FanIn[T any](chs ...<-chan T) <-chan T {
    out := make(chan T)
    var wg sync.WaitGroup
    wg.Add(len(chs))
    for _, ch := range chs {
        go func(c <-chan T) { defer wg.Done(); for v := range c { out <- v } }(ch)
    }
    go func() { wg.Wait(); close(out) }()
    return out
}

func main() {
    a, b := make(chan int, 1), make(chan int, 1)
    a <- 1; close(a); b <- 2; close(b)
    for v := range FanIn(a,b) { fmt.Println(v) }
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
