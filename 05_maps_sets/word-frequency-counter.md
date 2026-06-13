# word frequency counter

## Live interview task
Count words in text using a map.

## Concepts covered
- maps
- zero value on lookup
- strings.Fields

## Candidate solution

```go
package main

import (
    "fmt"
    "strings"
)

func freq(text string) map[string]int {
    m := make(map[string]int)
    for _, w := range strings.Fields(strings.ToLower(text)) {
        m[w]++ // missing key reads as 0, then stores count
    }
    return m
}

func main() {
    fmt.Println(freq("Go go channels maps")) // map[channels:1 go:2 maps:1]
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `m[w]++` works because read of missing key returns zero value `0`.
- `strings.Fields` splits on Unicode whitespace — punctuation sticks to words (`"go,"` ≠ `"go"`).
- Map iteration order is **randomized** — sort keys for stable output.
- For large texts: stream lines instead of loading all into memory.

## Q&A

**Q: Complexity?**  
A: O(n) words, O(u) space for u unique words.

**Q: Case sensitivity?**  
A: `ToLower` for case-insensitive; use `strings.EqualFold` when comparing pairs.

**Q: Better tokenization?**  
A: `strings.FieldsFunc` with custom rune predicate; or regexp for punctuation strip.

**Q: Thread-safe?**  
A: No — use `sync.Mutex` around map or `sync.Map` for read-heavy caches.

**Q: Edge cases?**  
A: Empty string, only whitespace, single word repeated.
