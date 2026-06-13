# HTTP server production timeouts

## Live interview task
Replace `http.ListenAndServe` with an explicitly configured server and explain every timeout.

## Candidate solution

```go
srv := &http.Server{
	Addr:              *addr,
	Handler:           mux,
	ReadHeaderTimeout: 5 * time.Second,
	ReadTimeout:       15 * time.Second,
	WriteTimeout:      30 * time.Second,
	IdleTimeout:       60 * time.Second,
	MaxHeaderBytes:    1 << 20,
}
if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
	log.Fatal(err)
}
```

## Interview notes / pitfalls
- `ReadHeaderTimeout` limits slow headers; `ReadTimeout` also includes body reads.
- `WriteTimeout` can conflict with long streaming responses.
- Client, reverse-proxy, handler-context, and server timeouts solve different problems.
- Handle `http.ErrServerClosed` separately during graceful shutdown.
