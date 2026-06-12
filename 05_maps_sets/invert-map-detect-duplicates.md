# invert map detect duplicates

## Live interview task
Invert a map while detecting duplicate values.

## Concepts covered
- maps
- generics
- comparable
- errors

## Candidate solution

```go
package main

import (
    "errors"
    "fmt"
)

func invert[K comparable, V comparable](m map[K]V) (map[V]K, error) {
    out := make(map[V]K, len(m))
    for k, v := range m {
        if _, exists := out[v]; exists { return nil, errors.New("duplicate value") }
        out[v] = k
    }
    return out, nil
}

func main() { fmt.Println(invert(map[string]int{"a":1,"b":2})) }
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
