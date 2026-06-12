# mutex protected counter

## Live interview task
Protect shared state with sync.Mutex.

## Concepts covered
- sync.Mutex
- data races

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

type Counter struct { mu sync.Mutex; n int }
func (c *Counter) Inc() { c.mu.Lock(); c.n++; c.mu.Unlock() }
func (c *Counter) Value() int { c.mu.Lock(); defer c.mu.Unlock(); return c.n }

func main() { var c Counter; var wg sync.WaitGroup; for i:=0;i<1000;i++{ wg.Add(1); go func(){ defer wg.Done(); c.Inc() }() }; wg.Wait(); fmt.Println(c.Value()) }
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
