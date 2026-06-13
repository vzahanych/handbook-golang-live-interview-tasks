# first result wins

## Live interview task
Start several workers and return the first successful result (hedged requests).

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

func first(chs ...<-chan string) string {
    switch len(chs) {
    case 0:
        return ""
    case 1:
        return <-chs[0]
    default:
        select {
        case v := <-chs[0]:
            return v
        case v := <-chs[1]:
            return v
        default:
            select {
            case v := <-chs[0]:
                return v
            case v := <-chs[1]:
                return v
            case v := <-firstN(chs[2:]...):
                return v
            }
        }
    }
}

func firstN(chs ...<-chan string) <-chan string {
    out := make(chan string, 1)
    go func() { out <- first(chs...) }()
    return out
}

func main() {
    fmt.Println(first(
        query("fast", 10*time.Millisecond),
        query("slow", 100*time.Millisecond),
    )) // fast
}
```

## Simpler two-way version (interview)

```go
select {
case v := <-query("fast", 10*time.Millisecond):
    fmt.Println(v)
case v := <-query("slow", 100*time.Millisecond):
    fmt.Println(v)
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Buffer size 1 on result channel — slow goroutine won't block forever after winner selected.
- Loser goroutines still run — leak/cancel with `context` in production.
- `select` chooses pseudo-randomly if multiple ready — fine for equal priority.
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
