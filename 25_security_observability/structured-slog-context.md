# Structured slog with request context

## Live interview task
Log method, path, status, duration, and request ID with `log/slog`.

## Interview notes / pitfalls
- Use stable attribute names and typed values, not formatted message strings.
- Never log credentials, authorization headers, or arbitrary request bodies.
- Inject a logger or request attributes through context sparingly.
- Logging every successful high-volume request may need sampling.
