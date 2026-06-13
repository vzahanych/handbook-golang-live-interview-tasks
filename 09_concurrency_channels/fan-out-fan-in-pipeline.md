# fan out fan in pipeline

## Live interview task
Build a pipeline with fan-out square workers and a fan-in merge stage.

## Concepts covered
- pipeline
- fan-out
- fan-in

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            out <- n
        }
    }()
    return out
}

func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            out <- n * n
        }
    }()
    return out
}

func merge(chs ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    wg.Add(len(chs))
    for _, ch := range chs {
        go func(c <-chan int) {
            defer wg.Done()
            for v := range c {
                out <- v
            }
        }(ch)
    }
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}

func main() {
    in := gen(1, 2, 3, 4)
  // Fan-out: two sq stages on same input — duplicate work demo; real pipelines split work
    merged := merge(sq(in), sq(in))
    for v := range merged {
        fmt.Println(v)
    }
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Stages connected by channels — each stage in goroutine, `range` input, `close` output.
- Fan-out: multiple workers read same channel — Go **multiplexes** sends (each value to one worker).
- Fan-in: merge goroutines write one output — need `WaitGroup` before close out.
- Demo merges two `sq(in)` on same `in` — only one consumer gets each value from unbuffered fan-out; fix: `tee` or broadcast for true duplicate pipelines.

## Q&A

**Q: Real fan-out pattern?**  
A: One `jobs` channel, N identical workers ranging `jobs`.

**Q: Backpressure?**  
A: Unbuffered links slow consumer slows whole pipeline.

**Q: Context cancel?**  
A: Stop all stages on `ctx.Done()` — close upstream.

**Q: Complexity?**  
A: O(n) per stage; parallelism up to worker count.

**Q: Mention in interview?**  
A: Sharing one channel among multiple readers splits work — don't expect duplicate processing per value.
