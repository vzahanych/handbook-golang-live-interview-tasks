# sliding window max sum

## Live interview task
Find the maximum sum of any fixed-size window in a slice.

## Concepts covered
- slices
- sliding window

## Candidate solution

```go
package main

import "fmt"

func maxWindowSum(a []int, k int) (int, bool) {
    if k <= 0 || k > len(a) { return 0, false }
    sum := 0
    for _, v := range a[:k] { sum += v }
    best := sum
    for i := k; i < len(a); i++ {
        sum += a[i] - a[i-k]
        if sum > best { best = sum }
    }
    return best, true
}

func main() { fmt.Println(maxWindowSum([]int{2,1,5,1,3,2}, 3)) }
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
