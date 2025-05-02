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
	flag.Var(values.Basic[int](), "count", "number of items")
	flag.Var(values.Duration(), "timeout", "operation timeout")
	flag.Var(values.BasicList[string](), "tag", "tags (can be specified multiple times)")
	flag.Var(values.Stringer(url.Parse), "endpoint", "API endpoint URL")
	flag.Var(values.BasicSlice[string](","), "regions", "comma-separated list of regions")

	// Custom type with custom parser
	type LogLevel int
	flag.Var(values.Generic(
		func(s string) (LogLevel, error) {
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
	flag.VisitAll(func(f *flag.Flag) { fmt.Printf("%s: %v\n", f.Name, f.Value) })
}
```
```
$ go run . -count 12 -endpoint http://example.com -tag foo -tag bar -regions euw,eune
count: 12
endpoint: http://example.com
log-level:
regions: euw,eune
tag: [foo bar]
timeout:
```

This is the most flexible way for defining complex flag values. Alternatively, the `Registerer` gives up the generic approach but provides a simpler API for common types, akin to `flag.FlagSet`:

```go
func main() {
	reg := values.FlagSetRegisterer(flag.CommandLine)

	count := reg.Int("count", 10, "number of items")
	timeout := reg.Duration("timeout", 5*time.Second, "operation timeout")
	bind := reg.IPAddrPort("bind", netip.MustParseAddrPort("0.0.0.0:8080"), "binding address")

	flag.Parse()
	fmt.Printf("Count: %d\n", *count)
	fmt.Printf("Timeout: %v\n", *timeout)
	fmt.Printf("Bind: %v\n", *bind)
}
```
```
$ go run . -count 12 -bind 10.0.0.1:80
Count: 12
Timeout: 5s
Bind: 10.0.0.1:80
```
