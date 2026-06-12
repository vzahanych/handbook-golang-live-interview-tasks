# remove element no order

## Live interview task
Remove an element by index in O(1) when order does not matter.

## Concepts covered
- slices
- O(1) deletion

## Candidate solution

```go
package main

import "fmt"

func removeAtNoOrder[T any](s []T, i int) []T {
    if i < 0 || i >= len(s) { return s }
    s[i] = s[len(s)-1]
    var zero T
    s[len(s)-1] = zero
    return s[:len(s)-1]
}

func main() { fmt.Println(removeAtNoOrder([]int{10,20,30,40}, 1)) }
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- None specific; discuss edge cases and complexity.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
