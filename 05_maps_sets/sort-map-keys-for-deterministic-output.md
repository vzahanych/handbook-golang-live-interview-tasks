# sort map keys for deterministic output

## Live interview task
Print map entries in stable key order.

## Concepts covered
- map iteration order
- sorting

## Candidate solution

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    m := map[string]int{"b":2, "a":1, "c":3}
    keys := make([]string, 0, len(m))
    for k := range m { keys = append(keys, k) }
    sort.Strings(keys)
    for _, k := range keys { fmt.Println(k, m[k]) }
}
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
