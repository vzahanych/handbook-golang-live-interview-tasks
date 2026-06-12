# safe channel closing owner

## Live interview task
Show the channel closing rule: the sending owner closes the channel.

## Concepts covered
- close
- channel ownership

## Candidate solution

```go
package main

import "fmt"

func produce(n int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for i := 0; i < n; i++ { out <- i }
    }()
    return out
}

func main() { for v := range produce(3) { fmt.Println(v) } }
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
