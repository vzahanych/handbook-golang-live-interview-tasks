# reverse slice in place

## Live interview task
Reverse a slice in place using two indexes.

## Concepts covered
- slices
- generics
- in-place modification

## Candidate solution

```go
package main

import "fmt"

func reverse[T any](s []T) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

func main() {
    xs := []int{1, 2, 3, 4}
    reverse(xs)
    fmt.Println(xs) // [4 3 2 1]
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Mutates backing array — any other slice sharing that array sees the change.
- Empty or length-1 slice: loop does not run — correct no-op.
- For strings, convert to `[]rune` first (UTF-8) — byte reverse breaks multibyte chars.
- `slices.Reverse(s)` in stdlib (Go 1.21+) — mention you know the standard helper.

## Q&A

**Q: Time/space complexity?**  
A: O(n) time, O(1) extra space.

**Q: Reverse sub-slice only?**  
A: `reverse(s[low:high])` — bounds are on the sub-slice header.

**Q: Edge cases?**  
A: `nil` slice, single element, even vs odd length, slice with shared backing array.

**Q: Production?**  
A: Use `slices.Reverse` unless custom generic constraints needed.
