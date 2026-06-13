# clip slice capacity before append

## Live interview task
Use full slice expression `s[low:high:max]` so append does not overwrite shared backing array.

## Concepts covered
- full slice expression
- append allocation
- slice aliasing

## Candidate solution

```go
package main

import "fmt"

func main() {
    base := []int{1, 2, 3, 4}
    a := base[:2:2] // len=2, cap=2 — cannot grow into base[2:]
    b := append(a, 99)
    fmt.Println("base", base) // [1 2 3 4] unchanged
    fmt.Println("b", b)       // [1 2 99] new array likely
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `base[:2]` has cap to end of base — append may write into `base[2]`.
- `base[:2:2]` sets cap = 2 — append reallocates if len==cap.
- `copy` to new slice is alternative when cloning sub-slice.
- Interview pairs with slice-append-shared-backing-trap in category 03.

## Q&A

**Q: Syntax `s[low:high:max]`?**  
A: len = high-low, cap = max-low.

**Q: When needed?**  
A: Sub-slice then append without affecting parent.

**Q: `append([]T(nil), s...)`?**  
A: Clone idiom — always independent copy.

**Q: Complexity?**  
A: O(1) header change or O(k) if new alloc on append.

**Q: Gotcha without clip?**  
A: `sub[0]=99` mutates parent if shared cap.
