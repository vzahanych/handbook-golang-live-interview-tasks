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
    for j := range jobs {
        results <- j * j
    }
}

func main() {
    jobs := make(chan int)
    results := make(chan int)
    var wg sync.WaitGroup

    for w := 0; w < 3; w++ {
        wg.Add(1)
        go worker(jobs, results, &wg)
    }

    go func() {
        for i := 1; i <= 5; i++ {
            jobs <- i
        }
        close(jobs)
        wg.Wait()
        close(results)
    }()

    for r := range results {
        fmt.Println(r)
    }
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Close `jobs` so workers exit `for range`.
- Wait for workers **before** `close(results)` — else readers may see premature close.
- Unbuffered channels synchronize each handoff — may deadlock if no consumer.
- Fixed worker count bounds concurrency — vs unbounded `go` per job.

## Q&A

**Q: Deadlock scenario?**  
A: Main reads results after workers block on send with no reader — need consumer goroutine or buffers.

**Q: Buffered results?**  
A: `make(chan int, 100)` decouples worker from main read speed.

**Q: Error handling?**  
A: `results chan Result` with `Err error` field or separate errs channel.

**Q: Complexity?**  
A: O(jobs) work; W workers in parallel.

**Q: vs `ants` / pool libs?**  
A: Same pattern — interview wants raw channels + WaitGroup.
