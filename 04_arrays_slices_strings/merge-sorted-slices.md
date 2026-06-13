# merge sorted slices

## Live interview task
Merge two sorted integer slices into a new sorted slice.

## Concepts covered
- slices
- append
- two pointers

## Candidate solution

```go
package main

import "fmt"

func merge(a, b []int) []int {
    out := make([]int, 0, len(a)+len(b))
    i, j := 0, 0
    for i < len(a) && j < len(b) {
        if a[i] <= b[j] {
            out = append(out, a[i])
            i++
        } else {
            out = append(out, b[j])
            j++
        }
    }
    out = append(out, a[i:]...)
    out = append(out, b[j:]...)
    return out
}

func main() {
    fmt.Println(merge([]int{1, 3, 5}, []int{2, 4, 6})) // [1 2 3 4 5 6]
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Preallocate `cap = len(a)+len(b)` — one allocation, no growth copies.
- `<=` vs `<` — `<=` keeps stability for duplicates from `a` first.
- Tail append `a[i:]` handles remaining elements in O(1) append of a slice.
- Building block for merge sort; also merge k lists with heap.

## Q&A

**Q: Complexity?**  
A: O(n+m) time, O(n+m) space for output.

**Q: Merge in-place into `a` with spare capacity?**  
A: Merge from end backward to avoid overwrite — harder live-coding task.

**Q: Empty inputs?**  
A: Loop skipped; tail append copies the non-empty side.

**Q: Follow-up?**  
A: Merge K sorted arrays — min-heap O(N log K).
