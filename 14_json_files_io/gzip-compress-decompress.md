# gzip compress decompress

## Live interview task
Compress and decompress bytes with `compress/gzip`.

## Concepts covered
- gzip
- io pipelines
- Close writers

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
    if _, err := zw.Write([]byte("hello hello hello")); err != nil {
        panic(err)
    }
    if err := zw.Close(); err != nil { // flush footer — required
        panic(err)
    }

    zr, err := gzip.NewReader(bytes.NewReader(buf.Bytes()))
    if err != nil {
        panic(err)
    }
    defer zr.Close()

    plain, err := io.ReadAll(zr)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(plain))
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- **Must** `Close()` gzip writer — writes CRC and footer; `Flush` not enough for finish.
- `gzip.NewReader` on non-gzip data errors at header.
- Stream: `io.Copy(zw, src)` then `zw.Close()`.
- `gzip.BestSpeed` vs `DefaultCompression` — `zw.Level` via `gzip.NewWriterLevel`.

## Q&A

**Q: File compress?**  
A: `os.Create` + `gzip.NewWriter(file)`.

**Q: HTTP?**  
A: `w.Header().Set("Content-Encoding", "gzip")` + gzip writer wrapping ResponseWriter.

**Q: Size increase?**  
A: Small payloads may grow — gzip needs redundancy.

**Q: vs zlib?**  
A: gzip = zlib wrapper + header/footer — `compress/zlib` for raw.

**Q: Complexity?**  
A: O(n) on data size.
