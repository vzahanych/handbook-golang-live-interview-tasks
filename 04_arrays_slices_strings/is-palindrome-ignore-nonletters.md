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
        if !unicode.IsLetter(r[i]) && !unicode.IsDigit(r[i]) {
            i++
            continue
        }
        if !unicode.IsLetter(r[j]) && !unicode.IsDigit(r[j]) {
            j--
            continue
        }
        if unicode.ToLower(r[i]) != unicode.ToLower(r[j]) {
            return false
        }
        i++
        j--
    }
    return true
}

func main() {
    fmt.Println(palindrome("A man, a plan, a canal: Panama")) // true
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Two pointers from ends — skip non-alphanumeric with `unicode` package.
- `ToLower` per rune works for many scripts; full case-folding needs `strings.EqualFold` on substrings or `unicode.SimpleFold`.
- O(n) time; `[]rune` allocation — can optimize with byte indices + `utf8.DecodeRuneInString` to avoid full decode.
- LeetCode 125 classic — state skip logic clearly for interviewer.

## Q&A

**Q: Without `[]rune`?**  
A: Byte indices + decode rune at `i` and `j` — O(n) bytes, no full copy.

**Q: Only ASCII?**  
A: `isAlnum(b)` with `(b>='a'&&b<='z')||...` — faster, state assumptions.

**Q: Edge cases?**  
A: Empty, only punctuation (`true`), single char, Unicode letters.

**Q: Complexity?**  
A: O(n) time; O(n) extra if `[]rune`, O(1) with two byte pointers.
