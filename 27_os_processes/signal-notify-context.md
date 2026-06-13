# Signal-driven cancellation

## Live interview task
Cancel the application context on `SIGINT` or `SIGTERM`.

## Candidate solution

```go
ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
defer stop()
if err := run(ctx); err != nil { log.Fatal(err) }
```

## Interview notes / pitfalls
- Call `stop` to release signal resources and restore default behavior.
- Windows signal support differs from Unix.
- A second signal may intentionally force immediate termination.
