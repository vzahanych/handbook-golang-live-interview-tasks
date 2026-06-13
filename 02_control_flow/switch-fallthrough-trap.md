# switch fallthrough trap

## Live interview task
Demonstrate explicit `fallthrough` in a switch and explain why Go does not fall through by default.

## Concepts covered
- switch
- fallthrough
- case flow

## Candidate solution

```go
package main

import "fmt"

func grade(score int) string {
    switch {
    case score >= 90:
        fmt.Println("tier: A")
        fallthrough // continues into next case body
    case score >= 80:
        fmt.Println("tier: B")
    case score >= 70:
        fmt.Println("tier: C")
    default:
        fmt.Println("tier: F")
    }
    return "done"
}

func main() {
    grade(95) // prints tier: A then tier: B
    grade(85) // prints tier: B only
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Unlike C/Java, Go cases do **not** fall through unless you write `fallthrough`.
- `fallthrough` must be the last statement in a case — you cannot guard it with `if`.
- Next case condition is **not** re-evaluated — execution jumps to next case body only.
- Almost always a bug in interview code — mention you would remove `fallthrough` in production unless parsing/state machines.

## Q&A

**Q: Why did Go disable implicit fallthrough?**  
A: C-style missing `break` bugs; Go chose safety over brevity.

**Q: Expression switch with fallthrough?**  
A: Works the same: `switch n { case 1: fallthrough; case 2: ... }`.

**Q: Better pattern for overlapping tiers?**  
A: Separate logic or ordered if-else; do not rely on fallthrough for business rules.

**Q: Complexity?**  
A: O(1); readability cost is the real concern.
