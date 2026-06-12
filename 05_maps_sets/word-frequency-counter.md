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
        m[w]++
    }
    return m
}

func main() { fmt.Println(freq("Go go channels maps")) }
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
