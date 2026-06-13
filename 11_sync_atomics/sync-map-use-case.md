# sync map use case

## Live interview task
When to use `sync.Map` vs `map` + `RWMutex`.

## Concepts covered
- sync.Map
- concurrent map
- cache patterns

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var m sync.Map

    m.Store("lang", "go")

    if v, ok := m.Load("lang"); ok {
        fmt.Println(v)
    }

    // Range all entries
    m.Range(func(k, v any) bool {
        fmt.Println(k, v)
        return true // false stops iteration
    })

    // LoadOrStore — good for lazy init per key
    actual, loaded := m.LoadOrStore("id", 1)
    fmt.Println(actual, loaded)
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `sync.Map` optimized for: **stable key set**, many `Load`, fewer `Store`/`Delete`, keys disjoint across goroutines.
- **Not** a drop-in for all maps — `map+RWMutex` often faster for general use.
- Keys and values are `any` — type assert on load.
- No len() — must Range to count.

## Q&A

**Q: When sync.Map wins?**  
A: In-memory cache of read-mostly config, connection table keyed by ID with rare insert.

**Q: When mutex map wins?**  
A: Frequent writes, need size, typed API, range under consistent snapshot.

**Q: `LoadOrStore`?**  
A: Atomic get-or-insert — dedupe lazy init per key.

**Q: Delete?**  
A: `m.Delete(k)` — no size update exposed.

**Q: Compare to `map` panic?**  
A: sync.Map safe concurrent — ordinary map needs lock.

**Q: Generic typed wrapper?**  
A: Wrap sync.Map with type-safe Store/Load methods in app code.
