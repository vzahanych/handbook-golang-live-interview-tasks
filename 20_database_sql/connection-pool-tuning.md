# SQL connection pool tuning

## Live interview task
Explain and configure `SetMaxOpenConns`, `SetMaxIdleConns`, `SetConnMaxLifetime`, and `SetConnMaxIdleTime`.

## Interview notes / pitfalls
- The pool is inside `*sql.DB`; `sql.Open` usually does not establish a connection.
- Unlimited open connections can overload the database.
- A transaction occupies one connection until commit or rollback.
- Verify startup with `PingContext`, but do not make every request ping first.
