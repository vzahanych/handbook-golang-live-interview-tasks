# QueryContext and rows lifecycle

## Live interview task
Query active users with a context and scan them into a slice.

## Candidate solution

```go
rows, err := db.QueryContext(ctx, `SELECT id, name FROM users WHERE active = ?`, true)
if err != nil { return nil, err }
defer rows.Close()
var users []User
for rows.Next() {
	var user User
	if err := rows.Scan(&user.ID, &user.Name); err != nil { return nil, err }
	users = append(users, user)
}
return users, rows.Err()
```

## Interview notes / pitfalls
- Always check `rows.Err()` after the loop.
- Placeholders differ across drivers.
- Parameter binding prevents injection; it does not work for identifiers.
