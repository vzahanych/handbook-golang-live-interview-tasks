# two sum map

## Live interview task
Return indexes of two values whose sum equals target (one solution guaranteed or find any).

## Concepts covered
- maps
- complement lookup
- single-pass hash map

## Candidate solution

```go
package main

import "fmt"

func twoSum(a []int, target int) (int, int, bool) {
    seen := make(map[int]int, len(a)) // value -> index
    for i, v := range a {
        if j, ok := seen[target-v]; ok {
            return j, i, true
        }
        seen[v] = i
    }
    return 0, 0, false
}

func main() {
    fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // 0 1 true
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Store index **after** checking complement — avoids using same element twice (`target = 2*v` at same index).
- Pre-size map `make(map[int]int, len(a))` reduces rehashing.
- Variant: return values not indices; variant: sorted two-pointer O(n) after sort.
- LeetCode 1 — state the O(n) map approach vs O(n²) brute force.

## Q&A

**Q: Complexity?**  
A: O(n) time, O(n) space.

**Q: Duplicate values?**  
A: Map stores latest index — for "use distinct indices" check order matters; solution stores after check.

**Q: No solution?**  
A: Return `false` or error — clarify with interviewer.

**Q: Three sum?**  
A: Sort + two pointers outer loop — O(n²).

**Q: Negative numbers / zeros?**  
A: Same algorithm works.
