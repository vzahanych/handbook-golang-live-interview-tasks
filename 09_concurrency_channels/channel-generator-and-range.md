# channel generator and range

## Live interview task
Create a generator function that sends values then closes the channel.

## Concepts covered
- channels
- close
- range over channel

## Candidate solution

```go
package main

import "fmt"

func gen(n int) <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        for i := 0; i < n; i++ { ch <- i }
    }()
    return ch
}

func main() { for v := range gen(5) { fmt.Println(v) } }
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
