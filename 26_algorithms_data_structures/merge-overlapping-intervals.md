# Merge overlapping intervals

## Live interview task
Merge inclusive integer intervals and return them sorted by start.

## Solution outline
- Clone input if mutation is not allowed.
- Sort by start, then end.
- Extend the last output interval when the next interval overlaps; otherwise append.

## Interview notes / pitfalls
- Complexity is `O(n log n)` due to sorting.
- Define whether touching intervals such as `[1,2]` and `[3,4]` merge.
