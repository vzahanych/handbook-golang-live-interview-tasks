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
    go func(){ defer close(out); for _, n := range nums { out <- n } }()
    return out
}

func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func(){ defer close(out); for n := range in { out <- n*n } }()
    return out
}

func merge(chs ...<-chan int) <-chan int {
    out := make(chan int); var wg sync.WaitGroup; wg.Add(len(chs))
    for _, ch := range chs { go func(c <-chan int){ defer wg.Done(); for v := range c { out <- v } }(ch) }
    go func(){ wg.Wait(); close(out) }()
    return out
}

func main() { in := gen(1,2,3,4); for v := range merge(sq(in), sq(in)) { fmt.Println(v) } }
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
