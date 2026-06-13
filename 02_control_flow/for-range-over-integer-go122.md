# for range over integer go122

## Live interview task
Print indexes 0..n-1 using the Go 1.22 integer range form.

## Concepts covered
- for range
- integer range (`for i := range n`, Go 1.22+)
- per-iteration loop variables (Go 1.22+ — see [loop-variable-closure-capture](../01_language_basics/loop-variable-closure-capture.md))
- iterator range (`for v := range seq`, `iter.Seq` / `iter.Seq2`, Go 1.23+ — see [for-range-over-iter-seq-go123](./for-range-over-iter-seq-go123.md))
- stdlib iterators (`slices.*`, `maps.*`, Go 1.23+; `reflect.*` iterators, Go 1.23–1.26)

## Candidate solution

```go
package main

import "fmt"

func main() {
    const n = 5
    for i := range n {
        fmt.Println(i) // 0 1 2 3 4
    }
}
```

## Equivalent pre-1.22

```go
for i := 0; i < n; i++ {
    fmt.Println(i)
}
```

## Run

```bash
go run . # Go 1.22+ (integer range); tested on Go 1.26.x
```

## Interview notes / pitfalls
- `for i := range n` iterates `i` from `0` to `n-1` when `n` is a non-negative integer.
- Negative `n` — loop runs zero times (same as `for i := 0; i < n; i++` when n < 0).
- Do not confuse with ranging over a channel or map — integer range is index-only.
- `for range 3` is idiomatic for "do something three times" when you do not need the index.
- Go 1.22 also changed **loop variable semantics** (not syntax): each iteration gets its own `i` / `v` in every `for` and `for range` loop. That fixes closure/goroutine capture bugs; integer range follows the same rule.

### Other `for range` forms on Go 1.26.x (today)

| Since | Range expression | Example |
|-------|------------------|---------|
| always | slice, array, string, map, channel | `for i, v := range xs` |
| 1.22 | non-negative **integer** | `for i := range n` |
| 1.23 | **`iter.Seq[V]`** / **`iter.Seq2[K,V]`** | `for v := range count(n)` |
| 1.23 | **`slices` / `maps` helpers** | `for k := range maps.Keys(m)`; `for i, v := range slices.Backward(xs)` |
| 1.23 | **`reflect.Value.Seq` / `Seq2`** | `for v := range reflect.ValueOf(xs).Seq()` on slice/array/map/string/chan |
| 1.26 | **`reflect.Type/Value` metadata iterators** | `for field, val := range reflect.ValueOf(p).Fields()`; `for m := range t.Methods()` |

No new **language** range syntax after 1.23 — 1.24–1.26 add more stdlib values you can pass to `range` (mostly `reflect`). Full iterator example: [for-range-over-iter-seq-go123](./for-range-over-iter-seq-go123.md).

## Q&A

**Q: Can you write `for range n` with a variable declared outside?**  
A: No — `range n` always declares a new `i` (or use `_` if unused: `for range n`).

**Q: Float or `int64`?**  
A: Integer range requires an integer type; use a classic for loop for non-int bounds.

**Q: Why was this added?**  
A: Common pattern `for i := 0; i < n; i++` is verbose; aligns with other languages' `for i in range(n)`.

**Q: Complexity?**  
A: O(n) iterations, O(1) space.

**Q: Go version in interview?**  
A: On Go 1.26.x you have all of the above. Integer range needs 1.22+; iterator range (`iter.Seq`) needs 1.23+. Confirm the module `go` line — older `go` versions disable newer range forms even if the toolchain is newer.
