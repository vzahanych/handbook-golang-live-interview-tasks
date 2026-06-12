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
    fmt.Println(xs)
}
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
