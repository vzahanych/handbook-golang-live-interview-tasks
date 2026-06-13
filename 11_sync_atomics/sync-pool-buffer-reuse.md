# sync pool buffer reuse

## Live interview task
Use `sync.Pool` to reuse temporary buffers and reduce allocations.

## Concepts covered
- sync.Pool
- allocation reduction
- Reset before reuse

## Candidate solution

```go
package main

import (
    "bytes"
    "fmt"
    "sync"
)

var pool = sync.Pool{
    New: func() any { return new(bytes.Buffer) },
}

func render(name string) string {
    b := pool.Get().(*bytes.Buffer)
    b.Reset()
    defer pool.Put(b)

    b.WriteString("hello ")
    b.WriteString(name)
    return b.String() // copies bytes — safe after Put
}

func main() {
    fmt.Println(render("go"))
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `Pool` objects may be **cleared at GC** — no guarantee of reuse; best-effort cache.
- Always `Reset()` object before reuse — stale data leak between requests.
- Return **copy** of buffer bytes before `Put` — other goroutines may Get same buffer.
- Don't pool objects with sensitive data without zeroing — security consideration.

## Q&A

**Q: When use Pool?**  
A: High-frequency alloc of same-shaped temp objects (buffers, structs) in hot path.

**Q: When not?**  
A: Long-lived objects, unbounded growth, objects with finalizers.

**Q: Per-P goroutine pools?**  
A: Pool shards internally — reduces contention.

**Q: `New` required?**  
A: `Get` returns nil if empty and no `New` — always provide `New`.

**Q: Complexity?**  
A: Get/Put O(1); reduces GC pressure.
