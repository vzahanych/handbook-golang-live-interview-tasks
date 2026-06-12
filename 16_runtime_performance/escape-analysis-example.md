# escape analysis example

## Live interview task
Show code likely to escape to heap and how to inspect it.

## Concepts covered
- escape analysis
- heap allocation

## Candidate solution

```go
package main

import "fmt"

func ptr() *int {
    x := 42
    return &x // x must outlive the function call
}

func main() { fmt.Println(*ptr()) }
```

## Run

```bash
go build -gcflags=-m .
```

## Interview notes / pitfalls
- None specific; discuss edge cases and complexity.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
