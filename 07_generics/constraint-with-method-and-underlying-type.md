# constraint with method and underlying type

## Live interview task
Use a constraint that requires both an underlying type and a method.

## Concepts covered
- interface constraints
- type unions
- methods in constraints

## Candidate solution

```go
package main

import "fmt"

type ID int

func (id ID) String() string { return fmt.Sprintf("ID-%d", id) }

type PrintableInt interface {
    ~int
    String() string
}

func PrintPlusOne[T PrintableInt](v T) string {
    return fmt.Sprintf("%s next=%d", v.String(), int(v)+1)
}

func main() {
    fmt.Println(PrintPlusOne(ID(41))) // ID-41 next=42
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Constraint is intersection: underlying `~int` **and** method `String() string`.
- Plain `int` fails — no `String()` method unless you add one.
- `int(v)` conversion works because `~int` in constraint permits conversion to int.
- Go 1.18+ "general interfaces" in constraints — not same as runtime interface.

## Q&A

**Q: Union + methods?**  
A: `interface { ~int | ~string; String() string }` — both types need method.

**Q: `comparable` constraint?**  
A: Predeclared — types that support `==`.

**Q: `any` constraint?**  
A: No restrictions — like unconstrained type param.

**Q: Why methods in constraint?**  
A: Call `v.String()` inside generic fn without type switch.

**Q: stdlib example?**  
A: `slog.LogValuer`, `fmt.Stringer` in custom generic logging helpers.
