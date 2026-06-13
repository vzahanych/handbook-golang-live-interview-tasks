# Explicit state machine

## Live interview task
Model an order with `pending -> paid -> shipped` and `pending -> cancelled` transitions.

## Required behavior
- Keep state values as a defined type with validated transitions.
- Reject illegal and duplicate transitions consistently.
- Record transition time and actor.
- Make persistence concurrency-safe with a version check.

## Follow-up questions
- Where should transition rules live?
- How would events be published atomically with state changes?
