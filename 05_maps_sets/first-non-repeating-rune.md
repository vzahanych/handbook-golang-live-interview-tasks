# first non repeating rune

## Live interview task
Find the first non-repeating rune in a string.

## Concepts covered
- map[rune]int
- range over string

## Candidate solution

```go
package main

import "fmt"

func firstUnique(s string) (rune, bool) {
    counts := make(map[rune]int)
    for _, r := range s { counts[r]++ }
    for _, r := range s {
        if counts[r] == 1 { return r, true }
    }
    return 0, false
}

func main() { r, ok := firstUnique("swiss"); fmt.Printf("%c %v\n", r, ok) }
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
