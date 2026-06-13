# generic channel fan in

## Live interview task
Write a generic fan-in function merging multiple receive-only channels.

## Concepts covered
- generics
- channels
- fan-in
- WaitGroup

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

func FanIn[T any](chs ...<-chan T) <-chan T {
    out := make(chan T)
    var wg sync.WaitGroup
    wg.Add(len(chs))
    for _, ch := range chs {
        go func(c <-chan T) {
            defer wg.Done()
            for v := range c {
                out <- v
            }
        }(ch)
    }
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}

func main() {
    a, b := make(chan int, 1), make(chan int, 1)
    a <- 1
    close(a)
    b <- 2
    close(b)
    for v := range FanIn(a, b) {
        fmt.Println(v) // 1 and 2, order nondeterministic
    }
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Pass `ch` to goroutine param — avoid loop capture issues on older Go.
- Unbuffered `out` — senders block until consumer reads; buffer if producers run ahead.
- Close `out` only after all inputs drained — `WaitGroup` + closer goroutine.
- Zero channels: `wg.Add(0)`, close `out` immediately.

## Q&A

**Q: Deterministic merge order?**  
A: `select` round-robin or priority — fan-in is usually nondeterministic.

**Q: Context cancel?**  
A: Pass `ctx` — exit goroutines on `ctx.Done()`, drain or abandon carefully.

**Q: Complexity?**  
A: O(total messages); one goroutine per input channel.

**Q: vs `errgroup`?**  
A: errgroup for errors; fan-in for value streams.

**Q: Buffered out?**  
A: `make(chan T, n)` reduces blocking on fast producers.
