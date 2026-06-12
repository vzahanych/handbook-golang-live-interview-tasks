# avoid large range value copy

## Live interview task
Avoid copying large elements by ranging over indexes instead of values.

## Concepts covered
- range copy cost
- large structs

## Candidate solution

```go
package main

import "fmt"

type Big struct { Data [1024]byte; N int }

func sum(xs []Big) int {
    total := 0
    for i := range xs { // avoids copying Big into second range variable
        total += xs[i].N
    }
    return total
}

func main() { fmt.Println(sum([]Big{{N:1},{N:2}})) }
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
