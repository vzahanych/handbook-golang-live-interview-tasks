# Copy matching struct fields

## Live interview task
Copy exported fields with identical names and assignable types from one struct to another.

## Interview notes / pitfalls
- Destination must be a non-nil pointer to a settable struct.
- Use `AssignableTo`, not string comparisons of type names.
- Embedded fields, tags, conversion, pointers, and deep copy semantics need explicit rules.
- Cache reflection metadata if this runs frequently.
