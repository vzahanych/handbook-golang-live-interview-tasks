# file upload handler

## Live interview task
Accept a multipart file upload and report its size.

## Concepts covered
- multipart/form-data
- ParseMultipartForm
- io.Copy

## Candidate solution

```go
package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

func upload(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }

    const maxMem = 10 << 20 // 10 MiB
    if err := r.ParseMultipartForm(maxMem); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    f, header, err := r.FormFile("file")
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer f.Close()

    n, err := io.Copy(io.Discard, f)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "%s: %d bytes\n", header.Filename, n)
}

func main() {
    http.HandleFunc("/upload", upload)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Run

```bash
go run .
curl -F 'file=@/etc/hosts' localhost:8080/upload
```

## Interview notes / pitfalls
- `ParseMultipartForm` limits memory — large files spill to temp disk.
- `FormFile("file")` — field name must match HTML/curl form.
- Stream to destination with `io.Copy(dst, f)` — don't read all into memory.
- Validate filename, content-type, virus scan in production.

## Q&A

**Q: Max upload size?**  
A: `r.Body = MaxBytesReader(w, r.Body, max)` before parse.

**Q: Multiple files?**  
A: `r.MultipartForm.File["files"]` slice.

**Q: Progress?**  
A: `io.TeeReader` or custom `Reader` wrapper.

**Q: Security?**  
A: Sanitize paths, store outside web root, random object names.

**Q: Complexity?**  
A: O(bytes) streamed.
