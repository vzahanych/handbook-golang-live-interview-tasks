# Golang live interview tasks

A collection of **228** live-coding and discussion tasks covering the Go language, standard library, concurrency, HTTP, scraping, data access, testing, performance, and service design.

Each task contains a prompt, concepts or requirements, a candidate solution or solution outline, and interview pitfalls. Most tasks use only the standard library; Colly tasks require `github.com/gocolly/colly/v2`.

- [Complete task catalog](./TASK_CATALOG.md)
- [Official learning and API references](./REFERENCES.md)
- [Machine-readable manifest](./manifest.json)

## Categories

- [01_language_basics](./01_language_basics/README.md) - Language basics and source organization (8 tasks)
- [02_control_flow](./02_control_flow/README.md) - Control flow and statements (9 tasks)
- [03_types_values_memory](./03_types_values_memory/README.md) - Types, values, identity, assignability, memory model (7 tasks)
- [04_arrays_slices_strings](./04_arrays_slices_strings/README.md) - Arrays, slices, strings and bytes (15 tasks)
- [05_maps_sets](./05_maps_sets/README.md) - Maps, sets and grouping (10 tasks)
- [06_functions_methods_interfaces](./06_functions_methods_interfaces/README.md) - Functions, methods, interfaces (8 tasks)
- [07_generics](./07_generics/README.md) - Generics and type constraints (10 tasks)
- [08_errors_panic_defer](./08_errors_panic_defer/README.md) - Errors, panic, recover and defer (8 tasks)
- [09_concurrency_channels](./09_concurrency_channels/README.md) - Goroutines, channels and select (12 tasks)
- [10_context_cancellation](./10_context_cancellation/README.md) - Context, cancellation and timeouts (7 tasks)
- [11_sync_atomics](./11_sync_atomics/README.md) - sync package and atomics (8 tasks)
- [12_parallel_slice_calculations](./12_parallel_slice_calculations/README.md) - Parallel slice calculations (15 tasks)
- [13_http_servers_clients](./13_http_servers_clients/README.md) - HTTP servers and clients (14 tasks)
- [14_json_files_io](./14_json_files_io/README.md) - JSON, files and I/O (8 tasks)
- [15_testing_benchmarking](./15_testing_benchmarking/README.md) - Testing, benchmarking and fuzzing (7 tasks)
- [16_runtime_performance](./16_runtime_performance/README.md) - Runtime and performance interview tasks (9 tasks)
- [17_colly_scraping](./17_colly_scraping/README.md) - Colly scraping examples (22 tasks)
- [18_mini_projects](./18_mini_projects/README.md) - Small integrated projects (7 tasks)
- [19_cli_configuration](./19_cli_configuration/README.md) - CLI and configuration (5 tasks)
- [20_database_sql](./20_database_sql/README.md) - Database and SQL (5 tasks)
- [21_networking_protocols](./21_networking_protocols/README.md) - Networking and protocols (5 tasks)
- [22_time_scheduling](./22_time_scheduling/README.md) - Time and scheduling (5 tasks)
- [23_reflection_encoding](./23_reflection_encoding/README.md) - Reflection and encoding (5 tasks)
- [24_design_architecture](./24_design_architecture/README.md) - Design and architecture (5 tasks)
- [25_security_observability](./25_security_observability/README.md) - Security, reliability, and observability (5 tasks)
- [26_algorithms_data_structures](./26_algorithms_data_structures/README.md) - Algorithms and data structures (5 tasks)
- [27_os_processes](./27_os_processes/README.md) - OS, files, and processes (5 tasks)

## Regenerate indexes

```bash
go run ./tools/generate_catalog.go
```
