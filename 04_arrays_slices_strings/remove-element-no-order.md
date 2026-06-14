# remove element no order

## Live interview task
Remove an element by index in O(1) when order does not matter.

## Concepts covered
- slices
- O(1) deletion
- swap-with-last

## Candidate solution

```go
package main

import "fmt"

// removeAtNoOrder deletes s[i] in O(1) by swapping in the last element — order is not preserved.
//
// Example: [10, 20, 30, 40], i=1
//   s[1] = s[3]     → [10, 40, 30, 40]
//   zero last slot   → [10, 40, 30, 0]  (help GC if T holds pointers)
//   return s[:3]    → [10, 40, 30]     (20 is gone; 30 moved from index 2 to stay, 40 filled hole at 1)
func removeAtNoOrder[T any](s []T, i int) []T {
    if i < 0 || i >= len(s) {
        return s
    }
    s[i] = s[len(s)-1] // overwrite removed slot with tail element
    var zero T
    s[len(s)-1] = zero // clear dropped slot so pointers can be collected
    return s[:len(s)-1] // shrink length; cap unchanged
}

func main() {
    fmt.Println(removeAtNoOrder([]int{10, 20, 30, 40}, 1)) // [10 40 30] order not preserved
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Copy last element into hole `i`, shrink length — O(1).
- Used in game engines, connection pools, set-like slices where order is irrelevant.
- Invalid index: return unchanged or panic — document contract.
- Do not use when iteration order or stability matters.

## Q&A

**Q: Remove by value O(1)?**  
A: Find index O(n), then swap-delete O(1) — total O(n) search.

**Q: vs `append(s[:i], s[i+1:]...)`?**  
A: That shifts O(n) elements; swap-delete is O(1).

**Q: Concurrent access?**  
A: Not safe — need mutex; or single goroutine owns slice.

**Q: Edge cases?**  
A: Remove last element (degenerate swap), single-element slice, invalid index.
