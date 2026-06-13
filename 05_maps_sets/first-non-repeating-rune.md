# first non repeating rune

## Live interview task
Find the first non-repeating rune in a string.

## Concepts covered
- map[rune]int
- two-pass counting
- range over string

## Candidate solution

```go
package main

import "fmt"

func firstUnique(s string) (rune, bool) {
    counts := make(map[rune]int)
    for _, r := range s {
        counts[r]++
    }
    for _, r := range s {
        if counts[r] == 1 {
            return r, true
        }
    }
    return 0, false
}

func main() {
    r, ok := firstUnique("swiss")
    fmt.Printf("%c %v\n", r, ok) // w true
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- First pass counts, second pass finds first with count 1 — preserves left-to-right order.
- `range` over string yields runes, not bytes.
- Return `(0, false)` for empty/no unique — `0` may be ambiguous if `'0'` is valid; use `(rune, bool)` or error.
- Single-pass with ordered map possible (`linked` + map) — overkill unless asked.

## Q&A

**Q: Complexity?**  
A: O(n) runes time, O(k) space for k distinct runes.

**Q: One pass?**  
A: Need order — `[]rune` + count map, or index map `map[rune]int` storing first index and count.

**Q: ASCII only?**  
A: `map[byte]int` + byte loop if valid assumption.

**Q: All repeating?**  
A: Return `false` — e.g. `"ss"`.

**Q: LeetCode link?**  
A: First unique character (387) — same idea.
