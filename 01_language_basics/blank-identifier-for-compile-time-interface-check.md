# blank identifier for compile time interface check

## Live interview task
Use the blank identifier for a compile-time interface implementation assertion.

## Concepts covered
- blank identifier
- interfaces
- method sets

## Candidate solution

```go
package main

import "fmt"

type Stringer interface{ String() string }

type User struct{ Name string }

func (u User) String() string { return u.Name }

// Compile-time check: User must implement Stringer.
var _ Stringer = User{}

// Pointer receiver check (when methods use *User):
// var _ Stringer = (*User)(nil)

func main() { fmt.Println(User{"Ada"}) }
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `var _ Stringer = User{}` fails at **compile** time if `User` stops implementing `Stringer` — no runtime cost.
- Value vs pointer: if methods are on `*User`, assert with `(*User)(nil)` or `new(User)`.
- Do not confuse with type assertion `v.(Stringer)` — that is runtime.
- Place assertions near the type definition (often end of file) so refactors break early.

## Q&A

**Q: Why `var _ T = concrete{}` instead of a test?**  
A: Tests can be skipped or live in another package; this guarantees every build verifies the contract.

**Q: What if the interface is in another module?**  
A: Same pattern works — import the interface package and assert in the implementing package.

**Q: Can you assert multiple interfaces?**  
A: Yes — one line per interface: `var _ io.Reader = (*MyType)(nil)`.

**Q: Production tip?**  
A: Also document which methods are exported API vs interface satisfaction; use `go:generate` or linters (`implements`) in large codebases.

**Q: Complexity?**  
A: Zero runtime cost; compile-time only.
