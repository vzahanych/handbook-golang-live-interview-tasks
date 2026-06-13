# panic safe parser

## Live interview task
Convert a panic-prone helper into an error-returning parser using `recover`.

## Concepts covered
- recover
- error conversion
- avoid panic in libraries

## Candidate solution

```go
package main

import (
    "fmt"
    "strconv"
)

func mustAtoi(s string) int {
    n, err := strconv.Atoi(s)
    if err != nil {
        panic(err)
    }
    return n
}

func parse(s string) (n int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("parse %q: %v", s, r)
        }
    }()
    return mustAtoi(s), nil
}

func main() {
    fmt.Println(parse("42")) // 42 <nil>
    fmt.Println(parse("x"))  // 0 parse "x": ...
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Prefer `strconv.Atoi` directly with error — recover wrapper is for legacy `Must*` APIs.
- Named result `err` lets defer assign error on panic.
- `return mustAtoi(s), nil` — if panic, `n` may be partial; only `err` from defer matters.
- `regexp.MustCompile` pattern — panic at init acceptable; runtime parse should return error.

## Q&A

**Q: Better design?**  
A: Delete `mustAtoi`; use `strconv.Atoi` in `parse` — idiomatic Go.

**Q: When panic OK?**  
A: Programmer errors, init-time `Must`, impossible internal state.

**Q: Typed panic to error?**  
A: `switch r := r.(type) { case error: err = r }`.

**Q: Libraries?**  
A: Never panic across public API boundary without documenting it.

**Q: Test?**  
A: `parse("x")` returns error, no process crash.
