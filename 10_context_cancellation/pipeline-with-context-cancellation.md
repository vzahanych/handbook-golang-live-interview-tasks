# pipeline with context cancellation

## Live interview task
Cancel a pipeline early so upstream goroutines can exit.

## Concepts covered
- context
- pipeline cancellation
- leak prevention

## Candidate solution

```go
package main

import (
    "context"
    "fmt"
)

func gen(ctx context.Context) <-chan int {
    out := make(chan int)
    go func(){ defer close(out); for i := 0; ; i++ { select { case <-ctx.Done(): return; case out <- i: } } }()
    return out
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    for v := range gen(ctx) {
        fmt.Println(v)
        if v == 3 { cancel(); break }
    }
}
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
