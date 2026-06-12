# rotate slice left k

## Live interview task
Rotate a slice left by k positions using the three-reversal algorithm.

## Concepts covered
- slice expressions
- in-place algorithms

## Candidate solution

```go
package main

import "fmt"

func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

func rotateLeft(s []int, k int) {
    if len(s) == 0 { return }
    k %= len(s)
    if k < 0 { k += len(s) }
    reverse(s[:k])
    reverse(s[k:])
    reverse(s)
}

func main() {
    s := []int{1,2,3,4,5}
    rotateLeft(s, 2)
    fmt.Println(s)
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
