# Read lines from file or stdin

## Live interview task
Accept `-input path`; use stdin when the value is `-`, and return all non-empty trimmed lines.

## Interview notes / pitfalls
- Inject `io.Reader` where possible instead of hard-coding `os.Stdin`.
- Increase `bufio.Scanner` capacity when lines may exceed 64 KiB.
- Close only files opened by the function, not caller-owned stdin.
- Propagate `scanner.Err()` after iteration.
