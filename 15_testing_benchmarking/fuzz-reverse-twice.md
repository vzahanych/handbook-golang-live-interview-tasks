# fuzz reverse twice

## Live interview task
Fuzz test that reversing a string twice returns the original string.

## Concepts covered
- fuzzing
- unicode strings

## Candidate solution

```go
package reverse

import "testing"

func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 { r[i], r[j] = r[j], r[i] }
    return string(r)
}

func FuzzReverseTwice(f *testing.F) {
    f.Add("hello")
    f.Add("Go语言")
    f.Fuzz(func(t *testing.T, s string) {
        if got := Reverse(Reverse(s)); got != s { t.Fatalf("got %q want %q", got, s) }
    })
}
```

## Run

```bash
go test -fuzz=Fuzz
```

## Interview notes / pitfalls
- None specific; discuss edge cases and complexity.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
