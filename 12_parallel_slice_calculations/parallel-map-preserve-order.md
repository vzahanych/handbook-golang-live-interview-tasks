# parallel map preserve order

## Live interview task
Apply a function to each slice element in parallel while preserving output order.

## Concepts covered
- parallel map
- order preservation
- index-based work queue

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

func ParallelMap[A, B any](in []A, workers int, f func(A) B) []B {
    out := make([]B, len(in))
    jobs := make(chan int)
    var wg sync.WaitGroup

    for w := 0; w < workers; w++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for i := range jobs {
                out[i] = f(in[i])
            }
        }()
    }

    for i := range in {
        jobs <- i
    }
    close(jobs)
    wg.Wait()
    return out
}

func main() {
    fmt.Println(ParallelMap([]int{1, 2, 3}, 2, func(x int) int { return x * x }))
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Write by **index** `out[i]` — order preserved regardless of completion order.
- Jobs channel carries indices, not values — cheap for large elements.
- `f` must be **pure** or thread-safe — shared state in `f` needs sync.
- Worker count > len(in) wastes goroutines — cap workers.

## Q&A

**Q: Complexity?**  
A: O(n) tasks; O(workers) concurrency.

**Q: Error handling?**  
A: Return `([]B, error)` with `errgroup` or atomic first-error flag + cancel.

**Q: vs orderless map?**  
A: Collect pairs `(i, result)` then sort — slower than direct index write.

**Q: `in` immutable during map?**  
A: Workers only read — do not mutate `in` concurrently.

**Q: Edge cases?**  
A: `workers==0` — define as 1 or panic; empty `in` → empty `out`.
