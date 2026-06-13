# Optimistic locking

## Live interview task
Update a document only when its version still matches the version read by the caller.

## Candidate query

```sql
UPDATE documents
SET body = ?, version = version + 1
WHERE id = ? AND version = ?
```

## Interview notes / pitfalls
- Treat zero affected rows as a conflict or missing record, then distinguish if required.
- Return the new version to the caller.
- Blind retries can overwrite user intent; retry only when business semantics permit it.
