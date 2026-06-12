# iota bitmask permissions

## Live interview task
Implement bitmask permissions with iota and test add/remove/has operations.

## Concepts covered
- constant declarations
- iota
- bitwise operators

## Candidate solution

```go
package main

import "fmt"

type Perm uint8

const (
    Read Perm = 1 << iota
    Write
    Execute
)

func Has(mask, p Perm) bool { return mask&p != 0 }
func Add(mask, p Perm) Perm { return mask | p }
func Remove(mask, p Perm) Perm { return mask &^ p }

func main() {
    var p Perm
    p = Add(p, Read|Write)
    fmt.Println(Has(p, Read), Has(p, Execute))
    p = Remove(p, Write)
    fmt.Println(Has(p, Write))
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
