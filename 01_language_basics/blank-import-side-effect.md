# blank import side effect

## Live interview task
Use a blank import to register a driver or trigger package `init` for side effects only.

## Concepts covered
- import declarations
- blank identifier
- init side effects
- database/sql driver pattern

## Candidate solution

```go
package main

import (
    "database/sql"
    "fmt"

    _ "modernc.org/sqlite" // registers "sqlite" driver via init()
)

func main() {
    db, err := sql.Open("sqlite", ":memory:")
    if err != nil {
        panic(err)
    }
    defer db.Close()
    fmt.Println("driver registered:", db.Driver())
}
```

## Minimal pattern (no external deps)

```go
package main

import _ "fmt" // legal but pointless — illustrates syntax only

// Real use: _ "github.com/lib/pq"
//           _ "github.com/go-sql-driver/mysql"
//           _ "image/png"  // register PNG decoder
```

## Run

```bash
go mod init example
go get modernc.org/sqlite
go run .
```

## Interview notes / pitfalls
- Blank import `_ "pkg"` runs `pkg`'s `init` but does not expose names — used for registration.
- `database/sql` drivers call `sql.Register` in `init`; you must import the driver anonymously.
- Overuse creates hidden dependencies — document why each blank import exists.
- `go mod why` helps trace why a driver package is linked.

## Q&A

**Q: Why not `import "driver"` and call `driver.Register()`?**  
A: Drivers self-register in `init` so `sql.Open("postgres", dsn)` works without coupling main to a concrete driver package.

**Q: Other examples?**  
A: `image/png`, `image/jpeg` for `image.Decode`; pprof `import _ "net/http/pprof"`.

**Q: Can blank imports slow startup?**  
A: Yes — every imported package runs `init`. Keep the import graph lean.

**Q: Testing without real driver?**  
A: Use `sqlmock` or inject a `*sql.DB` from test helpers; avoid blank-importing prod drivers in unit tests when possible.
