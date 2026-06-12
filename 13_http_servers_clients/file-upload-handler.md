# file upload handler

## Live interview task
Accept a multipart file upload and report its size.

## Concepts covered
- multipart forms
- io.Copy

## Candidate solution

```go
package main

import (
    "fmt"
    "io"
    "net/http"
)

func upload(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseMultipartForm(10 << 20); err != nil { http.Error(w, err.Error(), 400); return }
    f, header, err := r.FormFile("file")
    if err != nil { http.Error(w, err.Error(), 400); return }
    defer f.Close()
    n, _ := io.Copy(io.Discard, f)
    fmt.Fprintf(w, "%s: %d bytes\n", header.Filename, n)
}

func main() { http.HandleFunc("/upload", upload); http.ListenAndServe(":8080", nil) }
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
