# zero values reference types

## Live interview task
Print useful behavior of zero values for slice, map, channel, function and pointer types.

## Concepts covered
- zero values
- nil
- reference types

## Candidate solution

```go
package main

import "fmt"

func main() {
    var s []int
    var m map[string]int
    var ch chan int
    var f func()
    var p *int
    fmt.Println("slice", s == nil, len(s), cap(s)) // true 0 0
    fmt.Println("map", m == nil, len(m))           // true 0
    fmt.Println("chan", ch == nil, f == nil, p == nil)
    s = append(s, 1) // ok on nil slice
    // m["x"] = 1    // panic: assignment to entry in nil map
    fmt.Println("after append", s)
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- **Nil slice**: `len==0`, `append` works, `range` works — idiomatic empty slice is `nil` or `[]T{}` (both fine for JSON often `[]` vs `null`).
- **Nil map**: reading missing key returns zero; writing panics — use `make(map[K]V)` before insert.
- **Nil channel**: send/receive block forever — used in `select` to disable a case.
- **Nil func**: calling panics — check before call.
- **Nil pointer**: dereference panics; `== nil` is true.

## Q&A

**Q: `var s []int` vs `s := []int{}`?**  
A: Both len 0; first is nil slice, second non-nil empty slice. `reflect.DeepEqual` differs; `len` same.

**Q: Reading nil map?**  
A: `m["missing"]` returns zero value, `ok` false in two-value form — safe.

**Q: Zero value struct?**  
A: All fields zero — nested slices/maps inside still nil.

**Q: `new(T)` vs `&T{}`?**  
A: Both yield `*T` to zeroed memory; `new` only for heap allocation syntax sugar.

**Q: Interview trap?**  
A: `return nil, nil` for `([]int, error)` — nil slice is valid success result.
