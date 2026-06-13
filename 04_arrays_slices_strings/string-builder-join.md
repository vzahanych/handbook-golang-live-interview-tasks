# string builder join

## Live interview task
Join strings with a separator using `strings.Builder` and capacity precomputation.

## Concepts covered
- strings.Builder
- preallocation
- strings

## Candidate solution

```go
package main

import (
    "fmt"
    "strings"
)

func join(parts []string, sep string) string {
    if len(parts) == 0 {
        return ""
    }
    n := len(sep) * (len(parts) - 1)
    for _, p := range parts {
        n += len(p)
    }
    var b strings.Builder
    b.Grow(n)
    b.WriteString(parts[0])
    for _, p := range parts[1:] {
        b.WriteString(sep)
        b.WriteString(p)
    }
    return b.String()
}

func main() {
    fmt.Println(join([]string{"a", "b", "c"}, ",")) // a,b,c
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `+` in loop allocates new string each time — O(n²) total bytes copied.
- `strings.Join(parts, sep)` is stdlib — implement manually to show you understand allocation.
- `Grow(n)` is hint only — Builder may still grow; `n` should be exact byte count.
- `Builder` must not be copied after first use.

## Q&A

**Q: Why not `bytes.Buffer`?**  
A: Both work; `strings.Builder` is specialized for string building (no `Write` interface overhead for string-only).

**Q: Complexity?**  
A: O(total bytes) time with one allocation when `Grow` is correct.

**Q: Single part?**  
A: Return `parts[0]` — no separator needed.

**Q: Interview follow-up?**  
A: Benchmark `+` vs `Builder` vs `Join` — classic Go perf question.
