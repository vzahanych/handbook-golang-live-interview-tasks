# defer modifies named result

## Live interview task
Show how a deferred function can observe and modify named return values.

## Concepts covered
- defer
- named results
- return

## Candidate solution

```go
package main

import "fmt"

func compute() (n int) {
    defer func() { n *= 2 }()
    return 21
}

func main() { fmt.Println(compute()) }
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
