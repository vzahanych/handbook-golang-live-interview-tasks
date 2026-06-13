# send on closed channel panic

## Live interview task
Demonstrate that sending on a closed channel panics, and show safe shutdown patterns.

## Concepts covered
- close
- panic
- channel state

## Candidate solution

```go
package main

import "fmt"

func main() {
    ch := make(chan int, 1)
    ch <- 1
    close(ch)

    // Receive OK after close
    v, ok := <-ch
    fmt.Println("recv", v, ok) // 1 true

    v, ok = <-ch
    fmt.Println("recv", v, ok) // 0 false

    // ch <- 2 // panic: send on closed channel
}
```

## Run

```bash
go run .
```

## Safe shutdown pattern

```go
func producer(done <-chan struct{}, out chan<- int) {
    defer close(out)
    for i := 0; i < 10; i++ {
        select {
        case <-done:
            return
        case out <- i:
        }
    }
}
```

## Interview notes / pitfalls
- **Send on closed** → panic immediately.
- **Close closed** → panic.
- **Receive on closed** → zero value, `ok=false` — safe.
- Use `select` with `done` channel or `context` to stop sending instead of closing from receiver side.

## Q&A

**Q: How detect closed before send?**  
A: Can't reliably — design ownership so only closer stops sends.

**Q: `recover` send panic?**  
A: Possible but smell — fix protocol instead.

**Q: sync.Once close?**  
A: `var once sync.Once; once.Do(func(){ close(ch) })` — idempotent close helper.

**Q: Test?**  
A: Document contract; race detector catches misuse.

**Q: Interview pairing?**  
A: Always teach with safe-channel-closing-owner example.
