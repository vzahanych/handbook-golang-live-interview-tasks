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
- None specific; discuss edge cases and complexity.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
