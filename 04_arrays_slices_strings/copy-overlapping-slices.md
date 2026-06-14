# copy overlapping slices

## Live interview task
Use `copy` correctly when source and destination overlap.

## Concepts covered
- copy builtin
- overlapping slices

## Candidate solution

```go
package main

import "fmt"

func main() {
    s := []int{1, 2, 3, 4, 5}
    // Shift the first three elements one slot right inside the same backing array:
    //   dst = s[1:] → indices 1..4  (len 4)
    //   src = s[:3] → indices 0..2  (len 3)
    // dst and src overlap at indices 1 and 2 — a naive loop can clobber unread source bytes.
    //
    // copy() is safe (memmove semantics). It copies min(len(dst), len(src)) = 3 elements:
    //   s[1] ← s[0] = 1
    //   s[2] ← s[1] = 2   (from original source, not the value just written)
    //   s[3] ← s[2] = 3
    // s[0] and s[4] are untouched.
    copy(s[1:], s[:3])
    fmt.Println(s) // [1 1 2 3 5]
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- **`copy(dst, src)` handles overlap** like C `memmove`, not `memcpy` — safe for `copy(s[i:], s[j:])` on one backing array.
- **Returns** `min(len(dst), len(src))` — here `min(4, 3) = 3`; extra dst slots are left unchanged (`s[4]` stays `5`).
- **Naive manual loop fails** on overlap: writing `dst[i] = src[i]` forward can overwrite `src[i+1]` before it is read when regions overlap.
- **In-place delete at `i`**: `copy(s[i:], s[i+1:])` then shorten with `s = s[:len(s)-1]`.
- **Independent clone**: `dst := make([]T, len(s)); copy(dst, s)` or `slices.Clone(s)` — different backing array, no overlap issue.

## Q&A

**Q: `copy(s[1:], s[:3])` step by step?**  
A: Before: `[1,2,3,4,5]`. Overlap at indices 1–2. After three assignments: `[1,1,2,3,5]` — element at index 0 unchanged, index 4 unchanged.

**Q: Clone independent slice?**  
A: `dst := make([]T, len(s)); copy(dst, s)` or `slices.Clone(s)`.

**Q: Complexity?**  
A: O(min(len(dst), len(src))).

**Q: String to bytes?**  
A: `copy(b, s)` — string is read-only source; cannot `copy` into string.

**Q: Edge cases?**  
A: Nil dst/src (returns 0), zero lengths, full overlap same start.
