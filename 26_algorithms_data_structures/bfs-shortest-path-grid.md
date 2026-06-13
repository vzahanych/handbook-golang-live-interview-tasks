# BFS shortest path in a grid

## Live interview task
Find the minimum number of orthogonal moves between two open cells.

## Interview notes / pitfalls
- BFS is correct for unweighted edges; mark cells visited when enqueued.
- Use a slice with a head index as a queue to avoid repeated front deletion.
- Validate ragged rows, blocked endpoints, and start equal to target.
- Store parents when the actual path is required.
