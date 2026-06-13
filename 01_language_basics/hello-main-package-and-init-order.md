# hello main package and init order

## Live interview task
Show package-level initialization order: constants, variables, init functions, then main.

## Concepts covered
- package clause
- package initialization
- init
- main

## Candidate solution

```go
package main

import "fmt"

const app = "interview"

var build = trace("var build")

func trace(s string) string {
    fmt.Println(s)
    return s
}

func init() { fmt.Println("init 1") }
func init() { fmt.Println("init 2") }

func main() {
    fmt.Println("main", app, build)
}
```

## Run

```bash
go run .
```

## Expected output

```
var build
init 1
init 2
main interview interview
```

Order: package-level `const`/`var` initializers (in declaration order) → all `init` functions (in source order) → `main`.

## Interview notes / pitfalls
- `init` runs after imported packages finish initializing (depth-first dependency order).
- Multiple `init` functions in one file run top-to-bottom; across files in a package, order is file-name order (not guaranteed across toolchains for same name — rely on explicit deps instead).
- `init` cannot be called manually; use it only for cheap, idempotent package setup.
- A `var` initializer like `trace("var build")` runs **before** any `init` — easy to get wrong in whiteboard questions.

## Q&A

**Q: What runs first — `init` or package-level `var`?**  
A: Package-level variables and constants initialize first, in dependency order within the package. Then all `init` functions run. Finally `main` runs in `package main`.

**Q: Can you have more than one `init` per package?**  
A: Yes. They all run, in the order they appear in source files (per file, top to bottom).

**Q: What happens if `init` panics?**  
A: The program aborts before `main` runs. Imported packages that already initialized stay initialized; the failing package does not complete init.

**Q: When is `init` appropriate in production?**  
A: Registering drivers (`database/sql`), global metrics, validating package-level config, or one-time caches. Avoid heavy I/O, network calls, or logic that makes testing hard — prefer explicit `New()` constructors.

**Q: Complexity?**  
A: O(1) per initializer; space O(1) for this example. Real packages: watch for init chains that slow cold start.
