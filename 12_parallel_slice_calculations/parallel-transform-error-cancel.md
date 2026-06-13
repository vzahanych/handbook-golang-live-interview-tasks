# Parallel slice transform with error cancellation

## Live interview task
Apply `func(T) (R, error)` concurrently, preserve input order, stop scheduling after the first error, and wait for all started workers.

## Required contract
- Validate `workers > 0`.
- Preallocate `results` to the input length.
- Send indexed jobs so workers write distinct result slots.
- Derive a cancelable context and use `sync.Once` for the first error.
- Make both producers and workers select on `ctx.Done()`.

## Interview notes / pitfalls
- Returning before workers exit leaks goroutines and can race with the caller.
- Cancellation does not interrupt an arbitrary transform; the transform must accept context for that.
- Distinct slice indexes can be written concurrently, but appending to one shared slice cannot.
- Decide whether partial results are returned with the error.

## Follow-up questions
- Replace the custom coordination with `errgroup.Group`.
- Limit memory by streaming ordered results.
- Collect all errors with `errors.Join` instead of canceling early.
