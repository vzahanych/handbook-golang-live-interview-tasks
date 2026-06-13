# iota bitmask permissions

## Live interview task
Implement bitmask permissions with iota and test add/remove/has operations.

## Concepts covered
- constant declarations
- iota
- bitwise operators

## Candidate solution

```go
package main

import "fmt"

type Perm uint8

const (
    Read Perm = 1 << iota // 1
    Write                 // 2
    Execute               // 4
)

func Has(mask, p Perm) bool { return mask&p != 0 }
func Add(mask, p Perm) Perm { return mask | p }
func Remove(mask, p Perm) Perm { return mask &^ p }

func main() {
    var p Perm
    p = Add(p, Read|Write)
    fmt.Println(Has(p, Read), Has(p, Execute)) // true false
    p = Remove(p, Write)
    fmt.Println(Has(p, Write)) // false
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `iota` resets to 0 in each `const` block; start a new `const (` block to reset.
- Use `1 << iota` for powers of two so each flag is a distinct bit.
- `&^` is bit clear (AND NOT): `mask &^ p` removes bit `p`.
- Checking `mask&p != 0` works for single-bit flags; for multi-bit `p`, `mask&p == p` means "has all bits in p".
- `uint8` gives 8 bits — fine for demo; use `uint32`/`uint64` for more flags.

## Q&A

**Q: What is `iota`?**  
A: A predeclared identifier that increments in a `const` block, starting at 0 for the first line.

**Q: How do you skip an iota value?**  
A: Use a blank identifier: `const ( A = 1 << iota; _; C )` — the `_` line still consumes an iota step.

**Q: Why not use a `map[string]bool` instead?**  
A: Bitmasks are O(1) space per flag, cache-friendly, and fast for set ops (`|`, `&`, `&^`). Maps are clearer when flags are dynamic or sparse.

**Q: Edge cases to test?**  
A: Empty mask (no permissions), add same flag twice (idempotent), remove flag not present, combine `Read|Write|Execute`, overflow if too many flags for the integer width.

**Q: Complexity?**  
A: Has/Add/Remove are O(1) time and space.
