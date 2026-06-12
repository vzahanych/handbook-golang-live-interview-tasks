# copy file with buffer

## Live interview task
Copy a file using io.Copy and proper close/error handling.

## Concepts covered
- io.Copy
- files
- defer

## Candidate solution

```go
package main

import (
    "io"
    "os"
)

func copyFile(dst, src string) error {
    in, err := os.Open(src); if err != nil { return err }
    defer in.Close()
    out, err := os.Create(dst); if err != nil { return err }
    defer func(){ _ = out.Close() }()
    if _, err := io.Copy(out, in); err != nil { return err }
    return out.Sync()
}

func main() { os.WriteFile("a.txt", []byte("go"), 0644); _ = copyFile("b.txt", "a.txt") }
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
