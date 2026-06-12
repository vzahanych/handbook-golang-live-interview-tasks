# rune safe reverse string

## Live interview task
Reverse a UTF-8 string by runes, not bytes.

## Concepts covered
- strings
- runes
- UTF-8

## Candidate solution

```go
package main

import "fmt"

func reverseString(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

func main() { fmt.Println(reverseString("Go语言")) }
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
