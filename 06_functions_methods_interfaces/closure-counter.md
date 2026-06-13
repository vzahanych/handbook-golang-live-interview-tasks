# closure counter

## Live interview task
Return a closure that captures and mutates local state.

## Concepts covered
- function literals
- closures
- heap escape

## Candidate solution

```go
package main

import "fmt"

func counter() func() int {
    n := 0
    return func() int {
        n++
        return n
    }
}

func main() {
    next := counter()
    fmt.Println(next(), next(), next()) // 1 2 3
    other := counter()
    fmt.Println(other()) // 1 — independent state
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Each call to `counter()` creates a **new** `n` — closures are independent.
- `n` escapes to heap because the closure outlives `counter`'s stack frame.
- Loop + closure: capture loop var correctly (Go 1.22+ per-iteration).
- Mutex needed if closure shared across goroutines.

## Q&A

**Q: Closure vs struct with method?**  
A: Closure for tiny state; struct when you need multiple methods or testing seams.

**Q: Complexity?**  
A: O(1) per call.

**Q: Can closures be garbage-collected?**  
A: Yes when no references remain — captured vars collected with closure.

**Q: Production use?**  
A: `http.HandlerFunc`, middleware factories, `sync.Once` patterns.

**Q: Edge cases?**  
A: Nil func call panics; returning closure that captures large structs — memory.
