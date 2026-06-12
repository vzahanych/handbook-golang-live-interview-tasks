# for range over integer go122

## Live interview task
Print indexes 0..n-1 using the Go 1.22 integer range form.

## Concepts covered
- for range
- integer range
- Go 1.22

## Candidate solution

```go
package main

import "fmt"

func main() {
    const n = 5
    for i := range n {
        fmt.Println(i)
    }
}
```

## Run

```bash
go run . # requires Go 1.22+
```

## Interview notes / pitfalls
- None specific; discuss edge cases and complexity.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
