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

// join concatenates parts with sep between them — one allocation when size is precomputed.
// Avoids s += in a loop, which copies the whole string on every append (O(n²) bytes).
//
// Example: join(["a","b","c"], ",") → "a,b,c"
//   n = 2 seps + len("a")+len("b")+len("c") = 2+3 = 5 bytes
func join(parts []string, sep string) string {
    if len(parts) == 0 {
        return ""
    }
    n := len(sep) * (len(parts) - 1) // separators sit between parts, not after the last
    for _, p := range parts {
        n += len(p) // total byte length of the result
    }
    var b strings.Builder
    b.Grow(n) // reserve backing buffer once (hint — may still grow if n is wrong)
    b.WriteString(parts[0])
    for _, p := range parts[1:] {
        b.WriteString(sep)
        b.WriteString(p)
    }
    return b.String() // copy buffer to an immutable string
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
