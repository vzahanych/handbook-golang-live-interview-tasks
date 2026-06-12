# clip slice capacity before append

## Live interview task
Use full slice expression to force append to allocate instead of overwriting shared backing array.

## Concepts covered
- full slice expression
- capacity control
- append

## Candidate solution

```go
package main

import "fmt"

func main() {
    base := []int{1,2,3,4}
    a := base[:2:2] // cap is clipped to len
    b := append(a, 99)
    fmt.Println(base, b)
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
