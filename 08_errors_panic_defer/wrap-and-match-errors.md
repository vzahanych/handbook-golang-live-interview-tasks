# wrap and match errors

## Live interview task
Wrap errors with context and detect sentinel errors with errors.Is.

## Concepts covered
- errors
- fmt.Errorf %w
- errors.Is

## Candidate solution

```go
package main

import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("not found")

func load(id string) error { return fmt.Errorf("load %s: %w", id, ErrNotFound) }

func main() {
    err := load("42")
    fmt.Println(err)
    fmt.Println(errors.Is(err, ErrNotFound))
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
