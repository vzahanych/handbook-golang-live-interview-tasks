# colly multipart form submit

## Live interview task
Submit a multipart form with a file field.

## Concepts covered
- Colly
- multipart POST

## Candidate solution

```go
package main

import (
    "log"
    "os"
    "github.com/gocolly/colly/v2"
)

func main() {
    os.WriteFile("upload.txt", []byte("hello"), 0644)
    c := colly.NewCollector()
    c.OnResponse(func(r *colly.Response) { log.Println("status", r.StatusCode) })
    err := c.PostMultipart("https://httpbin.org/post", map[string][]byte{"file": []byte("hello")})
    if err != nil { log.Println(err) }
}
```

## Run

```bash
go mod init scrape && go get github.com/gocolly/colly/v2 && go run .
```

## Interview notes / pitfalls
- `PostMultipart` map is field name → file bytes — not the same as `Post` form fields.
- Large files should stream from disk — map loads entire file into memory.
- Boundary and Content-Type set by Colly — don't set manually unless custom.

## Q&A

**Q: Upload real file?**  
A: `os.ReadFile` into map value or use lower-level `http.Client` with `multipart.Writer`.

**Q: Complexity?**  
A: O(file size) memory in demo — streaming O(1) buffer with io.Copy.

**Q: Edge cases?**  
A: Multiple files same field name, filename metadata for server validation.

**Q: vs Post?**  
A: Post = urlencoded; PostMultipart = `multipart/form-data` for binaries.

**Q: Production?**  
A: Size limits, virus scan, timeout on slow uploads.
