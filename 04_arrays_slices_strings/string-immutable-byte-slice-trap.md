# string immutable byte slice trap

## Live interview task
Explain why mutating a byte slice converted from a string does not change the string, and show safe patterns.

## Concepts covered
- strings immutability
- byte slices
- conversion

## Candidate solution

```go
package main

import "fmt"

func main() {
    s := "hello"
    b := []byte(s)
    b[0] = 'H'
    fmt.Println(s)  // hello — unchanged
    fmt.Println(string(b)) // Hello

    // Illegal at compile time:
    // s[0] = 'H'

    // Shared backing if compiler optimizes? Treat as copy-on-convert:
    b2 := []byte(s)
    b2[1] = 'a'
    fmt.Println(s) // still hello
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Strings are **immutable** — bytes cannot be modified in place.
- `[]byte(s)` allocates a **copy** (unless unsafe / noescape optimizations — do not rely on aliasing).
- `string(b)` copies bytes into immutable string — O(n) allocation.
- Mutating `b` after `[]byte(s)` never affects `s` in normal code.

## Q&A

**Q: Zero-copy conversion?**  
A: `unsafe` / `strings.Builder` internals — not for application code; mention only as advanced topic.

**Q: Range over string?**  
A: By runes (UTF-8 decoded); `for i := range s` gives byte index.

**Q: Concatenation cost?**  
A: `s + t` allocates new string — use `Builder` for many joins.

**Q: `[]byte` to `string` without copy?**  
A: Not safely in pure Go without copying — mutating `b` after `string(b)` is undefined if you used unsafe.

**Q: Interview one-liner?**  
A: "String is read-only UTF-8 bytes; `[]byte(s)` is a modifiable copy."
