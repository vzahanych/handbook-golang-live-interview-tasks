# parallel min max

## Live interview task
Compute min and max in parallel.

## Concepts covered
- parallel reduction
- min/max

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
    if len(a) == 0 { return 0, 0, false }
    workers := runtime.GOMAXPROCS(0); if workers > len(a) { workers = len(a) }
    parts := make([]pair, workers)
    var wg sync.WaitGroup
    for w:=0; w<workers; w++ { lo, hi := w*len(a)/workers, (w+1)*len(a)/workers; wg.Add(1); go func(w,lo,hi int){ defer wg.Done(); mn,mx:=a[lo],a[lo]; for _,v:=range a[lo+1:hi]{ if v<mn{mn=v}; if v>mx{mx=v} }; parts[w]=pair{mn,mx} }(w,lo,hi) }
    wg.Wait()
    mn, mx := parts[0].min, parts[0].max
    for _, p := range parts[1:] { if p.min < mn { mn = p.min }; if p.max > mx { mx = p.max } }
    return mn, mx, true
}

func main() { fmt.Println(MinMax([]int{5,1,9,3})) }
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
