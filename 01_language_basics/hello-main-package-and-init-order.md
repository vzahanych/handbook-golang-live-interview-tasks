# hello main package and init order

## Live interview task
Show package-level initialization order: constants, variables, init functions, then main.

## Concepts covered
- package clause
- package initialization
- init
- main

## Candidate solution

```go
package main

import "fmt"

const app = "interview"

var build = trace("var build")

func trace(s string) string {
    fmt.Println(s)
    return s
}

func init() { fmt.Println("init 1") }
func init() { fmt.Println("init 2") }

func main() {
    fmt.Println("main", app, build)
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- init order is per file order after dependency initialization

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
