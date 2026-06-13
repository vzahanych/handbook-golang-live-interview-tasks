# type alias vs new defined type

## Live interview task
Explain the difference between a type alias and a new defined type.

## Concepts covered
- type definitions
- alias declarations
- type identity

## Candidate solution

```go
package main

import "fmt"

type UserID int       // new defined type — distinct from int
type InternalID = int // type alias — identical to int

func printInt(n int) { fmt.Println(n) }

func main() {
    var u UserID = 10
    var i InternalID = 20
    printInt(i)      // ok: alias is int
    printInt(int(u)) // explicit conversion required: UserID ≠ int
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `type T int` creates a **new type** with `int` as underlying type — no implicit conversion.
- `type T = int` is an **alias** — `T` and `int` are identical for assignability, methods, and interfaces.
- Aliases are mainly for migrations (`type Request = http.Request`) or embedding legacy APIs.
- Defined types enable stronger typing (`UserID` vs `OrderID`) and distinct method sets.

## Q&A

**Q: Can `UserID` and `int` be compared with `==` without conversion?**  
A: No — they are different types. Use `UserID(42) == u` or convert one side.

**Q: Which one can add methods?**  
A: Both — but alias methods are the same as methods on the aliased type; defined types get their own method set.

**Q: When did aliases appear?**  
A: Go 1.9, primarily to ease gradual refactors when moving types between packages.

**Q: `byte` vs `uint8`?**  
A: `byte` is a predefined alias for `uint8` — same type identity.

**Q: Interview trap?**  
A: `type JSON = map[string]any` — still a map; mutations and assignments behave exactly like the aliased type.
