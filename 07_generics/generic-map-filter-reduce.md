# generic map filter reduce

## Live interview task
Implement Map, Filter and Reduce helper functions with type parameters.

## Concepts covered
- generic functions
- type inference

## Candidate solution

```go
package main

import "fmt"

func Map[A, B any](in []A, f func(A) B) []B {
    out := make([]B, len(in))
    for i, v := range in { out[i] = f(v) }
    return out
}

func Filter[T any](in []T, keep func(T) bool) []T {
    out := in[:0]
    for _, v := range in { if keep(v) { out = append(out, v) } }
    return out
}

func Reduce[T, R any](in []T, zero R, f func(R, T) R) R {
    acc := zero
    for _, v := range in { acc = f(acc, v) }
    return acc
}

func main() { fmt.Println(Map([]int{1,2,3}, func(x int) int { return x*x })) }
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
