# parallel min max

## Live interview task
Compute min and max in parallel using per-chunk reduction.

## Concepts covered
- parallel reduction
- min/max merge

## Candidate solution

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
)

type pair struct{ min, max int }

func MinMax(a []int) (int, int, bool) {
    if len(a) == 0 {
        return 0, 0, false
    }
    workers := runtime.GOMAXPROCS(0)
    if workers > len(a) {
        workers = len(a)
    }
    parts := make([]pair, workers)
    var wg sync.WaitGroup

    for w := 0; w < workers; w++ {
        lo := w * len(a) / workers
        hi := (w + 1) * len(a) / workers
        wg.Add(1)
        go func(w, lo, hi int) {
            defer wg.Done()
            mn, mx := a[lo], a[lo]
            for _, v := range a[lo+1 : hi] {
                if v < mn {
                    mn = v
                }
                if v > mx {
                    mx = v
                }
            }
            parts[w] = pair{mn, mx}
        }(w, lo, hi)
    }
    wg.Wait()

    mn, mx := parts[0].min, parts[0].max
    for _, p := range parts[1:] {
        if p.min < mn {
            mn = p.min
        }
        if p.max > mx {
            mx = p.max
        }
    }
    return mn, mx, true
}

func main() {
    fmt.Println(MinMax([]int{5, 1, 9, 3})) // 1 9 true
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Initialize chunk min/max from `a[lo]` — empty chunk impossible with fair split.
- Final merge O(workers) — cheap vs O(n) scan.
- Single pass per chunk finds both min and max — 2n comparisons total sequential equivalent.

## Q&A

**Q: One element?**  
A: min == max == that element.

**Q: Generic version?**  
A: Constraint `cmp.Ordered` and `cmp.Compare` in loop.

**Q: `slices.Min`/`Max`?**  
A: Sequential stdlib — mention for production small slices.

**Q: Complexity?**  
A: O(n) comparisons, O(workers) space.

**Q: NaN floats?**  
A: `cmp.Compare` handles ordering rules — discuss if float interview.
