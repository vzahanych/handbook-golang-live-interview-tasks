# for range stdlib iterators go123

## Live interview task
Use `for range` over standard-library iterator helpers from `slices`, `maps`, and `reflect`.

## Concepts covered
- for range
- slices.All / Values / Backward
- maps.All / Keys / Values
- reflect.Value.Seq / Seq2 / Fields
- Go 1.23+ (`reflect` metadata iterators from 1.26)

## Candidate solution

```go
package main

import (
    "fmt"
    "maps"
    "reflect"
    "slices"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    xs := []int{10, 20, 30}

    fmt.Print("slices.Backward: ")
    for i, v := range slices.Backward(xs) {
        fmt.Printf("(%d:%d) ", i, v)
    }
    fmt.Println()

    fmt.Print("slices.Values: ")
    for v := range slices.Values(xs) {
        fmt.Print(v, " ")
    }
    fmt.Println()

    m := map[string]int{"go": 1, "iter": 2}
    fmt.Print("maps.All: ")
    for k, v := range maps.All(m) {
        fmt.Printf("%s=%d ", k, v)
    }
    fmt.Println()

    fmt.Print("reflect.Value.Seq2: ")
    for i, val := range reflect.ValueOf(xs).Seq2() {
        fmt.Printf("(%v:%v) ", i.Interface(), val.Interface())
    }
    fmt.Println()

    p := Person{Name: "Ann", Age: 30}
    fmt.Print("reflect.Value.Fields: ")
    for field, val := range reflect.ValueOf(p).Fields() {
        fmt.Printf("%s=%v ", field.Name, val.Interface())
    }
    fmt.Println()
}
```

## Run

```bash
go run . # Go 1.23+; reflect.Value.Fields needs Go 1.26+
```

## Expected output

```
slices.Backward: (2:30) (1:20) (0:10) 
slices.Values: 10 20 30 
maps.All: go=1 iter=2 
reflect.Value.Seq2: (0:10) (1:20) (2:30) 
reflect.Value.Fields: Name=Ann Age=30 
```

(`maps.All` key order is nondeterministic.)

## Interview notes / pitfalls
- These functions return `iter.Seq` / `iter.Seq2` — same `for range` syntax as custom iterators ([for-range-over-iter-seq-go123](./for-range-over-iter-seq-go123.md)).
- **`slices`**: `All` → index+value; `Values` → values only; `Backward` → reverse index+value. Also `Collect`, `AppendSeq`, `Sorted`, `Chunk`, …
- **`maps`**: `All` → key+value; `Keys` / `Values` → one element per step. Map iteration order still random.
- **`reflect.Value.Seq`**: on a slice, yields **indices** only; use **`Seq2`** for index+value pairs (same as ranging the slice directly).
- **`reflect.Value.Fields`** (Go 1.26): struct fields as `(StructField, Value)` pairs — replaces manual `for i := range v.NumField()` loops.
- Go 1.26 also adds `Type.Fields`, `Type.Methods`, `Type.Ins`, `Type.Outs`, `Value.Methods`.

## Q&A

**Q: Why use `slices.Backward` instead of a classic loop?**  
A: Same behavior, composes with other iterator helpers, and reads as “reverse this sequence” without index arithmetic.

**Q: `maps.Keys(m)` vs `for k := range m`?**  
A: Equivalent consumption; iterator form plugs into `slices.Collect`, pipelines, and functions that take `iter.Seq`.

**Q: When to use `reflect.Value.Seq`?**  
A: When the value kind is only known at run time — generic reflection code that must iterate slice/array/map/string/chan without a type switch per kind.

**Q: Go version in interview?**  
A: `slices`/`maps` iterators need 1.23+. `Value.Fields` and related `Type.*` iterators need 1.26+.

**Q: Complexity?**  
A: Same as the underlying collection — O(n) over n elements; iterators are lazy views, not extra copies of the whole slice/map.
