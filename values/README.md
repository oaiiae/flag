# flag/values

A package providing enhanced flag parsing for Go's standard `flag` package with support for:

- **Generic Values**: Define flags for any type with custom parse/format functions
- **Basic Types**: All Go basic types (int, string, bool, float, ...)
- **Standard Library Types**: Support for time.Duration, time.Time, net/url.URL, net/netip.Addr, net/mail.Address, ...
- **Collections**: Support for both repeated flags (lists) and delimited values (slices)
- **Simplified Registration**: Streamlined API for registering complex flag types

The package extensively uses generics to simplify the API.

## Usage

Use the value builders directly with `flag.Var()`:

```go
func main() {
    // Basic scalar value
    flag.Var(values.Basic[int](), "count", "number of items")

    // Time duration with automatic parsing
    flag.Var(values.Duration(), "timeout", "operation timeout")

    // List flags (can be specified multiple times)
    flag.Var(values.BasicList[string](), "tag", "tags (can be specified multiple times)")

    // URL with automatic parsing and validation
    flag.Var(values.Stringer[*url.URL](url.Parse), "endpoint", "API endpoint URL")

    // Slice flag (comma-separated values in one flag)
    flag.Var(values.BasicSlice[string](","), "regions", "comma-separated list of regions")

    // Custom type with custom parser
    type LogLevel int
    flag.Var(values.Generic(
        func(s string) (LogLevel, error) {
            // Custom parsing logic
            switch s {
            case "debug":
                return 0, nil
            case "info":
                return 1, nil
            default:
                return 0, fmt.Errorf("unknown log level: %s", s)
            }
        },
        func(l LogLevel) string {
            switch l {
            case 0:
                return "debug"
            case 1:
                return "info"
            default:
                return "unknown"
            }
        },
    ), "log-level", "logging level")

    flag.Parse()

    // Print all flag values
    flag.VisitAll(func(f *flag.Flag) {
        fmt.Printf("%s: %v\n", f.Name, f.Value)
    })
}
```

This is the most flexible way for defining complex flag values. Alternatively, the `Registerer` gives up the generic approach but provides a simpler API for common types, akin to `flag.FlagSet`:

```go
func main() {
    // Create a registerer for the standard flag package
    r := values.FlagSetRegisterer(flag.CommandLine)

    // Register various flag types
    count := r.Int("count", 10, "number of items")
    timeout := r.Duration("timeout", 5*time.Second, "operation timeout")
    urls := r.URLList("url", nil, "URLs to process (can be specified multiple times)")

    flag.Parse()

    // Use the parsed values
    fmt.Printf("Count: %d\n", *count)
    fmt.Printf("Timeout: %v\n", *timeout)
    fmt.Printf("URLs: %v\n", *urls)
}
```
