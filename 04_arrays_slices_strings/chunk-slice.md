# chunk slice

## Live interview task
Split a slice into chunks of at most n elements.

## Concepts covered
- slices
- capacity
- subslice sharing

## Candidate solution

```go
package main

import "fmt"

func chunk[T any](s []T, n int) [][]T {
    if n <= 0 { panic("chunk size must be positive") }
    out := make([][]T, 0, (len(s)+n-1)/n)
    for len(s) > 0 {
        end := n
        if end > len(s) { end = len(s) }
        out = append(out, s[:end])
        s = s[end:]
    }
    return out
}

func main() { fmt.Println(chunk([]int{1,2,3,4,5}, 2)) }
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
