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

// produce owns out — only this goroutine sends and closes it.
// Returns receive-only <-chan int so callers cannot close or send by mistake.
func produce(n int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out) // close after last send — signals "no more values" to consumers
        for i := 0; i < n; i++ {
            out <- i
        }
    }()
    return out
}

// consume reads every value still in the channel; close means "no more coming",
// not "discard what's left". for range drains buffered values, then exits.
func consume(in <-chan int) {
    for v := range in {
        fmt.Println(v)
    }
}

func main() {
    consume(produce(3)) // prints 0, 1, 2 then range returns
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- **Close only from sender** — receivers must not close (usually).
- Send on closed channel **panics**; receive on closed returns zero, ok=false.
- Close signals "no more values will be sent" — receivers still get any values already in the buffer, then `ok=false`.
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
