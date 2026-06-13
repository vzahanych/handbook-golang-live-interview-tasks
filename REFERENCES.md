# Go interview reference list

Use these references to verify language rules, standard-library behavior, version-specific details, and production guidance behind the tasks.

## Language and tooling

- [The Go Programming Language Specification](https://go.dev/ref/spec)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go memory model](https://go.dev/ref/mem)
- [Go command documentation](https://pkg.go.dev/cmd/go)
- [Go modules reference](https://go.dev/ref/mod)
- [Go release history](https://go.dev/doc/devel/release)

## Standard library APIs

- [Standard library index](https://pkg.go.dev/std)
- [slices](https://pkg.go.dev/slices)
- [maps](https://pkg.go.dev/maps)
- [context](https://pkg.go.dev/context)
- [sync](https://pkg.go.dev/sync)
- [sync/atomic](https://pkg.go.dev/sync/atomic)
- [net/http](https://pkg.go.dev/net/http)
- [net/url](https://pkg.go.dev/net/url)
- [net/netip](https://pkg.go.dev/net/netip)
- [database/sql](https://pkg.go.dev/database/sql)
- [encoding/json](https://pkg.go.dev/encoding/json)
- [reflect](https://pkg.go.dev/reflect)
- [time](https://pkg.go.dev/time)
- [os/exec](https://pkg.go.dev/os/exec)
- [log/slog](https://pkg.go.dev/log/slog)

## Concurrency, testing, and performance

- [Go concurrency patterns: pipelines and cancellation](https://go.dev/blog/pipelines)
- [Data race detector](https://go.dev/doc/articles/race_detector)
- [Fuzzing](https://go.dev/doc/security/fuzz/)
- [Diagnostics](https://go.dev/doc/diagnostics)
- [Profiling Go programs](https://go.dev/blog/pprof)

## HTTP, security, and scraping

- [Go web application security](https://go.dev/doc/security/)
- [Go vulnerability management](https://go.dev/doc/security/vuln/)
- [Colly documentation](https://go-colly.org/docs/)
- [Colly API reference](https://pkg.go.dev/github.com/gocolly/colly/v2)
- [Robots exclusion protocol, RFC 9309](https://www.rfc-editor.org/rfc/rfc9309)

## Interview use

- Confirm the Go version before asking version-sensitive questions.
- Treat candidate solutions as discussion starters, not production-ready templates.
- Evaluate contracts, edge cases, cancellation, ownership, complexity, and tests in addition to syntax.
