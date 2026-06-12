# struct padding order fields

## Live interview task
Compare struct sizes when fields are ordered differently.

## Concepts covered
- struct padding
- unsafe.Sizeof
- alignment

## Candidate solution

```go
package main

import (
    "fmt"
    "unsafe"
)

type Bad struct {
    A bool
    B int64
    C bool
}

type Better struct {
    B int64
    A bool
    C bool
}

func main() {
    fmt.Println(unsafe.Sizeof(Bad{}), unsafe.Sizeof(Better{}))
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
