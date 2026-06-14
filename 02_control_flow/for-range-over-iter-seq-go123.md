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

// count returns a lazy iterator over 0..n-1 (not a slice).
// for v := range count(n) calls yield(i) each step; yield returns false on break.
func count(n int) iter.Seq[int] {
    return func(yield func(int) bool) {
        for i := range n {
            if !yield(i) { // consumer stopped — exit early
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

## How the lazy iterator works

An `iter.Seq[V]` is just `func(yield func(V) bool)`. There is no slice, no
channel, and no goroutine — it is a single function that *drives the loop body*.

**Inverted control flow (push model).** With a normal `for` you call into the
data ("give me the next element"). Here the iterator calls *you*: each element
is **pushed** to the loop body through `yield`. The compiler rewrites

```go
for v := range count(3) { use(v) }
```

into roughly:

```go
count(3)(func(v int) bool {   // the loop body becomes the yield callback
    use(v)
    return true               // true = keep going, false = stop (break)
})
```

So `count` never builds `[]int{0,1,2}`. It runs its own `for i := range n`
loop and hands one value at a time to `yield`. The value `2` is only ever
computed *after* the body has finished with `0` and `1`.

**Why this is "lazy".** Each element is produced on demand, right before the
body consumes it, and is discarded after. Concretely this means:

- **O(1) extra memory** — nothing materializes the whole sequence. `count(1_000_000)` allocates nothing; a `[]int` of a million ints would allocate ~8 MB.
- **Infinite / unbounded sources work.** An iterator can `yield` forever (e.g. a counter or a paginated API); the consumer decides when to stop. A slice cannot be infinite.
- **Short-circuiting skips work.** `break` (or `return false` from `yield`) stops the iterator *immediately* — remaining elements are never generated. Find-first over a huge source touches only what it needs.

**Tracing `count(5)` with an early `break`:**

```go
for v := range count(5) {
    if v == 2 { break }
    fmt.Print(v, " ")   // prints: 0 1
}
```

1. `count` starts its loop, computes `i=0`, calls `yield(0)` → body prints `0`, returns `true`.
2. `i=1`, `yield(1)` → prints `1`, returns `true`.
3. `i=2`, `yield(2)` → body hits `break`, so `yield` returns `false`.
4. `count` sees `!yield(...)` and `return`s. **`i=3` and `i=4` are never computed.**

That `return false` propagation is the whole reason every iterator checks
`if !yield(v) { return }`. Forget it and a `break` cannot stop the source —
and worse, calling `yield` again after it returned `false` **panics**.

## Interview notes / pitfalls
- An iterator is a function that calls `yield` for each element; return `false` from `yield` to stop early (same as `break` in `for range`).
- `yield` must not be called after it returns `false` — it panics.
- `iter.Seq[V]` → `for v := range seq`; `iter.Seq2[K,V]` → `for k, v := range seq`.
- Prefer stdlib helpers when they exist — see [for-range-stdlib-iterators-go123](./for-range-stdlib-iterators-go123.md).
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
