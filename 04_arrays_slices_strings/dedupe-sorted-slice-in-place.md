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

func dedupeSorted(s []int) []int {
    if len(s) < 2 {
        return s
    }
    w := 1
    for r := 1; r < len(s); r++ {
        if s[r] != s[w-1] {
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
