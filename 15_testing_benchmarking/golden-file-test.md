# golden file test

## Live interview task
Use a golden file in `testdata/` to test formatted output.

## Concepts covered
- testdata
- golden files
- snapshot testing

## Candidate solution

```go
package report

import (
    "os"
    "testing"
)

func Render() string {
    return "name,count\ngo,3\n"
}

func TestRenderGolden(t *testing.T) {
    want, err := os.ReadFile("testdata/report.golden")
    if err != nil {
        t.Fatal(err)
    }
    got := Render()
    if got != string(want) {
        t.Fatalf("got %q want %q", got, string(want))
    }
}
```

## Setup

```bash
mkdir -p testdata
printf 'name,count\ngo,3\n' > testdata/report.golden
go test
```

## Interview notes / pitfalls
- `testdata/` ignored by `go build` — convention for fixtures.
- Update golden intentionally — `-update` flag pattern in custom test helper.
- Normalize line endings (`\n` vs `\r\n`) in cross-platform CI.
- Large goldens — consider smaller unit tests; goldens for reports/CLI output.

## Q&A

**Q: vs inline string in test?**  
A: Golden better for multi-line output; inline for small.

**Q: `cmp.Diff`?**  
A: Better failure messages — `google/go-cmp`.

**Q: Regenerate workflow?**  
A: `if os.Getenv("UPDATE_GOLDEN") != "" { os.WriteFile(...) }`.

**Q: Binary output?**  
A: `bytes.Equal` with golden bytes file.

**Q: Complexity?**  
A: O(output size) compare.
