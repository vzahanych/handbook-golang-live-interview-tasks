# type switch format any

## Live interview task
Use a type switch to format values of different dynamic types.

## Concepts covered
- interfaces
- type switches
- dynamic type
- nil interface gotcha

## Candidate solution

```go
package main

import "fmt"

func describe(v any) string {
    switch x := v.(type) {
    case nil:
        return "nil"
    case int:
        return fmt.Sprintf("int:%d", x)
    case string:
        return fmt.Sprintf("string:%q", x)
    case fmt.Stringer:
        return "stringer:" + x.String()
    default:
        return fmt.Sprintf("%T", x)
    }
}

func main() {
    fmt.Println(describe("go"))
    fmt.Println(describe(42))
    fmt.Println(describe(nil))
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `case nil` matches an interface with **no dynamic type** (untyped nil or empty interface holding no value).
- An interface holding a **typed nil pointer** (e.g. `(*int)(nil)`) is **not** `case nil` — it falls through to `default` or a pointer case.
- Case order: concrete types before interfaces — if `fmt.Stringer` were before `string`, strings would match `Stringer` (strings implement it).
- `x` in `case T:` has type `T` inside that branch.

## Q&A

**Q: `any` vs `interface{}`?**  
A: Identical — `any` is a predeclared alias (Go 1.18+).

**Q: Two-type assert `v.(type)`?**  
A: Only legal in a type switch. Elsewhere use `x, ok := v.(int)`.

**Q: How to handle multiple types the same way?**  
A: `case int, int64, int32:` shared body.

**Q: Complexity?**  
A: O(1) per switch; type switch uses runtime type info.

**Q: JSON unmarshaling pattern?**  
A: `switch v := raw.(type)` on `map[string]any` values — very common in live coding.
