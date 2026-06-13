# Unsafe string and byte conversion

## Live interview task
Review a zero-allocation `string` to `[]byte` conversion using `unsafe` and explain why it is dangerous.

## Interview notes / pitfalls
- Strings are immutable by contract; mutating aliased bytes can corrupt memory assumptions.
- Lifetime and garbage-collector visibility must remain valid.
- `unsafe.String`, `unsafe.Slice`, and layout assumptions are version-sensitive tools.
- Prefer ordinary conversion unless profiling proves the allocation matters and ownership is airtight.
