# defer in loop performance trap

## Live interview task
Show why `defer` inside a tight loop is problematic and refactor to a scoped helper.

## Concepts covered
- defer
- loops
- resource management

## Problematic version

```go
func processAll(paths []string) error {
    for _, p := range paths {
        f, err := os.Open(p)
        if err != nil {
            return err
        }
        defer f.Close() // f is reassigned each iteration; defer evaluates f now and queues Close for that handle — but every Close runs only when processAll returns, so all files stay open until then
        // ... use f
    }
    return nil
}
```

## Candidate solution

```go
package main

import (
    "fmt"
    "os"
)

func processOne(path string) error {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer func() {
        fmt.Println("close", path)
        _ = f.Close()
    }()
    fmt.Println("work", path)
    // ... use f
    return nil
}

func processAll(paths []string) error {
    for _, p := range paths {
        if err := processOne(p); err != nil {
            return err
        }
    }
    return nil
}

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintln(os.Stderr, "usage:", os.Args[0], "file...")
        os.Exit(1)
    }
    if err := processAll(os.Args[1:]); err != nil {
        panic(err)
    }
}
```

## Run

```bash
printf '%s' a > a; printf '%s' b > b; printf '%s' c > c
go run . a b c
```

## Expected output

```
work a
close a
work b
close b
work c
close c
```

## Interview notes / pitfalls
- Defers in a loop stack up until the **function** returns — each iteration queues another Close (for the handle opened that round), but nothing closes until exit, so FDs pile up until then.
- Fix: extract body to `func() { defer close(); ... }()` or `processOne` helper.
- Alternative: explicit `Close()` before next iteration (error handling is verbose).
- Same issue with `defer mu.Unlock()` in a loop holding locks too long.

## Q&A

**Q: How many defers if loop runs 10k times?**  
A: 10k deferred calls queued — all run at function exit.

**Q: Does Go optimize defer in loops?**  
A: No — still stacks defer records.

**Q: When is defer-in-loop acceptable?**  
A: Small bounded iterations (e.g. 3 files) where holding all open briefly is intentional — still document it.

**Q: Complexity?**  
A: Bad version: O(n) resources held; fixed version: O(1) resources at a time.
