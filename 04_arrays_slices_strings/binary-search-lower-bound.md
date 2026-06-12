# binary search lower bound

## Live interview task
Implement lower_bound: first index whose value is >= target.

## Concepts covered
- binary search
- index expressions

## Candidate solution

```go
package main

import "fmt"

func lowerBound(a []int, target int) int {
    lo, hi := 0, len(a)
    for lo < hi {
        mid := lo + (hi-lo)/2
        if a[mid] < target { lo = mid + 1 } else { hi = mid }
    }
    return lo
}

func main() { fmt.Println(lowerBound([]int{1,2,4,4,7}, 4)) }
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
