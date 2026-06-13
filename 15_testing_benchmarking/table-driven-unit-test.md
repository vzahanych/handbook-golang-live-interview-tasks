# table driven unit test

## Live interview task
Write a table-driven unit test for a pure function.

## Concepts covered
- testing
- table-driven tests
- t.Run subtests

## Candidate solution

```go
// file: calc/add.go
package calc

func Add(a, b int) int { return a + b }

// file: calc/add_test.go
package calc

import "testing"

func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive", 1, 2, 3},
        {"negative", -1, -2, -3},
        {"zero", 0, 0, 0},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Add(tt.a, tt.b); got != tt.want {
                t.Fatalf("got %d want %d", got, tt.want)
            }
        })
    }
}
```

## Run

```bash
go test ./calc/...
go test -run TestAdd/negative ./calc/...
```

## Interview notes / pitfalls
- Table-driven is idiomatic Go — one test function, many cases.
- `t.Run` gives named subtests — filter with `-run`, parallel with `t.Parallel()` per case.
- Use `t.Helper()` in test helpers for correct line numbers.
- Prefer `cmp.Diff` or `testify` only if team uses them — stdlib `t.Fatalf` fine in interviews.

## Q&A

**Q: `Fatal` vs `Error`?**  
A: `Fatal` stops test immediately; `Error` continues — use Fatal when rest invalid.

**Q: Parallel subtests?**  
A: `t.Parallel()` inside `t.Run` — parent must not use shared mutable state.

**Q: Error type in table?**  
A: `wantErr bool` or `wantErr error` with `errors.Is`.

**Q: Setup/teardown?**  
A: `t.Cleanup(func(){...})` per subtest or shared helper.

**Q: Complexity?**  
A: O(cases * work per case).
