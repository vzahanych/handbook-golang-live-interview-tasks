# type alias vs new defined type

## Live interview task
Explain the difference between a type alias and a new defined type.

## Concepts covered
- type definitions
- alias declarations
- type identity

## Candidate solution

```go
package main

import "fmt"

type UserID int       // new defined type
type InternalID = int // alias, identical to int

func printInt(n int) { fmt.Println(n) }

func main() {
    var u UserID = 10
    var i InternalID = 20
    printInt(i)      // ok: alias is int
    printInt(int(u)) // explicit conversion: UserID is distinct from int
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
