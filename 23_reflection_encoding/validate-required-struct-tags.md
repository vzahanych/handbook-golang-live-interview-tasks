# Validate required struct tags

## Live interview task
Use reflection to report zero-valued fields tagged `validate:"required"`.

## Required behavior
- Accept a struct or pointer to struct and reject other inputs.
- Handle nil pointers without panicking.
- Use `reflect.Value.IsZero` and field metadata from `reflect.Type`.
- Skip unexported fields or explain how package visibility affects access.

## Interview notes / pitfalls
- Reflection panics when methods are used on the wrong kind.
- Recursive validation needs cycle detection.
- Generated or explicit validation is usually clearer on hot paths.
