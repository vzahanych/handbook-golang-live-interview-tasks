# Union-find connected components

## Live interview task
Count connected components from pairs of integer node IDs.

## Interview notes / pitfalls
- Implement path compression and union by size or rank.
- Include isolated nodes supplied outside the edge list.
- Clarify whether unknown IDs are created lazily.
- Near-constant amortized operations do not mean worst-case literal O(1).
