# benchmark string concat

## Live interview task
Benchmark string concatenation variants.

## Concepts covered
- benchmarks
- strings.Builder

## Candidate solution

```go
package concat

import (
    "strings"
    "testing"
)

var sink string

func BenchmarkBuilder(b *testing.B) {
    parts := []string{"a","b","c","d"}
    for i := 0; i < b.N; i++ {
        var sb strings.Builder
        for _, p := range parts { sb.WriteString(p) }
        sink = sb.String()
    }
}
```

## Run

```bash
go test -bench=. -benchmem
```

## Interview notes / pitfalls
- None specific; discuss edge cases and complexity.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
