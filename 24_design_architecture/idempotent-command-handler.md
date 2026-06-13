# Idempotent command handler

## Live interview task
Process a payment command exactly once from the caller's perspective, despite retries.

## Required behavior
- Require an idempotency key scoped to the operation and caller.
- Persist key, request fingerprint, status, and response atomically with side effects.
- Return the stored response for an identical retry.
- Reject reuse of the same key with different input.

## Interview notes / pitfalls
- Networks cannot generally guarantee literal exactly-once execution.
- In-memory deduplication fails across restarts and multiple instances.
