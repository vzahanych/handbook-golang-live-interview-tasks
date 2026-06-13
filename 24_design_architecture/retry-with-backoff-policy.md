# Retry with backoff policy

## Live interview task
Retry only transient failures with exponential backoff, jitter, context cancellation, and a maximum attempt count.

## Interview notes / pitfalls
- Retry classification is more important than the loop itself.
- Respect server hints such as `Retry-After` where applicable.
- Do not retry non-idempotent operations without an idempotency mechanism.
- Cap both individual delays and total elapsed time.
