# rune safe reverse string

## Live interview task
Reverse a UTF-8 string by runes, not bytes.

## Concepts covered
- strings
- runes
- UTF-8

## Candidate solution

```go
package main

import "fmt"

func reverseString(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

func main() {
    fmt.Println(reverseString("Go语言")) // 言语oG
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Reversing bytes of `"语言"` breaks UTF-8 — invalid string or wrong chars.
- `[]rune(s)` allocates and decodes — O(n) runes, O(n) space.
- Combining characters (emoji with ZWJ) — rune reverse may still be wrong visually; mention Unicode normalization for production i18n.
- Byte reverse of ASCII-only strings is fine.

## Q&A

**Q: Complexity?**  
A: O(n) runes time and space for the `[]rune` buffer.

**Q: Without extra `[]rune`?**  
A: Walk bytes from both ends swapping whole rune sequences — complex; `[]rune` is interview-expected.

**Q: `for _, r := range s`?**  
A: Iterates by rune, index is byte offset — use for in-place byte tricks only with care.

**Q: Edge cases?**  
A: Empty string, single rune, ASCII, combining marks, invalid UTF-8 (replacement char).
