# Parallel deduplication with sharded maps

## Live interview task
Deduplicate a large `[]string` concurrently while preserving first-occurrence order.

## Solution outline
1. Split the input into indexed chunks.
2. Workers insert values into hash-selected shards, each guarded by its own mutex.
3. Store the smallest observed index for every value.
4. Merge shard entries into `(index, value)` pairs and sort by index.

## Interview notes / pitfalls
- A single concurrent append corrupts the result unless synchronized.
- A plain concurrent map loses stable ordering.
- Sharding reduces lock contention but adds complexity; benchmark before using it.
- Hash collisions are safe when the full string remains the map key.

## Follow-up questions
- Can a two-pass sequential merge be faster?
- How would memory usage change for millions of unique strings?
- Would `sync.Map` be appropriate for write-once keys?
