# map increment efficient

## Live interview task
Use `m[key]++` for efficient map counter updates.

## Concepts covered
- map update idiom
- zero value read
- hash map performance

## Candidate solution

```go
package main

import "fmt"

func main() {
    counts := map[string]int{}
    for _, w := range []string{"go", "go", "map"} {
        counts[w]++
    }
    fmt.Println(counts) // map[go:2 map:1]
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Missing key read returns 0 — increment works without `if ok`.
- Not concurrent-safe — mutex or separate maps per goroutine then merge.
- For large key sets, pre-size: `make(map[string]int, expectedUnique)`.
- vs `map[string]struct{}` set — increment pattern is frequency only.

## Q&A

**Q: Decrement?**  
A: `counts[k]--` — ensure key exists or accept zero to -1.

**Q: Float counts?**  
A: Same idiom with `map[K]float64`.

**Q: Complexity?**  
A: O(1) average per increment.

**Q: Memory?**  
A: O(unique keys).

**Q: Production?**  
A: Sharded maps for concurrent counters if atomic not enough.
