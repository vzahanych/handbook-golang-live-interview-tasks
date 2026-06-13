# switch without expression classifier

## Live interview task
Classify an integer using a switch without an expression (boolean switch).

## Concepts covered
- switch statements
- tagless switch (true switch)
- case order

## Candidate solution

```go
package main

import "fmt"

func classify(n int) string {
    switch {
    case n < 0:
        return "negative"
    case n == 0:
        return "zero"
    case n%2 == 0:
        return "positive even"
    default:
        return "positive odd"
    }
}

func main() {
    fmt.Println(classify(7))  // positive odd
    fmt.Println(classify(-1)) // negative
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Tagless `switch` is equivalent to `switch true` — cases are boolean expressions evaluated top to bottom.
- **First matching case wins** — order matters: put `n < 0` before `n%2 == 0`.
- No automatic fallthrough (unlike C) — each case ends with implicit `break`.
- Use `fallthrough` keyword explicitly to run the next case (rare, often a smell).

## Q&A

**Q: Switch vs if-else chain?**  
A: Style preference for many branches; switch can be slightly clearer for disjoint cases. Compiler lowers both similarly.

**Q: Can cases call functions?**  
A: Yes — `case isPrime(n):` is valid if `isPrime` returns bool.

**Q: Type switch difference?**  
A: `switch x := v.(type)` inspects dynamic type, not boolean conditions.

**Q: Complexity?**  
A: O(1) for fixed number of cases; O(1) space.

**Q: Production tip?**  
A: For enums, prefer `switch status { case Pending: ... }` with a typed constant — exhaustiveness checked by `exhaustive` linter.
