# Top K with a min-heap

## Live interview task
Return the largest K integers from a stream without storing the full stream.

## Solution outline
- Maintain a min-heap of at most K values using `container/heap`.
- Push until full; then replace the root only when a larger value arrives.
- Sort the final heap if deterministic descending output is required.

## Interview notes / pitfalls
- Complexity is `O(n log k)` time and `O(k)` space.
- Define behavior for `k <= 0`, duplicates, and fewer than K inputs.
