# select default nonblocking receive

## Live interview task
Implement non-blocking receive with `select` and `default`.

## Concepts covered
- select
- default
- non-blocking receive
- generics

## Candidate solution

```go
package main

import "fmt"

func tryRecv[T any](ch <-chan T) (T, bool) {
    select {
    case v := <-ch:
        return v, true
    default:
        var zero T
        return zero, false
    }
}

func main() {
    ch := make(chan int, 1)
    fmt.Println(tryRecv(ch)) // 0 false
    ch <- 10
    fmt.Println(tryRecv(ch)) // 10 true
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `default` makes the whole `select` non-blocking — if no case ready, `default` runs immediately.
- Non-blocking **send**: `select { case ch <- v: ...; default: ... }`.
- Busy-loop with `default` burns CPU — use blocking receive or `time.After` / context in real code.
- Closing a channel: receive returns zero value immediately with `ok == false` — distinguish from empty buffered channel.

## Q&A

**Q: What if multiple cases are ready?**  
A: Go pseudo-randomly picks one — do not rely on priority without separate logic.

**Q: `nil` channel in select?**  
A: Receive/send on `nil` channel blocks forever; in `select`, nil cases are ignored (never chosen).

**Q: When is this pattern used?**  
A: Event loops, draining buffered work without blocking, combining with `time.Ticker` in one loop.

**Q: Complexity?**  
A: O(1) per attempt.

**Q: Production alternative?**  
A: `context.Context` cancellation, `chan struct{}` done signal, or `sync/atomic` flags for simple state.
