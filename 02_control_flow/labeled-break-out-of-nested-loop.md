# labeled break out of nested loop

## Live interview task
Search a matrix and break out of two nested loops when a target is found.

## Concepts covered
- labeled statements
- break
- nested loops

## Candidate solution

```go
package main

import "fmt"

func find(matrix [][]int, target int) (row, col int, ok bool) {
outer:
    for i, r := range matrix {
        for j, v := range r {
            if v == target {
                row, col, ok = i, j, true
                break outer
            }
        }
    }
    return
}

func main() { fmt.Println(find([][]int{{1,2},{3,4}}, 3)) }
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
