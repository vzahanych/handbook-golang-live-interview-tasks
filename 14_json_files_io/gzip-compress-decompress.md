# gzip compress decompress

## Live interview task
Compress and decompress bytes with gzip.

## Concepts covered
- gzip
- bytes.Buffer
- io

## Candidate solution

```go
package main

import (
    "bytes"
    "compress/gzip"
    "fmt"
    "io"
)

func main() {
    var buf bytes.Buffer
    zw := gzip.NewWriter(&buf)
    zw.Write([]byte("hello hello hello"))
    zw.Close()
    zr, _ := gzip.NewReader(&buf)
    plain, _ := io.ReadAll(zr)
    zr.Close()
    fmt.Println(string(plain))
}
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
