# nil interface vs interface holding nil pointer

## Live interview task
Explain why an interface value holding a nil pointer is not nil.

## Concepts covered
- interface representation
- nil
- dynamic type

## Candidate solution

```go
package main

import "fmt"

type Reader struct{}
func (*Reader) Read() {}

func main() {
    var p *Reader = nil
    var x any = p
    fmt.Println(p == nil) // true
    fmt.Println(x == nil) // false: interface has dynamic type *Reader
    fmt.Printf("%T %#v\n", x, x)
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
