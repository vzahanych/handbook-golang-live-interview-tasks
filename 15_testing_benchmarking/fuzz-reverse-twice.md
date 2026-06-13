# fuzz reverse twice

## Live interview task
Fuzz test that reversing a UTF-8 string twice returns the original.

## Concepts covered
- fuzzing (Go 1.18+)
- rune-safe reverse
- unicode edge cases

## Candidate solution

```go
package reverse

import "testing"

func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

func FuzzReverseTwice(f *testing.F) {
    f.Add("hello")
    f.Add("Go语言")
    f.Add("")
    f.Fuzz(func(t *testing.T, s string) {
        if got := Reverse(Reverse(s)); got != s {
            t.Fatalf("got %q want %q", got, s)
        }
    })
}
```

## Run

```bash
go test -fuzz=FuzzReverseTwice -fuzztime=10s
```

## Interview notes / pitfalls
- Seed corpus with `f.Add` — known cases run before random mutations.
- Fuzz finds crash inputs — saved under `testdata/fuzz/FuzzReverseTwice`.
- Byte-reverse would fail fuzz on UTF-8 — rune reverse passes this property.
- Combining marks / emoji ZWJ may fail "visual" reverse — property test is codepoint-level.

## Q&A

**Q: Stop fuzzing?**  
A: Ctrl+C or `-fuzztime=30s`.

**Q: Regression?**  
A: Commit `testdata/fuzz` corpus — `go test` replays failures.

**Q: vs property-based (rapid)?**  
A: Built-in fuzz no extra deps.

**Q: Fuzz with invalid UTF-8?**  
A: `range` string yields replacement chars — document behavior.

**Q: Complexity per input?**  
A: O(runes) per reverse.
