# sync pool buffer reuse

## Live interview task
Use sync.Pool to reuse temporary buffers.

## Concepts covered
- sync.Pool
- allocations
- temporary objects

## Candidate solution

```go
package main

import (
    "bytes"
    "fmt"
    "sync"
)

var pool = sync.Pool{New: func() any { return new(bytes.Buffer) }}

func render(name string) string {
    b := pool.Get().(*bytes.Buffer)
    b.Reset()
    defer pool.Put(b)
    b.WriteString("hello ")
    b.WriteString(name)
    return b.String()
}

func main() { fmt.Println(render("go")) }
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
