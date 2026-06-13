# Generic LRU cache

## Live interview task
Implement a fixed-capacity generic LRU cache with `Get` and `Put` in amortized O(1).

## Required design
- Map keys to nodes in a doubly linked list.
- Keep most recently used at the front.
- Update recency on both reads and writes.
- Evict the back node after inserting beyond capacity.

## Interview notes / pitfalls
- Decide whether zero capacity is valid.
- Concurrency safety is an additional contract, not automatic.
- `container/list` stores `any`; a custom generic node avoids assertions.
