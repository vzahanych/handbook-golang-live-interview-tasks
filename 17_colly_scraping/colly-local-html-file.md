# colly local html file

## Live interview task
Scrape a local HTML file by allowing file URLs.

## Concepts covered
- Colly
- local files

## Candidate solution

```go
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "github.com/gocolly/colly/v2"
)

func main() {
    os.WriteFile("page.html", []byte(`<html><title>Local</title><a href="/x">x</a></html>`), 0644)
    abs, _ := filepath.Abs("page.html")
    c := colly.NewCollector()
    c.OnHTML("title", func(e *colly.HTMLElement) { fmt.Println(e.Text) })
    c.Visit("file://" + abs)
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
