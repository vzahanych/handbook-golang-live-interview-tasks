# Colly concurrent multiple-URL scraper

## Live interview task
Visit a slice of URLs asynchronously, extract titles, cap parallelism, and preserve input order.

## Candidate solution

```go
type Result struct { URL, Title, Error string }

func Scrape(urls []string, parallelism int) ([]Result, error) {
	if parallelism <= 0 {
		return nil, errors.New("parallelism must be positive")
	}
	results := make([]Result, len(urls))
	c := colly.NewCollector(colly.Async(true), colly.AllowURLRevisit())
	if err := c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: parallelism}); err != nil {
		return nil, err
	}
	c.OnRequest(func(r *colly.Request) {
		i, _ := strconv.Atoi(r.Ctx.Get("index"))
		results[i].URL = r.URL.String()
	})
	c.OnHTML("title", func(e *colly.HTMLElement) {
		i, _ := strconv.Atoi(e.Request.Ctx.Get("index"))
		results[i] = Result{URL: e.Request.URL.String(), Title: e.Text}
	})
	c.OnError(func(r *colly.Response, err error) {
		i, _ := strconv.Atoi(r.Ctx.Get("index"))
		results[i] = Result{URL: r.Request.URL.String(), Error: err.Error()}
	})
	for i, target := range urls {
		ctx := colly.NewContext(); ctx.Put("index", strconv.Itoa(i))
		if err := c.Request("GET", target, nil, ctx, nil); err != nil {
			results[i] = Result{URL: target, Error: err.Error()}
		}
	}
	c.Wait()
	return results, nil
}
```

## Interview notes / pitfalls
- Pass the index directly through Colly context; URL lookup would fail for duplicate targets.
- `AllowURLRevisit` is required when duplicate inputs must produce duplicate result rows.
- Call `Wait` when `Async(true)` is enabled.
- Protect shared state unless each callback owns a distinct index.
- Validate `parallelism` before constructing the limit rule.
