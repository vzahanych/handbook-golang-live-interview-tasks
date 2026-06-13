# map delete while iterate

## Live interview task
Safely delete map entries matching a predicate during iteration.

## Concepts covered
- maps
- delete builtin
- iteration semantics

## Candidate solution

```go
package main

import "fmt"

func deleteEvens(m map[int]int) {
    for k, v := range m {
        if v%2 == 0 {
            delete(m, k) // safe during range in Go
        }
    }
}

func main() {
    m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
    deleteEvens(m)
    fmt.Println(m) // map[1:1 3:3] — order nondeterministic
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Go allows **delete during range** — unspecified whether deleted entry is visited in current iteration.
- Do not assume all matching keys deleted in one pass if you depend on visit order — collect keys first if needed.
- Alternative: `for k := range m { if pred(m[k]) { delete(m, k) } }` — same semantics.
- **Never** add new keys during iteration — behavior undefined for those entries in current loop.

## Q&A

**Q: Collect keys first?**  
A: `var del []K; for k,v := range m { if pred(v) { del = append(del, k) } }; for _, k := range del { delete(m,k) }` — predictable, safe if predicate depends on others.

**Q: Complexity?**  
A: O(n) over map size.

**Q: Clear entire map?**  
A: `clear(m)` (Go 1.21+) or `for k := range m { delete(m, k) }` or reassign `m = make(...)`.

**Q: Slice delete while range?**  
A: Different — modifying slice length during range is tricky; maps are special-cased for delete.

**Q: Interview trap?**  
A: Assuming you cannot delete during iteration — in Go you can, unlike some languages.
