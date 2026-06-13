# binary search lower bound

## Live interview task
Implement `lower_bound`: first index `i` where `a[i] >= target` (sorted ascending).

## Concepts covered
- binary search
- half-open intervals

## Candidate solution

```go
package main

import "fmt"

func lowerBound(a []int, target int) int {
    lo, hi := 0, len(a)
    for lo < hi {
        mid := lo + (hi-lo)/2
        if a[mid] < target {
            lo = mid + 1
        } else {
            hi = mid
        }
    }
    return lo
}

func main() {
    a := []int{1, 2, 4, 4, 7}
    fmt.Println(lowerBound(a, 4))  // 2 — first 4
    fmt.Println(lowerBound(a, 5))  // 4 — would insert here
    fmt.Println(lowerBound(a, 0))  // 0
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Half-open range `[lo, hi)` — invariant: answer in `[lo, hi]`.
- `mid := lo + (hi-lo)/2` avoids int overflow (style point in interviews).
- `sort.Search` in stdlib — same semantics.
- `lower_bound` vs `upper_bound`: upper uses `a[mid] <= target` → first index **>** target.

## Q&A

**Q: Complexity?**  
A: O(log n) time, O(1) space.

**Q: Target not present?**  
A: Returns insertion index where it would go — may equal `len(a)`.

**Q: Duplicates?**  
A: Lower bound lands on **first** equal element.

**Q: `sort.Search` equivalent?**  
A: `sort.Search(len(a), func(i int) bool { return a[i] >= target })`.

**Q: Edge cases?**  
A: Empty slice returns 0; single element; all less than target.
