# URL resolution and normalization

## Live interview task
Resolve relative links against a base URL and keep only same-origin HTTP links.

## Interview notes / pitfalls
- Use `url.Parse` and `ResolveReference`, not string concatenation.
- Normalize default ports and host case for origin comparison.
- Fragments do not change the fetched resource and can be removed for deduplication.
- Path normalization can change semantics on unusual servers; define the contract first.
