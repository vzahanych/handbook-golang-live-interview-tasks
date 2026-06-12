# switch without expression classifier

## Live interview task
Classify an integer using a switch without an expression.

## Concepts covered
- switch statements
- default true switch expression

## Candidate solution

```go
package main

import "fmt"

func classify(n int) string {
    switch {
    case n < 0:
        return "negative"
    case n == 0:
        return "zero"
    case n%2 == 0:
        return "positive even"
    default:
        return "positive odd"
    }
}

func main() { fmt.Println(classify(7)) }
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
