# sync once lazy init

## Live interview task
Initialize a value exactly once with `sync.Once`.

## Concepts covered
- sync.Once
- lazy initialization
- singleton

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

var (
    once   sync.Once
    config map[string]string
)

func Config() map[string]string {
    once.Do(func() {
        config = map[string]string{"env": "dev"}
    })
    return config
}

func main() {
    fmt.Println(Config(), Config()) // same map, init once
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `once.Do(f)` runs `f` exactly once — concurrent callers block until first completes.
- If `f` panics, `Do` considers unfinished — **retry on next call** (rare edge case).
- Don't return mutable global from `Config()` without warning — callers can mutate shared map.
- Prefer `sync.Once` inside struct for per-instance lazy init.

## Q&A

**Q: vs `init()`?**  
A: `init` always runs at startup; `Once` defers until first use — faster cold start.

**Q: Pass init error?**  
A: `Once` no error return — use `sync.OnceValues` (Go 1.21+) or manual `sync.Mutex` + `err`.

**Q: Double-checked locking?**  
A: `Once` implements correct DCL — don't hand-roll.

**Q: Reset Once?**  
A: Not supported — new `Once` variable or redesign.

**Q: Complexity?**  
A: First call pays init; later O(1) atomic fast path.
