# io reader line counter

## Live interview task
Write a function that depends on io.Reader and is easy to test.

## Concepts covered
- interfaces
- io.Reader
- bufio.Scanner

## Candidate solution

```go
package main

import (
    "bufio"
    "fmt"
    "io"
    "strings"
)

func countLines(r io.Reader) (int, error) {
    sc := bufio.NewScanner(r)
    n := 0
    for sc.Scan() { n++ }
    return n, sc.Err()
}

func main() {
    n, _ := countLines(strings.NewReader("a\nb\nc\n"))
    fmt.Println(n)
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
