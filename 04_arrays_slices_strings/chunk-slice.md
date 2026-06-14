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

// chunk groups s into consecutive windows of up to n elements.
//
// Example: chunk([1,2,3,4,5], 2) → [1,2] | [3,4] | [5]
//   pass 1: s=[1,2,3,4,5] → take s[:2], then s becomes [3,4,5]
//   pass 2: s=[3,4,5]     → take s[:2], then s becomes [5]
//   pass 3: s=[5]         → end=min(2,1)=1, take s[:1], then s becomes []
//
// Each chunk is a subslice header pointing into the same backing array as s.
func chunk[T any](s []T, n int) [][]T {
    if n <= 0 {
        panic("chunk size must be positive")
    }
    // ceil(len(s)/n) without floats: (5+2-1)/2 = 3 slots for [[1,2],[3,4],[5]]
    out := make([][]T, 0, (len(s)+n-1)/n)
    for len(s) > 0 {
        end := n
        if end > len(s) {
            end = len(s) // last chunk when len(s) is not a multiple of n
        }
        out = append(out, s[:end]) // view only — elements stay in original array
        s = s[end:]                  // drop the head; repeat on the tail
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
- **`out` holds slice headers, not copies** — each chunk is `s[i:j]` into the original backing array. Writing `chunks[0][0] = 99` changes the source slice too.
- **`s = s[end:]`** rebinds `s` to the unprocessed tail; the loop stops when the tail is empty (`len(s) == 0`).
- **Last chunk** is shorter when `len(s) % n != 0` — `end = min(n, len(s))` handles that in one branch.
- **Independent chunks** (no shared memory): copy each window, e.g. `append([]T(nil), s[:end]...)`.
- **Batch processing**: same pattern for DB rows, HTTP pages, or worker batches of size `n`.

## Q&A

**Q: Complexity?**  
A: O(len(s)) time; O(k) chunk headers where k = ceil(len/n).

**Q: `n <= 0`?**  
A: Panic or return error — define contract; never infinite loop.

**Q: Empty input?**  
A: Returns empty `[][]T`, not `nil` slice of chunks — both OK, be consistent.

**Q: Production?**  
A: Use `slices.Chunk` (Go 1.23+) if available in your Go version.
