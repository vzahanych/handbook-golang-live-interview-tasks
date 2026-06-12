# dedupe sorted slice in place

## Live interview task
Remove duplicates from a sorted slice without allocating a new backing array.

## Concepts covered
- slices
- write index
- in-place filtering

## Candidate solution

```go
package main

import "fmt"

func dedupeSorted(s []int) []int {
    if len(s) < 2 { return s }
    w := 1
    for r := 1; r < len(s); r++ {
        if s[r] != s[w-1] {
            s[w] = s[r]
            w++
        }
    }
    return s[:w]
}

func main() { fmt.Println(dedupeSorted([]int{1,1,2,2,2,3})) }
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
