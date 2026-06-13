# read file lines scanner

## Live interview task
Read text line by line with `bufio.Scanner`.

## Concepts covered
- os.Open
- defer close
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
    if err != nil {
        return nil, err
    }
    defer f.Close()

    sc := bufio.NewScanner(f)
    var out []string
    for sc.Scan() {
        out = append(out, strings.TrimSpace(sc.Text()))
    }
    return out, sc.Err()
}

func main() {
    _ = os.WriteFile("demo.txt", []byte("a\nb\n"), 0644)
    xs, _ := lines("demo.txt")
    fmt.Println(xs)
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Always check `sc.Err()` after loop — EOF is not an error, I/O errors are.
- Default max token 64KB — `sc.Buffer(make([]byte, 0, 64*1024), max)` for long lines.
- `Scanner` splits on lines — use `ScanWords` or custom `SplitFunc` for other tokens.
- For very large files, process per line without storing all in slice.

## Q&A

**Q: vs `ReadString('\n')`?**  
A: Scanner simpler; Reader more control for partial reads.

**Q: `os.ReadFile`?**  
A: Whole file in memory — fine for small files only.

**Q: Empty file?**  
A: Returns empty slice, nil err.

**Q: Windows `\r\n`?**  
A: `TrimSpace` removes `\r`; or `sc.Text()` strips `\n` only.

**Q: Complexity?**  
A: O(bytes) time, O(lines) memory if collecting all.
