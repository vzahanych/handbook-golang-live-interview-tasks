# errors join go120

## Live interview task
Combine multiple errors with `errors.Join` (Go 1.20+) and check membership with `errors.Is`.

## Concepts covered
- errors.Join
- multi-error
- validation aggregates

## Candidate solution

```go
package main

import (
    "errors"
    "fmt"
)

var ErrName = errors.New("invalid name")
var ErrAge = errors.New("invalid age")

func validate(name string, age int) error {
    var errs []error
    if name == "" {
        errs = append(errs, ErrName)
    }
    if age < 0 {
        errs = append(errs, ErrAge)
    }
    return errors.Join(errs...)
}

func main() {
    err := validate("", -1)
    fmt.Println(err)
    fmt.Println(errors.Is(err, ErrName)) // true
    fmt.Println(errors.Is(err, ErrAge))  // true
}
```

## Run

```bash
go run . # Go 1.20+
```

## Interview notes / pitfalls
- `errors.Join` returns `nil` if all args nil — like `append` idiom.
- `Unwrap()` on joined error returns `[]error` — `Is`/`As` check each in chain per Go 1.20+ rules.
- Different from `%w` wrap — join is **sibling** errors, not layered context.
- Use for validation collecting all field errors; wrap for call stack context.

## Q&A

**Q: Join vs `fmt.Errorf("%w; %w")`?**  
A: Join is structured multi-error; double `%w` not supported in one `Errorf`.

**Q: `err.Error()` output?**  
A: Joined messages separated by newlines.

**Q: Handle in HTTP?**  
A: Return 400 with all validation messages from joined unwrap.

**Q: Empty join?**  
A: `errors.Join()` → nil.

**Q: vs multierror packages?**  
A: Stdlib sufficient since 1.20 for most cases.
