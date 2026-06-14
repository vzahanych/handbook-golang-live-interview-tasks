# first result wins

## Live interview task
Start **n** workers (hedged requests) and return the first result — e.g. three queries with delays 300ms / 100ms / 10ms → `"fast"` wins.

## Concepts covered
- select
- racing goroutines
- buffered channel

## Candidate solution

```go
package main

import (
    "fmt"
    "time"
)

func query(name string, d time.Duration) <-chan string {
    ch := make(chan string, 1) // buffer 1 so sender never blocks after result
    go func() {
        time.Sleep(d)
        ch <- name
    }()
    return ch
}

// first returns the value from whichever worker channel delivers first.
// Works for any n >= 1 — no recursive select or channel-per-pair trickery.
//
// Pattern: one goroutine per input channel, all racing to send into shared out.
// Example with 3 workers (10ms, 100ms, 300ms):
//   all goroutines block on <-c until their worker finishes
//   "fast" finishes first → out <- "fast" succeeds
//   "medium"/"slow" finish later → out is full → default branch (no block)
//   main receives "fast" from <-out
func first(chs ...<-chan string) string {
    switch len(chs) {
    case 0:
        return ""
    case 1:
        return <-chs[0] // nothing to race
    }
    out := make(chan string, 1) // holds exactly one winner; lets losers drop their send
    for _, ch := range chs {
        go func(c <-chan string) { // pass c — avoid loop variable capture bug
            v := <-c // wait for this worker only
            select {
            case out <- v: // non-blocking: only first successful send wins
            default:        // later finishers exit without blocking main
            }
        }(ch)
    }
    return <-out // blocks until some goroutine wins the race
}

func main() {
    workers := []struct {
        name string
        wait time.Duration
    }{
        {"slow", 300 * time.Millisecond},
        {"medium", 100 * time.Millisecond},
        {"fast", 10 * time.Millisecond},
    }
    chs := make([]<-chan string, len(workers))
    for i, w := range workers {
        chs[i] = query(w.name, w.wait)
    }
    fmt.Println(first(chs...)) // fast
}
```

## Two-way shortcut (interview)

When you only have two workers, a plain `select` is enough:

```go
select {
case v := <-query("fast", 10*time.Millisecond):
    fmt.Println(v)
case v := <-query("slow", 100*time.Millisecond):
    fmt.Println(v)
}
```

For **n > 2**, use the `first` helper above (one goroutine per channel + single result channel).

## Run

```bash
go run .
```

## Interview notes / pitfalls
- **`first` scales to any n**: spawn one goroutine per input channel; each waits on its channel, then tries a non-blocking send to a shared `out` (buffer 1). First send wins; losers hit `default` and exit.
- Buffer size 1 on **query** result channel — slow worker won't block forever after it finishes (even though nobody reads the loser).
- Loser goroutines still run — leak/cancel with `context` in production.
- `select` chooses pseudo-randomly if multiple ready at once — fine for equal priority.
- Hedged requests pattern in distributed systems — mention canceling losers via context.

## Q&A

**Q: Goroutine leak?**  
A: Yes if slow query never read — use ctx cancel or timeout.

**Q: First error wins?**  
A: Return `(T, error)` channels or `select` on err ch.

**Q: Complexity?**  
A: O(1) select; wall time = min(latencies).

**Q: vs `errgroup`?**  
A: errgroup waits for all; this returns first success.

**Q: Production?**  
A: `context.WithTimeout` + multiple backend calls + cancel on first response.
