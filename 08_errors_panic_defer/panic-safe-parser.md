# panic safe parser

## Live interview task
Convert a panic-prone helper into an error-returning parser.

## Concepts covered
- recover
- error conversion

## Candidate solution

```go
package main

import (
    "fmt"
    "strconv"
)

func mustAtoi(s string) int {
    n, err := strconv.Atoi(s)
    if err != nil { panic(err) }
    return n
}

func parse(s string) (n int, err error) {
    defer func() { if r := recover(); r != nil { err = fmt.Errorf("parse %q: %v", s, r) } }()
    return mustAtoi(s), nil
}

func main() { fmt.Println(parse("x")) }
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
