# strings builder grow once

## Live interview task
Reduce allocations by growing strings.Builder once.

## Concepts covered
- strings.Builder
- Grow

## Candidate solution

```go
package main

import (
    "fmt"
    "strings"
)

func repeat(word string, n int) string {
    var b strings.Builder
    b.Grow(len(word) * n)
    for i := 0; i < n; i++ { b.WriteString(word) }
    return b.String()
}

func main() { fmt.Println(repeat("go", 3)) }
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
