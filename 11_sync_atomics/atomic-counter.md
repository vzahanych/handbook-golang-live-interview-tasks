# atomic counter

## Live interview task
Use typed atomics for a concurrent counter.

## Concepts covered
- sync/atomic
- typed atomics

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    var n atomic.Int64
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ { wg.Add(1); go func(){ defer wg.Done(); n.Add(1) }() }
    wg.Wait()
    fmt.Println(n.Load())
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
