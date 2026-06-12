# copy overlapping slices

## Live interview task
Use copy correctly when source and destination overlap.

## Concepts covered
- copy
- overlapping slices

## Candidate solution

```go
package main

import "fmt"

func main() {
    s := []int{1, 2, 3, 4, 5}
    copy(s[1:], s[:3])
    fmt.Println(s) // [1 1 2 3 5]
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
