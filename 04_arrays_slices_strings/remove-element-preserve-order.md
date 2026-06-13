# remove element preserve order

## Live interview task
Remove all occurrences of a value while preserving order.

## Concepts covered
- slices
- comparable
- GC-friendly clearing

## Candidate solution

```go
package main

import "fmt"

func removeAll[T comparable](s []T, bad T) []T {
    w := 0
    for _, v := range s {
        if v != bad {
            s[w] = v
            w++
        }
    }
    var zero T
    for i := w; i < len(s); i++ {
        s[i] = zero // release references for GC (pointers, strings, slices)
    }
    return s[:w]
}

func main() {
    fmt.Println(removeAll([]string{"a", "x", "b", "x"}, "x")) // [a b]
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- In-place filter: read index implicit in `range`, write index `w`.
- Zeroing tail prevents memory leak when `T` holds pointers (slice of `*BigStruct`).
- `slices.DeleteFunc` (Go 1.21+) is the stdlib version.
- Preserving order costs O(n); swap-with-last is O(1) per delete but shuffles.

## Q&A

**Q: Why zero the tail?**  
A: `s[:w]` hides elements but cap still references them — GC cannot collect if pointers remain.

**Q: `T` not comparable?**  
A: Use predicate `func(T) bool` instead of `bad T`.

**Q: Complexity?**  
A: O(n) time, O(1) extra space.

**Q: Edge cases?**  
A: No matches (return same content), all match (empty slice), empty input.
