# shadowing short variable trap

## Live interview task
Find and fix the classic short variable declaration shadowing bug.

## Concepts covered
- short variable declarations
- scope
- shadowing

## Buggy version (interviewer may show this)

```go
func load(ok bool) (string, error) {
    var err error
    value := "default"

    if !ok {
        err, _ := errors.New("load failed"), value // BUG: err is NEW inner err
    } else {
        value = "loaded"
    }
    return value, err // always nil err from outer scope
}
```

## Candidate solution

```go
package main

import (
    "errors"
    "fmt"
)

func load(ok bool) (string, error) {
    var err error
    value := "default"

    if !ok {
        err = errors.New("load failed") // assignment, not :=
    } else {
        value = "loaded"
    }
    return value, err
}

func main() {
    fmt.Println(load(false)) // default load failed
    fmt.Println(load(true))  // loaded <nil>
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `:=` declares **at least one** new variable in the **current block**; other names on the left are redeclared (shadowed) in the inner block.
- Shadowing is legal but `go vet` / `shadow` analyzer often flags it — learn to spot it in `if`, `switch`, and `for` bodies.
- Named return values can be shadowed the same way inside nested blocks.
- `err` shadowing is the #1 production bug pattern in Go error handling.

## Q&A

**Q: When is `:=` safe with `err`?**  
A: When `err` has not been declared in an outer scope in the same function, e.g. `data, err := os.ReadFile(path)` at the start of a function.

**Q: How do tools catch this?**  
A: `go vet` (shadow), staticcheck, gopls diagnostics in the IDE.

**Q: Does shadowing affect the caller?**  
A: No — only the inner block sees the shadowed name. The outer `err` is unchanged, which is exactly the bug.

**Q: Edge cases?**  
A: Shadowing in `defer` closures, nested `if` chains, and `for` loops with `:=` on the same `err`.

**Q: Complexity?**  
A: O(1); this is about correctness, not performance.
