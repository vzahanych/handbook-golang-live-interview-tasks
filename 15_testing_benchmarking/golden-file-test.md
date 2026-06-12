# golden file test

## Live interview task
Use a golden file to test formatted output.

## Concepts covered
- testdata
- golden files

## Candidate solution

```go
package report

import (
    "os"
    "testing"
)

func Render() string { return "name,count\ngo,3\n" }

func TestRenderGolden(t *testing.T) {
    want, err := os.ReadFile("testdata/report.golden")
    if err != nil { t.Fatal(err) }
    if got := Render(); got != string(want) { t.Fatalf("got %q want %q", got, want) }
}
```

## Run

```bash
mkdir -p testdata && printf "name,count
go,3
" > testdata/report.golden && go test
```

## Interview notes / pitfalls
- None specific; discuss edge cases and complexity.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
