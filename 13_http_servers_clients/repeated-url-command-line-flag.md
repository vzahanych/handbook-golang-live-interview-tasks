# Repeated URL command-line flag

## Live interview task
Implement a `-url` flag that may be repeated and reject startup when no URL is supplied.

## Candidate solution

```go
type urlList []string

func (u *urlList) String() string { return strings.Join(*u, ",") }

func (u *urlList) Set(raw string) error {
	parsed, err := url.ParseRequestURI(raw)
	if err != nil || (parsed.Scheme != "http" && parsed.Scheme != "https") {
		return fmt.Errorf("invalid HTTP URL %q", raw)
	}
	*u = append(*u, parsed.String())
	return nil
}

func main() {
	var urls urlList
	flag.Var(&urls, "url", "target URL; repeat for multiple targets")
	flag.Parse()
	if len(urls) == 0 {
		log.Fatal("at least one -url is required")
	}
}
```

Imports: `flag`, `fmt`, `log`, `net/url`, `strings`.

## Interview notes / pitfalls
- `flag.String` only keeps one value; implement `flag.Value` for repetition.
- Parsing a URI is not enough for SSRF protection.
- Configuration validation should happen before the server starts.
