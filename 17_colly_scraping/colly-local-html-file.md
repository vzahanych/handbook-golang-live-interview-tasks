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
- `file://` URLs need absolute path — `filepath.Abs` avoids cwd surprises.
- Default collector may block file scheme — may need `colly.AllowURLRevisit` or transport tweaks per version.
- Ideal for unit tests — no network flake in CI.

## Q&A

**Q: Why file URLs?**  
A: Deterministic tests for selectors without mocking HTTP server.

**Q: Complexity?**  
A: O(file size) parse — same as HTTP response body.

**Q: Edge cases?**  
A: Windows paths (`file:///C:/...`), relative links in local HTML won't resolve to real hosts.

**Q: vs httptest?**  
A: httptest exercises full HTTP stack; file:// tests parsing only.

**Q: Production?**  
A: Fixture HTML in `testdata/` for golden selector tests.
