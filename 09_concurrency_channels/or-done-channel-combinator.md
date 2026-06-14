# or done channel combinator

## Live interview task
Implement **`or`**: given n “done” channels (`<-chan struct{}`), return one channel that **closes as soon as any input closes** — logical OR for cancellation signals. Use case: stop a pipeline when *either* the user cancels, a timeout fires, or an upstream stage finishes. Example: `a` is already closed → `<-or(a)` returns immediately and the combined channel is closed.

## Concepts covered
- select
- channel combinators
- recursion

## Candidate solution

```go
package main

import "fmt"

func or(chs ...<-chan struct{}) <-chan struct{} {
    done := make(chan struct{})
    go func() {
        defer close(done)
        switch len(chs) {
        case 0:
            return
        case 1:
            <-chs[0]
        default:
            select {
            case <-chs[0]:
            case <-chs[1]:
            case <-or(chs[2:]...):
            }
        }
    }()
    return done
}

func main() {
    a := make(chan struct{})
    close(a)
    <-or(a)
    fmt.Println("closed")
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Recursive `or` builds binary tree of selects — classic Go Concurrency Patterns talk.
- Go 1.22+ `for range` over channel — still need combinator for "first of N closes".
- Production: prefer `context.Context` — `ctx.Done()` is standard OR of cancel + timeout + parent.
- Goroutine leak if inputs never close — always pair with cancel.

## Q&A

**Q: Why not `select` with N cases in loop?**  
A: `select` cases must be static at compile time in one statement — dynamic N needs recursion or reflect.

**Q: vs `context`?**  
A: Context is idiomatic cancellation; channel `or` teaches select composition.

**Q: Already closed channel?**  
A: Receive returns immediately — `or` completes.

**Q: Complexity?**  
A: O(log n) goroutines in recursive version.

**Q: Edge case len 0?**  
A: Close output immediately — no inputs to wait on.
