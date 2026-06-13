# Atomic file replacement

## Live interview task
Write configuration so readers see either the old complete file or the new complete file.

## Solution outline
- Create a temporary file in the destination directory.
- Set permissions, write, `Sync`, and close it.
- Rename it over the destination.
- Sync the directory when crash durability is required.

## Interview notes / pitfalls
- Rename atomicity and replacement behavior vary by filesystem and platform.
- Cross-filesystem rename is not atomic.
- Clean up the temporary file on every failure path.
