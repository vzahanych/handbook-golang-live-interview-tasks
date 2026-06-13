# Subcommands with FlagSet

## Live interview task
Build `todo add -text ...` and `todo list -json` without a third-party CLI package.

## Solution outline
- Require at least one positional argument for the command.
- Create one `flag.FlagSet` per subcommand with `ContinueOnError`.
- Parse only `os.Args[2:]` in the selected set.
- Return errors from `run(args, stdout, stderr)` so `main` only maps them to exit codes.

## Interview notes / pitfalls
- The global `flag` set makes tests interfere with one another.
- Separate parsing, business logic, and process exit for testability.
