# Transactional funds transfer

## Live interview task
Move money between accounts atomically and reject insufficient funds.

## Required behavior
- Begin with `BeginTx(ctx, ...)` and defer rollback.
- Debit with a conditional update such as `balance >= amount`.
- Verify exactly one affected row, then credit the receiver.
- Commit only after every check succeeds.

## Interview notes / pitfalls
- A read-then-write sequence can race without locks or a conditional update.
- `defer tx.Rollback()` is safe after a successful commit.
- Discuss isolation levels, deadlocks, retries, and decimal money representation.
