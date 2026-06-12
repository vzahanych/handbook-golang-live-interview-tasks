# preallocate slice before append

## Live interview task
Compare appending with and without capacity preallocation.

## Concepts covered
- make slice capacity
- append
- allocations

## Candidate solution

```go
package main

import "fmt"

func build(n int) []int {
    s := make([]int, 0, n)
    for i := 0; i < n; i++ { s = append(s, i) }
    return s
}

func main() { fmt.Println(len(build(1000))) }
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
