# custom error type with errors as

## Live interview task
Create a custom error type and extract it with errors.As.

## Concepts covered
- custom errors
- errors.As
- wrapping

## Candidate solution

```go
package main

import (
    "errors"
    "fmt"
)

type StatusError struct { Code int; Msg string }
func (e *StatusError) Error() string { return fmt.Sprintf("status %d: %s", e.Code, e.Msg) }

func call() error { return fmt.Errorf("api: %w", &StatusError{Code: 503, Msg: "unavailable"}) }

func main() {
    var se *StatusError
    if errors.As(call(), &se) { fmt.Println(se.Code) }
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
