# is palindrome ignore nonletters

## Live interview task
Check whether a string is a palindrome while ignoring spaces, punctuation and case.

## Concepts covered
- runes
- unicode
- two pointers

## Candidate solution

```go
package main

import (
    "fmt"
    "unicode"
)

func palindrome(s string) bool {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < j; {
        if !unicode.IsLetter(r[i]) && !unicode.IsDigit(r[i]) { i++; continue }
        if !unicode.IsLetter(r[j]) && !unicode.IsDigit(r[j]) { j--; continue }
        if unicode.ToLower(r[i]) != unicode.ToLower(r[j]) { return false }
        i++; j--
    }
    return true
}

func main() { fmt.Println(palindrome("A man, a plan, a canal: Panama")) }
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
