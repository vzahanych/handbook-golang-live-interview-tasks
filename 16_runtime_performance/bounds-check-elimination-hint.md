# bounds check elimination hint

## Live interview task
Write slice code in a bounds-check-friendly way.

## Concepts covered
- bounds check elimination
- index expressions

## Candidate solution

```go
package main

import "fmt"

func sum4(s []int) int {
    _ = s[3] // one early check can help compiler eliminate later checks
    return s[0] + s[1] + s[2] + s[3]
}

func main() { fmt.Println(sum4([]int{1,2,3,4})) }
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
