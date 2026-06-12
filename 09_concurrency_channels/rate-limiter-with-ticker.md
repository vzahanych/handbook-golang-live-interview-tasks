# rate limiter with ticker

## Live interview task
Throttle work using time.Ticker.

## Concepts covered
- time.Ticker
- rate limiting

## Candidate solution

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(100 * time.Millisecond)
    defer ticker.Stop()
    for i := 0; i < 3; i++ {
        <-ticker.C
        fmt.Println("request", i, time.Now().Format("15:04:05.000"))
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
