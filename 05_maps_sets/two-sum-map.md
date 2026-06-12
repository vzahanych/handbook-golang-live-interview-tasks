# two sum map

## Live interview task
Return indexes of two values whose sum equals target.

## Concepts covered
- maps
- lookup comma-ok

## Candidate solution

```go
package main

import "fmt"

func twoSum(a []int, target int) (int, int, bool) {
    seen := make(map[int]int, len(a))
    for i, v := range a {
        if j, ok := seen[target-v]; ok { return j, i, true }
        seen[v] = i
    }
    return 0, 0, false
}

func main() { fmt.Println(twoSum([]int{2,7,11,15}, 9)) }
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
