# dedupe sorted slice in place

## Live interview task
Remove duplicates from a sorted slice without allocating a new backing array.

## Concepts covered
- slices
- write index (read/write pointers)
- in-place filtering

## Candidate solution

```go
package main

import "fmt"

// dedupeSorted keeps one copy of each value in a sorted slice, in place.
// Requires ascending order — equal values are adjacent, so we only compare neighbors.
//
// Two indices:
//   r — read cursor, scans every element
//   w — write cursor, length of the unique prefix built so far
//
// Example [1, 1, 2, 2, 2, 3]:
//   start: s=[1,1,2,2,2,3], w=1 (s[0] is already the first unique)
//   r=1: s[1]==1 → skip duplicate
//   r=2: s[2]=2 is new → s[1]=2, w=2 → [1,2,2,2,2,3]
//   r=3,4: still 2 → skip
//   r=5: s[5]=3 is new → s[2]=3, w=3 → [1,2,3,2,2,3]
//   return s[:3] → [1 2 3] (tail may linger in backing array/cap)
func dedupeSorted(s []int) []int {
    if len(s) < 2 {
        return s
    }
    w := 1 // unique prefix is s[:w]; s[0] is always kept
    for r := 1; r < len(s); r++ {
        if s[r] != s[w-1] { // new value — different from last unique written
            s[w] = s[r]
            w++
        }
    }
    return s[:w]
}

func main() {
    fmt.Println(dedupeSorted([]int{1, 1, 2, 2, 2, 3})) // [1 2 3]
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Requires **sorted** input — unsorted dedupe needs a map/set.
- Return `s[:w]` — capacity may still hold old tail (GC concern for pointers — zero tail if needed).
- Same pattern as `removeAll`, `partition` — "write index" is a core live-coding pattern.
- `slices.Compact` (Go 1.21+) does this for sorted comparable slices.

## Q&A

**Q: Complexity?**  
A: O(n) time, O(1) extra space.

**Q: All duplicates?**  
A: If entire slice is one value, result length 1.

**Q: Unsorted input?**  
A: `map[T]struct{}` or sort first O(n log n).

**Q: Stable unique for unsorted?**  
A: Use map + single pass preserving first occurrence order.
