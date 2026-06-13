# Execute a command with context

## Live interview task
Run a subprocess with a timeout and capture stdout and stderr separately.

## Candidate pattern

```go
ctx, cancel := context.WithTimeout(parent, 2*time.Second)
defer cancel()
cmd := exec.CommandContext(ctx, name, args...)
var stdout, stderr bytes.Buffer
cmd.Stdout, cmd.Stderr = &stdout, &stderr
err := cmd.Run()
```

## Interview notes / pitfalls
- Never invoke a shell with untrusted concatenated input.
- Decide how to limit captured output.
- Child processes may outlive a killed parent unless process groups are handled.
