# SQL null values

## Live interview task
Read an optional nickname and JSON-encode it as either a string or `null`.

## Solution outline
- Scan into `sql.NullString` or `*string`, depending on driver behavior and model conventions.
- Convert database transport types into a domain type at the repository boundary.
- Test null, empty string, and non-empty string independently.

## Interview notes / pitfalls
- SQL `NULL` and `""` are different values.
- Leaking `sql.NullString` through every application layer couples the domain to SQL.
