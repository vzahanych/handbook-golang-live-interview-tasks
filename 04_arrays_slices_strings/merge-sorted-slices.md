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
        if a[i] <= b[j] { out = append(out, a[i]); i++ } else { out = append(out, b[j]); j++ }
    }
    out = append(out, a[i:]...)
    out = append(out, b[j:]...)
    return out
}

func main() { fmt.Println(merge([]int{1,3,5}, []int{2,4,6})) }
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
