# Repository transaction boundary

## Live interview task
Design a use case that updates two repositories in one database transaction.

## Interview notes / pitfalls
- Separate repositories that each open their own transaction cannot provide atomicity.
- Pass a transaction-scoped interface or expose a unit-of-work callback.
- Keep SQL transaction mechanics out of pure domain logic where practical.
- Make rollback and panic behavior explicit.
