# escape analysis example

## Live interview task
Show code that escapes to the heap and how to inspect with `-gcflags=-m`.

## Concepts covered
- escape analysis
- stack vs heap
- compiler diagnostics

## Candidate solution

```go
package main

import "fmt"

func escapes() *int {
    x := 42
    return &x // x escapes — must live after function returns
}

func noEscape() int {
    x := 42
    return x // x stays on stack
}

func main() {
    fmt.Println(*escapes(), noEscape())
}
```

## Run

```bash
go build -gcflags=-m .
# look for "moved to heap: x"
```

## Interview notes / pitfalls
- Returning pointer to local variable → escape to heap.
- Closing over variable in goroutine → often escapes.
- `fmt.Println` interface args may cause escape — hard to avoid in demos.
- Escape ≠ leak — GC collects unreachable heap objects.

## Q&A

**Q: Why escape analysis?**  
A: Stack allocation is cheap (frame pop); heap needs GC.

**Q: Reduce escapes?**  
A: Return values not pointers; pass buffers in; use generics without interface boxing.

**Q: `new(T)` vs `&T{}`?**  
A: Both heap allocate if result escapes or too large for stack.

**Q: Production tool?**  
A: `pprof` alloc_space — find hot heap allocs.

**Q: Interview command?**  
A: `go build -gcflags=-m` or `-m=2` verbose.
