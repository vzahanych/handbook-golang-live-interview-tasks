# first result wins

## Live interview task
Start several workers and use the first successful result.

## Concepts covered
- select
- buffered result channel
- racing requests

## Candidate solution

```go
package main

import (
    "fmt"
    "time"
)

func query(name string, d time.Duration) <-chan string {
    ch := make(chan string, 1)
    go func(){ time.Sleep(d); ch <- name }()
    return ch
}

func main() {
    select {
    case v := <-query("fast", 10*time.Millisecond): fmt.Println(v)
    case v := <-query("slow", 100*time.Millisecond): fmt.Println(v)
    }
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
