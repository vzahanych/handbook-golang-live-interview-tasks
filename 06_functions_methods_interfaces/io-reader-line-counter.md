# io reader line counter

## Live interview task
Write a function that depends on `io.Reader` and is easy to test.

## Concepts covered
- interfaces
- io.Reader
- bufio.Scanner
- testability

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
    for sc.Scan() {
        n++
    }
    return n, sc.Err()
}

func main() {
    n, err := countLines(strings.NewReader("a\nb\nc\n"))
    fmt.Println(n, err) // 3 <nil>
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `io.Reader` is the standard abstraction — files, networks, strings all work.
- `Scanner` default max token 64KB — `sc.Buffer()` for longer lines.
- `Scanner` splits on lines; for arbitrary delimiters use `ReadString` or `bufio.Reader`.
- Always return `sc.Err()` — last scan may fail after loop ends.

## Q&A

**Q: Test without file?**  
A: `strings.NewReader`, `bytes.NewReader`, `strings.NewReader` in table tests.

**Q: Count bytes not lines?**  
A: `io.Copy` to `io.Discard` or `bufio.Reader.Read` loop.

**Q: Complexity?**  
A: O(n) bytes read, O(1) extra space (buffer).

**Q: Empty input?**  
A: Returns 0 lines — no `\n` means one line? `strings.NewReader("")` → 0 scans → 0 lines. Document semantics.

**Q: File last line without newline?**  
A: `Scanner` still counts final line as one token.
