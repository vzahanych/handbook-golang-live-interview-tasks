# waitgroup basic worker start

## Live interview task
Run several goroutines and wait for all of them.

## Concepts covered
- goroutines
- sync.WaitGroup
- loop variables

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
        i := i
        wg.Add(1)
        go func() { defer wg.Done(); fmt.Println("worker", i) }()
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
