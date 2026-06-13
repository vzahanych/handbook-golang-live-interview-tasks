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
    copy(s[1:], s[:3])
    fmt.Println(s) // [1 1 2 3 5]
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `copy(dst, src)` handles overlap like `memmove` — safe for `copy(s[i:], s[j:])`.
- Returns number of elements copied = `min(len(dst), len(src))`.
- Manual loop `for i := range dst { dst[i] = src[i] }` on overlap can clobber unread source — wrong.
- Shifting left in-place: often `copy(s[i:], s[i+1:])` to delete at i.

## Q&A

**Q: `copy(s[1:], s[:3])` step by step?**  
A: dst len 4, src len 3 → copies 3 elems: index 1←0, 2←1, 3←2 → `[1,1,2,3,5]`.

**Q: Clone independent slice?**  
A: `dst := make([]T, len(s)); copy(dst, s)` or `slices.Clone(s)`.

**Q: Complexity?**  
A: O(min(len(dst), len(src))).

**Q: String to bytes?**  
A: `copy(b, s)` — string is read-only source; cannot `copy` into string.

**Q: Edge cases?**  
A: Nil dst/src (returns 0), zero lengths, full overlap same start.
