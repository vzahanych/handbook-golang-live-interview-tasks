# Type switch versus reflection

## Live interview task
Format a small closed set of values, then explain when reflection would be justified.

## Interview notes / pitfalls
- Type switches are compile-time readable and appropriate for closed sets.
- Reflection is useful for open-ended structural inspection such as encoders.
- Generics solve operations shared by compile-time type sets, not arbitrary runtime metadata.
- Avoid reflection when a small interface expresses the behavior directly.
