# Debounce events

## Live interview task
Emit the latest input only after no new value arrives for a duration.

## Solution outline
- Keep one timer owned by the debounce goroutine.
- On input, save the latest value and safely stop/reset the timer.
- On timer fire, emit the saved value with context-aware send.
- On input close, define whether to flush the pending value.

## Interview notes / pitfalls
- A new goroutine and `time.After` per event causes avoidable allocations and races.
- Slow consumers require a documented backpressure or drop policy.
