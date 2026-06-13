# Colly request timeout and context cancellation

## Live interview task
Connect a caller context and deadline to a Colly scrape.

## Solution outline
- Configure `collector.SetRequestTimeout` for each HTTP request.
- Register `OnRequest` and replace the underlying request context through a custom transport or use a goroutine that aborts the collector when the caller is done.
- Return a result only after `collector.Wait()` for asynchronous collectors.
- Normalize Colly errors and `ctx.Err()` without hiding the cancellation cause.

## Interview notes / pitfalls
- Colly's request timeout and `context.Context` cancellation are related but not identical.
- Never leave cancellation-monitor goroutines blocked after the scrape completes.
- A single collector shared across unrelated calls makes cancellation boundaries difficult.
