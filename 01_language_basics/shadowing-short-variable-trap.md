# shadowing short variable trap

## Live interview task
Find and fix the classic short variable declaration shadowing bug.

## Concepts covered
- short variable declarations
- scope
- shadowing

## Candidate solution

```go
package main

import (
    "errors"
    "fmt"
)

func load(ok bool) (string, error) {
    var err error
    value := "default"

    if !ok {
        // Use assignment, not :=, otherwise err is shadowed inside the if block.
        err = errors.New("load failed")
    } else {
        value = "loaded"
    }
    return value, err
}

func main() {
    fmt.Println(load(false))
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- := redeclares at least one new variable in the current block

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
