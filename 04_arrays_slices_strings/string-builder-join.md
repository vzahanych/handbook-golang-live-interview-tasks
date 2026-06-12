# string builder join

## Live interview task
Join strings with a separator using strings.Builder and capacity precomputation.

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
    if len(parts) == 0 { return "" }
    n := len(sep) * (len(parts)-1)
    for _, p := range parts { n += len(p) }
    var b strings.Builder
    b.Grow(n)
    b.WriteString(parts[0])
    for _, p := range parts[1:] {
        b.WriteString(sep)
        b.WriteString(p)
    }
    return b.String()
}

func main() { fmt.Println(join([]string{"a","b","c"}, ",")) }
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- None specific; discuss edge cases and complexity.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
