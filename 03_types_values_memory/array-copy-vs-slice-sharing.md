# array copy vs slice sharing

## Live interview task
Demonstrate that arrays copy their elements while slices share an underlying array.

## Concepts covered
- value representation
- arrays
- slices
- copy semantics

## Candidate solution

```go
package main

import "fmt"

func main() {
    a := [3]int{1, 2, 3}
    b := a
    b[0] = 99
    fmt.Println("array", a, b)

    s := []int{1, 2, 3}
    t := s
    t[0] = 99
    fmt.Println("slice", s, t)
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
