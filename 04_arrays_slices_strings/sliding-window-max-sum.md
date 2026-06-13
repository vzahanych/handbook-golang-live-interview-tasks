# sliding window max sum

## Live interview task
Find the maximum sum of any contiguous window of size `k` in a slice.

## Concepts covered
- slices
- sliding window
- fixed window

## Candidate solution

```go
package main

import "fmt"

func maxWindowSum(a []int, k int) (int, bool) {
    if k <= 0 || k > len(a) {
        return 0, false
    }
    sum := 0
    for _, v := range a[:k] {
        sum += v
    }
    best := sum
    for i := k; i < len(a); i++ {
        sum += a[i] - a[i-k]
        if sum > best {
            best = sum
        }
    }
    return best, true
}

func main() {
    fmt.Println(maxWindowSum([]int{2, 1, 5, 1, 3, 2}, 3)) // 9 (5+1+3)
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Fixed window: add incoming, subtract outgoing — O(n) not O(n·k).
- Return `(0, false)` for invalid `k` — distinguish from valid zero sum.
- Variant: **variable** window (longest subarray with sum ≤ target) — two pointers, different problem.
- Variant: max in window (deque) — O(n) harder follow-up.

## Q&A

**Q: Complexity?**  
A: O(n) time, O(1) space.

**Q: All negative numbers?**  
A: Still works — best is least negative window.

**Q: `k == len(a)`?**  
A: One window — sum of entire array.

**Q: Follow-up: max **element** in each window?**  
A: Monotonic deque storing indices — O(n).

**Q: Edge cases?**  
A: `k > len`, `k <= 0`, empty `a`.
