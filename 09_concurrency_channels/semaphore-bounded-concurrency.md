# semaphore bounded concurrency

## Live interview task
Limit concurrent work with a buffered channel used as a semaphore.

## Concepts covered
- buffered channels
- semaphore
- bounded concurrency

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    sem := make(chan struct{}, 2)
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        i := i
        wg.Add(1)
        go func() {
            defer wg.Done()
            sem <- struct{}{}
            defer func(){ <-sem }()
            time.Sleep(10 * time.Millisecond)
            fmt.Println("done", i)
        }()
    }
    wg.Wait()
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
