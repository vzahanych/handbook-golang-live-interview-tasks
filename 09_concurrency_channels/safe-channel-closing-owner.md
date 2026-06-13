# safe channel closing owner

## Live interview task
Demonstrate the channel closing rule: only the **sender** (owner) closes the channel.

## Concepts covered
- close
- channel ownership
- range exit

## Candidate solution

```go
package main

import "fmt"

func produce(n int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for i := 0; i < n; i++ {
            out <- i
        }
    }()
    return out
}

func consume(in <-chan int) {
    for v := range in {
        fmt.Println(v)
    }
}

func main() {
    consume(produce(3))
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- **Close only from sender** — receivers must not close (usually).
- Send on closed channel **panics**; receive on closed returns zero, ok=false.
- Close signals "no more values" — not a data value; use for range termination.
- Double close panics — use `sync.Once` for shared shutdown.

## Q&A

**Q: Why close?**  
A: Let consumers exit `for range` without guessing when stream ends.

**Q: Buffered channel close?**  
A: Receivers drain buffer after close, then get zero values.

**Q: Who closes worker pool jobs?**  
A: Producer closes `jobs` after enqueue done; workers never close jobs.

**Q: Detect close without range?**  
A: `v, ok := <-ch` — `!ok` means closed.

**Q: One-liner?**  
A: "Don't close channels on the receiving side."
