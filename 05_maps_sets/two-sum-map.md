# two sum map

## Live interview task
Given a slice and a `target`, return **two distinct indexes** whose values add up to `target` (each index used at most once). Use a map for O(n) complement lookup instead of nested loops. Example: `[2, 7, 11, 15]`, target `9` → indexes `0` and `1` because `2 + 7 = 9`.

## Concepts covered
- maps
- complement lookup
- single-pass hash map

## Candidate solution

```go
package main

import "fmt"

// twoSum finds two distinct indexes i, j with a[i]+a[j] == target.
// Single pass: for each v at i, check if complement (target-v) was seen earlier.
//
// Example: [2, 7, 11, 15], target=9
//   i=0, v=2: need 7 — not in map → store 2→0
//   i=1, v=7: need 2 — seen at index 0 → return 0, 1
func twoSum(a []int, target int) (int, int, bool) {
    seen := make(map[int]int, len(a)) // value → index where we saw it
    for i, v := range a {
        if j, ok := seen[target-v]; ok {
            return j, i, true // j is earlier index, i is current
        }
        seen[v] = i // store after check — avoids pairing an element with itself
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
