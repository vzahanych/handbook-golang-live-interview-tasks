# sync cond broadcast

## Live interview task
Coordinate goroutines with sync.Cond.

## Concepts covered
- sync.Cond
- condition loops

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    mu := sync.Mutex{}
    cond := sync.NewCond(&mu)
    ready := false
    go func(){ mu.Lock(); for !ready { cond.Wait() }; fmt.Println("go"); mu.Unlock() }()
    mu.Lock(); ready = true; cond.Broadcast(); mu.Unlock()
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Always wait in a loop because conditions can change before a goroutine wakes.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
