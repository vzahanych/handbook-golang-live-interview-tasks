# assignability with underlying types

## Live interview task
Show assignability and conversion rules for named types with identical underlying types.

## Concepts covered
- assignability
- underlying types
- conversions

## Candidate solution

```go
package main

import "fmt"

type Meters int
type Seconds int

func main() {
    var m Meters = 10
    var s Seconds = 2
    // m = s // compile error: different named types
    m = Meters(s)
    fmt.Println(m)
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Named types with same underlying type are **not** assignable without conversion.
- Underlying type: peel aliases with `~` in generics; for named types it's the type in `type Name Underlying`.
- Conversions between named types with identical underlying types are **allowed** if both are not defined over incompatible structs.
- Struct tags / methods differ — conversion copies bits only; no semantic safety.

## Q&A

**Q: `type Meters = int` vs `type Meters int`?**  
A: Alias is assignable to `int`; defined type requires conversion.

**Q: Why strong typing for units?**  
A: Prevents `Meters + Seconds` at compile time — use distinct defined types.

**Q: Generic constraint `~int`?**  
A: Matches any type whose underlying type is `int`, including `Meters`.

**Q: Unsafe conversion?**  
A: `unsafe.Pointer` between unrelated types is a different, dangerous topic — not allowed for arbitrary named types without `unsafe`.

**Q: Complexity?**  
A: O(1); compile-time check.
