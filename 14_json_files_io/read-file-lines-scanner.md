# read file lines scanner

## Live interview task
Read text line by line with bufio.Scanner.

## Concepts covered
- os.Open
- defer
- bufio.Scanner

## Candidate solution

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func lines(path string) ([]string, error) {
    f, err := os.Open(path)
    if err != nil { return nil, err }
    defer f.Close()
    sc := bufio.NewScanner(f)
    var out []string
    for sc.Scan() { out = append(out, strings.TrimSpace(sc.Text())) }
    return out, sc.Err()
}

func main() { _ = os.WriteFile("demo.txt", []byte("a\nb\n"), 0644); xs, _ := lines("demo.txt"); fmt.Println(xs) }
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
