# remove element preserve order

## Live interview task
Remove all occurrences of a value while preserving order.

## Concepts covered
- slices
- comparable
- GC-friendly clearing

## Candidate solution

```go
package main

import "fmt"

func removeAll[T comparable](s []T, bad T) []T {
    w := 0
    for _, v := range s {
        if v != bad {
            s[w] = v
            w++
        }
    }
    var zero T
    for i := w; i < len(s); i++ { s[i] = zero } // release references
    return s[:w]
}

func main() { fmt.Println(removeAll([]string{"a","x","b","x"}, "x")) }
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
