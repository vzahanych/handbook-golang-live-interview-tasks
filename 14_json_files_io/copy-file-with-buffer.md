# copy file with buffer

## Live interview task
Copy a file using `io.Copy` with proper close and error handling.

## Concepts covered
- io.Copy
- os.File
- defer
- Sync

## Candidate solution

```go
package main

import (
    "io"
    "os"
)

func copyFile(dst, src string) error {
    in, err := os.Open(src)
    if err != nil {
        return err
    }
    defer in.Close()

    out, err := os.Create(dst)
    if err != nil {
        return err
    }

    if _, err := io.Copy(out, in); err != nil {
        _ = out.Close()
        return err
    }
    if err := out.Sync(); err != nil {
        _ = out.Close()
        return err
    }
    return out.Close()
}

func main() {
    _ = os.WriteFile("a.txt", []byte("go"), 0644)
    _ = copyFile("b.txt", "a.txt")
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- On `Copy` error, close `out` and **remove partial dst** — `os.Remove(dst)` for atomic copy pattern.
- `defer out.Close()` after successful Copy can hide Close error — explicit Close on success path as shown.
- `io.Copy` uses 32KB internal buffer — efficient.
- `os.Rename` same filesystem is atomic move — copy+rename for safe replace.

## Q&A

**Q: Copy permissions?**  
A: `os.Stat(src)` then `chmod` on dst or use `io.Copy` + `Chmod`.

**Q: `cp` syscall?**  
A: Go portable — `Copy` works everywhere.

**Q: Large files?**  
A: Streaming — constant memory.

**Q: `io.CopyBuffer`?**  
A: Custom buffer from `sync.Pool`.

**Q: Complexity?**  
A: O(bytes) time, O(1) memory.
