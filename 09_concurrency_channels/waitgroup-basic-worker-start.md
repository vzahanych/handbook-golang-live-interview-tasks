# waitgroup basic worker start

## Live interview task
Run several goroutines and wait for all of them with `sync.WaitGroup`.

## Concepts covered
- goroutines
- sync.WaitGroup
- loop variables

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
        i := i // omit on Go 1.22+ with per-iteration semantics
        wg.Add(1)
        go func() {
            defer wg.Done()
            fmt.Println("worker", i)
        }()
    }
    wg.Wait()
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `wg.Add(1)` **before** `go` — race if Add after goroutine starts and Wait runs early.
- `Done()` once per `Add` — mismatch panics or hangs.
- Copy `WaitGroup` by value breaks — pass pointer or share one var.
- `Wait` blocks until counter zero — no way to wait with timeout (use channel/context instead).

## Q&A

**Q: `Add` inside goroutine?**  
A: Risky — Wait may run before all Adds; prefer Add in parent before launch.

**Q: vs `errgroup`?**  
A: errgroup waits + first error + cancel — better for worker pools with errors.

**Q: Complexity?**  
A: O(1) sync overhead per goroutine.

**Q: Edge cases?**  
A: Zero goroutines — Wait returns immediately.

**Q: Production?**  
A: Always `defer wg.Done()` at start of goroutine body.
