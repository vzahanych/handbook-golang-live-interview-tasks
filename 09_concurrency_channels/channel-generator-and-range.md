# channel generator and range

## Live interview task
Create a generator function that sends values then closes the channel.

## Concepts covered
- channels
- close
- range over channel

## Candidate solution

```go
package main

import "fmt"

func gen(n int) <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        for i := 0; i < n; i++ {
            ch <- i
        }
    }()
    return ch
}

func main() {
    for v := range gen(5) {
        fmt.Println(v) // 0..4
    }
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `range` on channel receives until **closed** — producer must `close(ch)`.
- Receive from closed channel: zero value + `ok==false` ends range.
- Unbuffered channel blocks until receiver ready — backpressure by default.
- Return `<-chan int` (send-only from caller view) — hides close responsibility inside generator.

## Q&A

**Q: Buffered generator?**  
A: `make(chan int, n)` — producer may not block until buffer full.

**Q: Who closes?**  
A: Sender/owner closes — never close from receiver (usually).

**Q: Context cancel?**  
A: `select { case ch<-v: case <-ctx.Done(): return }` in generator.

**Q: Complexity?**  
A: O(n) values; one goroutine per generator typical.

**Q: Edge cases?**  
A: `n==0` — close immediately, range no-ops.
