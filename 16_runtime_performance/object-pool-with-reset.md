# object pool with reset

## Live interview task
Reuse objects with a pool — reset state before returning to pool.

## Concepts covered
- sync.Pool
- manual pool
- reset before reuse

## Candidate solution

```go
package main

import (
    "bytes"
    "fmt"
    "sync"
)

var pool = sync.Pool{
    New: func() any {
        return bytes.NewBuffer(make([]byte, 0, 1024))
    },
}

func process() string {
    b := pool.Get().(*bytes.Buffer)
    b.Reset()
    defer pool.Put(b)

    b.WriteString("hello")
    return b.String()
}

func main() {
    fmt.Println(process())
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- **Reset** before reuse — stale bytes/data leak between requests.
- `sync.Pool` objects may be GC'd anytime — pool is cache not storage.
- Return copies if next user may Get same buffer while you hold `[]byte` view.
- Clear sensitive data before Put in security-sensitive pools.

## Q&A

**Q: vs manual `[]*Buffer` stack?**  
A: sync.Pool handles per-P sharding and GC; manual pool for deterministic tests.

**Q: Pool poison?**  
A: Forgotten Reset — test with distinctive first user data.

**Q: When not to pool?**  
A: Rare allocation, huge objects, hard-to-reset state.

**Q: Complexity?**  
A: Get/Put O(1); reduces GC pressure.

**Q: Interview tie-in?**  
A: Same as sync-pool-buffer-reuse in category 11.
