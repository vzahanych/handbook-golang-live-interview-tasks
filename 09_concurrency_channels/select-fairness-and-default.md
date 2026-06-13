# select fairness and default

## Live interview task
Explain `select` behavior when multiple cases are ready, and use `default` for non-blocking ops.

## Concepts covered
- select
- pseudo-random choice
- default branch

## Candidate solution

```go
package main

import "fmt"

func main() {
    a, b := make(chan int, 1), make(chan int, 1)
    a <- 1
    b <- 2

    // Both ready — Go picks one pseudo-randomly (not strict alternation)
    select {
    case v := <-a:
        fmt.Println("from a", v)
    case v := <-b:
        fmt.Println("from b", v)
    }

  // Non-blocking try-receive
    ch := make(chan int)
    select {
    case v := <-ch:
        fmt.Println(v)
    default:
        fmt.Println("no data ready")
    }
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Multiple ready cases: **uniform pseudo-random** among ready — prevents starvation but nondeterministic.
- `default` makes entire select non-blocking — busy spin if misused.
- Nil channel case never selected — use to disable branches in dynamic `select` loops.
- Empty `select{}` blocks forever — sometimes intentional deadlock detection.

## Q&A

**Q: Fair round-robin?**  
A: Not built-in — implement with counter or separate scheduler.

**Q: `select` with single case?**  
A: Compiles — blocks like receive/send with possible ctx in other arms added later.

**Q: Timeout pattern?**  
A: `select { case <-ch: case <-time.After(d): }` — prefer `context.WithTimeout`.

**Q: vs mutex?**  
A: Channels for coordination; mutex for shared state — "share memory by communicating".

**Q: Production default?**  
A: Avoid tight loops with `default` — burns CPU.
