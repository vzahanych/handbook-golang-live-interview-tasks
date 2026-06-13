# Custom JSON time format

## Live interview task
Marshal a date as `YYYY-MM-DD` and reject timestamps containing a time component.

## Solution outline
- Define a named type wrapping `time.Time`.
- Implement `MarshalJSON` and `UnmarshalJSON`.
- Use `strconv.Quote`/`Unquote` and `time.Parse` rather than slicing JSON bytes blindly.
- Decide how zero values and `null` are represented.

## Interview notes / pitfalls
- Custom methods change behavior everywhere the type is encoded.
- Validate trailing data and return errors with field context at the API boundary.
