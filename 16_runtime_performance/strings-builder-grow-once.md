# strings builder grow once

## Live interview task
Pre-grow `strings.Builder` once to match final size and reduce reallocations.

## Concepts covered
- strings.Builder
- Grow
- allocation reduction

## Candidate solution

```go
package main

import (
    "fmt"
    "strings"
)

func repeat(word string, n int) string {
    var b strings.Builder
    b.Grow(len(word) * n)
    for i := 0; i < n; i++ {
        b.WriteString(word)
    }
    return b.String()
}

func main() {
    fmt.Println(repeat("go", 3)) // gogogo
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `Grow(n)` reserves at least n more bytes — may allocate once.
- `String()` copies buffer to immutable string — one final alloc unavoidable.
- Do not copy `Builder` after first use — internal state invalid.
- `Reset()` reuses buffer for next build in hot loops.

## Q&A

**Q: vs `bytes.Buffer`?**  
A: Builder optimized for string result; no `String()` extra copy in some paths.

**Q: Wrong Grow estimate?**  
A: Still works — extra growth if underestimate.

**Q: `WriteByte`/`WriteRune`?**  
A: Same buffer — account rune UTF-8 size in Grow.

**Q: Benchmark?**  
A: `+` loop vs Builder with `-benchmem`.

**Q: `strings.Repeat`?**  
A: Stdlib for single rune/word repeat — use when fits.
