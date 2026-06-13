# Configuration precedence

## Live interview task
Load a port from a default, `APP_PORT`, and `-port`, with flags taking highest precedence.

## Solution outline
- Start with the default.
- Parse and validate the environment value with `strconv.Atoi`.
- Bind the resulting value as the flag default, then call `flag.Parse`.
- Validate the final range `1..65535` once.

## Interview notes / pitfalls
- `os.Getenv` cannot distinguish unset from explicitly empty; use `os.LookupEnv`.
- Report which source contained an invalid value without logging secrets.
- Centralize validation after precedence has been resolved.
