# custom error type with errors as

## Live interview task
Create a custom error type and extract it with `errors.As`.

## Concepts covered
- custom errors
- `errors.As`
- pointer receiver on `Error()`

## Candidate solution

```go
package main

import (
    "errors"
    "fmt"
)

type StatusError struct {
    Code int
    Msg  string
}

func (e *StatusError) Error() string {
    return fmt.Sprintf("status %d: %s", e.Code, e.Msg)
}

func call() error {
    return fmt.Errorf("api: %w", &StatusError{Code: 503, Msg: "unavailable"})
}

func main() {
    var se *StatusError
    if errors.As(call(), &se) {
        fmt.Println(se.Code) // 503
    }
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `errors.As(err, &target)` — `target` must be pointer to **pointer** or pointer to interface holding pointer type.
- `As` finds first matching type in unwrap chain — like type assertion but recursive.
- Use `As` for **structured** errors (HTTP code, field name); `Is` for sentinels.
- Value vs pointer: `As` into `*StatusError` matches `*StatusError` in chain, not `StatusError` value unless wrapped as value.

## Q&A

**Q: `As` vs type switch on `err`?**  
A: `As` handles wrapped errors; type switch only top dynamic type unless you unwrap manually.

**Q: Implement `Unwrap() error`?**  
A: On custom type to participate in chain — optional if only leaf error.

**Q: HTTP handler pattern?**  
A: `var se *StatusError; if errors.As(err, &se) { w.WriteHeader(se.Code) }`.

**Q: Nil pointer in wrap?**  
A: Wrapping typed nil pointer still causes `err != nil` — separate gotcha.

**Q: Complexity?**  
A: O(unwrap depth).
