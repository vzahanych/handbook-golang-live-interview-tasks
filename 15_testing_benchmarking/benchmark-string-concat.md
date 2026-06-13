# benchmark string concat

## Live interview task
Benchmark string concatenation — compare `+` loop vs `strings.Builder`.

## Concepts covered
- testing.B
- benchmarks
- benchmem

## Candidate solution

```go
package concat

import (
    "strings"
    "testing"
)

var sink string

func concatPlus(parts []string) string {
    s := ""
    for _, p := range parts {
        s += p
    }
    return s
}

func concatBuilder(parts []string) string {
    var b strings.Builder
    b.Grow(16)
    for _, p := range parts {
        b.WriteString(p)
    }
    return b.String()
}

func BenchmarkPlus(b *testing.B) {
    parts := []string{"a", "b", "c", "d", "e"}
    for i := 0; i < b.N; i++ {
        sink = concatPlus(parts)
    }
}

func BenchmarkBuilder(b *testing.B) {
    parts := []string{"a", "b", "c", "d", "e"}
    for i := 0; i < b.N; i++ {
        sink = concatBuilder(parts)
    }
}
```

## Run

```bash
go test -bench=. -benchmem
```

## Interview notes / pitfalls
- Assign to package `sink` — prevents compiler dead-code elimination.
- `+` in loop is O(n²) copies — Builder O(n).
- `-benchmem` shows allocations per op — key metric in interviews.
- `b.ResetTimer()` after expensive setup outside loop.

## Q&A

**Q: `b.N`?**  
A: Framework adjusts until stable timing — don't hardcode iterations.

**Q: `strings.Join`?**  
A: Best when you have slice of strings already — one allocation with size hint.

**Q: Compare results?**  
A: `benchstat` for A/B across commits.

**Q: Parallel benchmark?**  
A: `b.RunParallel` for concurrent code paths.

**Q: Expected winner?**  
A: Builder/Join — fewer allocs and faster for n>2.
