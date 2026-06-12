# worker pool jobs results

## Live interview task
Implement a worker pool with jobs and results channels.

## Concepts covered
- worker pool
- directional channels
- WaitGroup

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

func worker(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for j := range jobs { results <- j * j }
}

func main() {
    jobs := make(chan int)
    results := make(chan int)
    var wg sync.WaitGroup
    for w := 0; w < 3; w++ { wg.Add(1); go worker(jobs, results, &wg) }
    go func(){ for i := 1; i <= 5; i++ { jobs <- i }; close(jobs); wg.Wait(); close(results) }()
    for r := range results { fmt.Println(r) }
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
