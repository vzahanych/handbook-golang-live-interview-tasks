# object pool with reset

## Live interview task
Reuse objects with an explicit pool and reset before putting them back.

## Concepts covered
- manual pooling
- reset
- allocations

## Candidate solution

```go
package main

import "fmt"

type Buffer struct{ data []byte }
var pool []*Buffer

func get() *Buffer {
    if len(pool) == 0 { return &Buffer{data: make([]byte, 0, 1024)} }
    b := pool[len(pool)-1]; pool = pool[:len(pool)-1]; return b
}
func put(b *Buffer) { b.data = b.data[:0]; pool = append(pool, b) }

func main() { b := get(); b.data = append(b.data, "go"...); fmt.Println(string(b.data)); put(b) }
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
