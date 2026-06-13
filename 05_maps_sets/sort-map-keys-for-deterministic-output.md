# sort map keys for deterministic output

## Live interview task
Print map entries in stable key order.

## Concepts covered
- map iteration order
- sorting
- deterministic output

## Candidate solution

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    m := map[string]int{"b": 2, "a": 1, "c": 3}
    keys := make([]string, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    for _, k := range keys {
        fmt.Println(k, m[k])
    }
}
```

## Run

```bash
go run .
```

## Expected output (always)

```
a 1
b 2
c 3
```

## Interview notes / pitfalls
- Map iteration order is **intentionally randomized** (since Go 1) — never depend on it for logic or tests without sorting.
- Collect keys → sort → lookup values — standard pattern.
- `slices.Sorted(maps.Keys(m))` (Go 1.23+) simplifies key collection.
- JSON `Marshal` sorts map keys for stable output.

## Q&A

**Q: Why random iteration?**  
A: Prevent reliance on implementation detail; encourage explicit ordering.

**Q: Complexity?**  
A: O(k log k) for k keys to sort.

**Q: Sort by value?**  
A: Sort slice of structs `{k, v}` with `sort.Slice` by `.v`.

**Q: Tests comparing maps?**  
A: `reflect.DeepEqual` or `maps.Equal` — order irrelevant for equality.

**Q: Edge cases?**  
A: Empty map — no output; single key.
