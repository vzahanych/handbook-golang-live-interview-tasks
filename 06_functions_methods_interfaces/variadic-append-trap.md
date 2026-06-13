# variadic append trap

## Live interview task
Explain variadic parameters and the classic `append(slice, otherSlice...)` spread syntax.

## Concepts covered
- variadic functions
- slice spread
- append

## Candidate solution

```go
package main

import "fmt"

func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func merge(base []int, more ...int) []int {
    return append(base, more...)
}

func main() {
    fmt.Println(sum(1, 2, 3)) // 6

    a := []int{1, 2}
    b := []int{3, 4}
    // append(a, b) would append slice as single element — wrong type
    fmt.Println(merge(a, b...)) // [1 2 3 4]

    fmt.Println(sum(a...)) // spread slice into variadic call
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Variadic `...int` is `[]int` inside function — `nums` is a slice.
- Spread `b...` only valid when `b` is slice of matching element type.
- `append(s, s...)` on overlapping slice — `copy`/`append` rules apply; dangerous.
- Empty variadic: `sum()` → empty slice, loop zero times → 0.

## Q&A

**Q: Pass nil slice to variadic?**  
A: `nums` is nil, `range` is no-op — valid.

**Q: One slice arg without `...`?**  
A: `merge(a, b)` compile error — `b` is `[]int`, need `int`.

**Q: `fmt.Println` variadic?**  
A: `func Println(a ...any)` — accepts anything.

**Q: Complexity?**  
A: merge O(len(more)); sum O(n).

**Q: Interview follow-up?**  
A: Implement `max(nums ...int)` with empty input error.
