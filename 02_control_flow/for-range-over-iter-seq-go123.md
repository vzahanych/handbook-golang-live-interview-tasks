# for range over iter seq go123

## Live interview task
Implement a custom iterator with `iter.Seq` / `iter.Seq2` and consume it with `for range`.

## Concepts covered
- for range
- iter.Seq / iter.Seq2
- yield callback and early stop
- Go 1.23+

## Candidate solution

```go
package main

import (
    "fmt"
    "iter"
    "maps"
)

func count(n int) iter.Seq[int] {
    return func(yield func(int) bool) {
        for i := range n {
            if !yield(i) {
                return
            }
        }
    }
}

func pairs(keys ...string) iter.Seq2[int, string] {
    return func(yield func(int, string) bool) {
        for i, k := range keys {
            if !yield(i, k) {
                return
            }
        }
    }
}

func main() {
    fmt.Print("Seq: ")
    for v := range count(5) {
        fmt.Print(v, " ")
    }
    fmt.Println()

    fmt.Print("Seq2: ")
    for i, k := range pairs("go", "iter", "range") {
        fmt.Printf("(%d:%s) ", i, k)
    }
    fmt.Println()

    m := map[string]int{"a": 1, "b": 2}
    fmt.Print("maps.Keys: ")
    for k := range maps.Keys(m) {
        fmt.Print(k, " ")
    }
    fmt.Println()
}
```

## Run

```bash
go run . # Go 1.23+
```

## Expected output

```
Seq: 0 1 2 3 4 
Seq2: (0:go) (1:iter) (2:range) 
maps.Keys: a b 
```

(`maps.Keys` order is nondeterministic.)

## Interview notes / pitfalls
- An iterator is a function that calls `yield` for each element; return `false` from `yield` to stop early (same as `break` in `for range`).
- `yield` must not be called after it returns `false` — it panics.
- `iter.Seq[V]` → `for v := range seq`; `iter.Seq2[K,V]` → `for k, v := range seq`.
- Prefer stdlib helpers when they exist: `maps.Keys`, `slices.Backward`, `slices.All`, etc.
- Previewed in Go 1.22 behind `GOEXPERIMENT=rangefunc`; stable language feature from **Go 1.23**.
- See also: [for-range-over-integer-go122](./for-range-over-integer-go122.md) for `for i := range n`.

## Q&A

**Q: How is this different from integer range?**  
A: Integer range is built-in syntax for `0..n-1`. Iterator range works over any `iter.Seq` / `iter.Seq2` — custom sequences, lazy pipelines, stdlib adapters.

**Q: Why not return a slice?**  
A: Iterators can be lazy (generate on demand), composable, and stoppable without allocating the full sequence.

**Q: Can you range over a plain func?**  
A: Only functions with the iterator signature (`func(yield func(V) bool)` or the two-value form). Random func types do not work.

**Q: Go version in interview?**  
A: Needs Go 1.23+ in `go.mod`. On Go 1.26.x the feature is always available when the language version allows it.

**Q: Complexity?**  
A: O(n) if the iterator yields n elements; O(1) extra space when lazy (no materialized slice).
