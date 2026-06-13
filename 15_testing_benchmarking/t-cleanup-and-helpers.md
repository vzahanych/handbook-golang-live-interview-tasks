# t cleanup and helpers

## Live interview task
Use `t.Cleanup` for teardown and `t.Helper()` in test helpers.

## Concepts covered
- t.Cleanup
- t.Helper
- test isolation

## Candidate solution

```go
package store

import (
    "os"
    "testing"
)

func tempFile(t *testing.T) string {
    t.Helper()
    f, err := os.CreateTemp("", "test-*")
    if err != nil {
        t.Fatal(err)
    }
    t.Cleanup(func() {
        _ = os.Remove(f.Name())
        _ = f.Close()
    })
    return f.Name()
}

func TestWrite(t *testing.T) {
    path := tempFile(t)
    if err := os.WriteFile(path, []byte("data"), 0644); err != nil {
        t.Fatal(err)
    }
}
```

## Run

```bash
go test
```

## Interview notes / pitfalls
- `t.Cleanup` runs LIFO after test (or subtest) ends — even on `Fatal`.
- Prefer `Cleanup` over `defer` in helpers — runs even when helper calls `t.Fatal` before return.
- `t.Helper()` marks function — failures report caller line, not helper internals.
- Multiple cleanups stack — reverse order like defer.

## Q&A

**Q: vs `defer` in test?**  
A: `defer` in test body fine; Cleanup better inside shared helpers.

**Q: Subtest cleanup?**  
A: Registered on `t` of subtest — runs when subtest completes.

**Q: Panic in test?**  
A: Cleanup still runs.

**Q: Parallel tests?**  
A: Each `t` has own cleanup — don't share temp files across parallel tests without sync.

**Q: Interview mention?**  
A: Shows idiomatic Go 1.14+ test hygiene.
