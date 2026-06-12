# map increment efficient

## Live interview task
Use m[key]++ instead of separate read and write.

## Concepts covered
- map update
- performance idiom

## Candidate solution

```go
package main

import "fmt"

func main() {
    counts := map[string]int{}
    for _, w := range []string{"go","go","map"} { counts[w]++ }
    fmt.Println(counts)
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
