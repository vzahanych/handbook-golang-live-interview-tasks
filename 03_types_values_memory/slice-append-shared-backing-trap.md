# slice append shared backing trap

## Live interview task
Show how `append` on a sub-slice can overwrite elements in the parent slice when capacity allows.

## Concepts covered
- slices
- append
- capacity
- aliasing

## Candidate solution

```go
package main

import "fmt"

func main() {
    orig := []int{1, 2, 3, 4}
    sub := orig[:2:2] // len=2, cap=2 — cannot append without new array
    sub = append(sub, 99)
    fmt.Println("safe", orig, sub) // [1 2 3 4] [1 2 99]

    orig2 := []int{1, 2, 3, 4}
    sub2 := orig2[:2] // len=2, cap=4 — shares tail capacity
    sub2 = append(sub2, 99)
    fmt.Println("trap", orig2, sub2) // [1 2 99 4] [1 2 99]
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `s[low:high]` cap extends to parent cap — append may write into parent's unused slots.
- Full slice expression `s[low:high:max]` limits cap — forces reallocation on append if `max-low` exceeded.
- Classic interview: pass sub-slice to func that appends — caller's data corrupted.
- Fix: `append([]T(nil), sub...)` or `copy` to new buffer.

## Q&A

**Q: When does append allocate?**  
A: When `len == cap` before append — new array ~2× cap (implementation-dependent growth).

**Q: `copy` vs append clone?**  
A: `dst := make([]T, len(s)); copy(dst, s)` — explicit; `append(nil, s...)` idiomatic.

**Q: Passing slice to goroutine?**  
A: Copy or ensure no concurrent append on shared backing array.

**Q: Complexity?**  
A: Append amortized O(1); copy O(n).

**Q: Reference?**  
A: [Go slices intro](https://go.dev/blog/slices-intro) — highly cited in interviews.
