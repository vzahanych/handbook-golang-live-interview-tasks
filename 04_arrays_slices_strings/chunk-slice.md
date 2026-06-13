# chunk slice

## Live interview task
Split a slice into chunks of at most n elements.

## Concepts covered
- slices
- subslice sharing
- capacity

## Candidate solution

```go
package main

import "fmt"

func chunk[T any](s []T, n int) [][]T {
    if n <= 0 {
        panic("chunk size must be positive")
    }
    out := make([][]T, 0, (len(s)+n-1)/n)
    for len(s) > 0 {
        end := n
        if end > len(s) {
            end = len(s)
        }
        out = append(out, s[:end])
        s = s[end:]
    }
    return out
}

func main() {
    fmt.Println(chunk([]int{1, 2, 3, 4, 5}, 2)) // [[1 2] [3 4] [5]]
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Chunks are **subslices** of the original backing array — mutating a chunk may affect others.
- For independent chunks: `append([]T(nil), s[:end]...)` per chunk.
- Last chunk may be smaller than `n` — handle `end > len(s)`.
- Batch processing pattern: process DB rows / API pages in groups of n.

## Q&A

**Q: Complexity?**  
A: O(len(s)) time; O(k) chunk headers where k = ceil(len/n).

**Q: `n <= 0`?**  
A: Panic or return error — define contract; never infinite loop.

**Q: Empty input?**  
A: Returns empty `[][]T`, not `nil` slice of chunks — both OK, be consistent.

**Q: Production?**  
A: Use `slices.Chunk` (Go 1.23+) if available in your Go version.
