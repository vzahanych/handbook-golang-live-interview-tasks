# Embed static assets

## Live interview task
Embed a `static` directory and serve it below `/assets/`.

## Candidate pattern

```go
//go:embed static
var assets embed.FS

sub, err := fs.Sub(assets, "static")
if err != nil { return err }
mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(sub))))
```

## Interview notes / pitfalls
- `go:embed` patterns are evaluated at build time.
- Set cache policy and content-security headers deliberately.
- Embedded large assets increase binary size and deployment memory pressure.
