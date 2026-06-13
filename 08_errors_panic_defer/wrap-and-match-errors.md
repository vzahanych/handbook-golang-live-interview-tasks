# wrap and match errors

## Live interview task
Wrap errors with context and detect sentinel errors with `errors.Is`.

## Concepts covered
- errors
- `fmt.Errorf` `%w`
- `errors.Is`

## Candidate solution

```go
package main

import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("not found")

func load(id string) error {
    return fmt.Errorf("load %s: %w", id, ErrNotFound)
}

func main() {
    err := load("42")
    fmt.Println(err)                      // load 42: not found
    fmt.Println(errors.Is(err, ErrNotFound)) // true
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `%w` wraps one error — chain with `fmt.Errorf("ctx: %w", err)` for stack of context.
- `errors.Is(err, target)` walks the unwrap chain — works through `%w` wraps.
- Use **sentinel** `var ErrX = errors.New(...)` for stable `errors.Is` checks.
- `err == ErrNotFound` fails after wrap — always `errors.Is` for wrapped errors.
- `errors.Join` (Go 1.20+) combines multiple errors — different from wrap.

## Q&A

**Q: `Is` vs `==`?**  
A: `==` only top level; `Is` traverses `Unwrap()`.

**Q: When not to wrap?**  
A: When returning sentinel unchanged and caller uses `Is` — still OK to add context with `%w`.

**Q: `%v` vs `%w`?**  
A: `%v` stringifies inner error but breaks `Is`/`As` chain.

**Q: Production pattern?**  
A: Wrap at boundaries (repo → service → handler); check with `Is`/`As` at decision points.

**Q: Complexity?**  
A: `Is` O(depth of wrap chain), typically small.
