# generic map filter reduce

## Live interview task
Implement Map, Filter and Reduce helper functions with type parameters.

## Concepts covered
- generic functions
- type inference
- in-place filter

## Candidate solution

```go
package main

import "fmt"

func Map[A, B any](in []A, f func(A) B) []B {
    out := make([]B, len(in))
    for i, v := range in {
        out[i] = f(v)
    }
    return out
}

func Filter[T any](in []T, keep func(T) bool) []T {
    out := in[:0]
    for _, v := range in {
        if keep(v) {
            out = append(out, v)
        }
    }
    return out
}

func Reduce[T, R any](in []T, zero R, f func(R, T) R) R {
    acc := zero
    for _, v := range in {
        acc = f(acc, v)
    }
    return acc
}

func main() {
    fmt.Println(Map([]int{1, 2, 3}, func(x int) int { return x * x }))
    fmt.Println(Filter([]int{1, 2, 3, 4}, func(x int) bool { return x%2 == 0 }))
    fmt.Println(Reduce([]int{1, 2, 3}, 0, func(a, x int) int { return a + x }))
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `Filter` reuses `in[:0]` backing array — mutates original slice header if caller keeps `in` reference.
- Type inference: `Map([]int{1}, fn)` infers `A=int`; sometimes need explicit `Map[int,int](...)`.
- Go 1.23+ `slices` has `DeleteFunc`, `Compact` — Map/Filter less hand-rolled in prod.
- `Reduce` on empty returns `zero` — document identity element.

## Q&A

**Q: Map complexity?**  
A: O(n) time, O(n) output space.

**Q: Filter without mutating input?**  
A: `out := make([]T, 0, len(in))` instead of `in[:0]`.

**Q: Lazy map?**  
A: Generators/channels — `func MapChan` for streaming.

**Q: Why generics over `interface{}`?**  
A: Type safety, no assertions, better compile-time checks.

**Q: Parallel Map?**  
A: Worker pool with order preservation — separate interview topic.
