# tee channel values

## Live interview task
Duplicate each input value to two output channels.

## Concepts covered
- select
- nil channels in select
- broadcast

## Candidate solution

```go
package main

import "fmt"

func tee[T any](in <-chan T) (<-chan T, <-chan T) {
    a, b := make(chan T), make(chan T)
    go func() {
        defer close(a)
        defer close(b)
        for v := range in {
            out1, out2 := a, b
            for i := 0; i < 2; i++ {
                select {
                case out1 <- v:
                    out1 = nil // disable after send — select won't block on done branch
                case out2 <- v:
                    out2 = nil
                }
            }
        }
    }()
    return a, b
}

func main() {
    in := make(chan int, 1)
    in <- 7
    close(in)
    a, b := tee(in)
    fmt.Println(<-a, <-b) // 7 7
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Set channel to `nil` in select after send — nil case disabled, prevents duplicate send to same branch blocking forever.
- Slow consumer on one branch blocks tee goroutine — backpressure on both outputs.
- Buffered outputs reduce blocking — trade memory.
- True broadcast with N>2: extend loop or fan-out copy pattern.

## Q&A

**Q: Deadlock if one consumer gone?**  
A: Yes — need context cancel or separate goroutines per output with queues.

**Q: vs io.TeeReader?**  
A: Same idea for streams — bytes to two readers.

**Q: Complexity?**  
A: O(1) per value tee goroutine work.

**Q: Generic benefit?**  
A: One implementation for any element type.

**Q: Production?**  
A: Rare raw tee — often log + process pipelines use separate middleware.
