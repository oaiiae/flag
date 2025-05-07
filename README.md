# flag/*

Flag is a collection of packages extending the usability of the standard Go `flag` package. It has zero dependencies and aims to provide a simple, Go-idiomatic framework for implementing complex command-line interfaces.

## flag/cli

Package `flag/cli` provides a very simple interface for building command-lines applications with:

- **Command trees**: build hierarchical command structures with subcommands
- **Flag compatibility**: compatible with Go's standard `flag` package
- **Environment support**: map flags to environment variables
- **Required flags**: enforce required flags for early failure
- **Run context**: wrap commands with custom code

```go
func main() {
    cmd := &cli.Command{
        Name:      os.Args[0],
        Usage:     "A sample application",
        UsageArgs: "[arguments]",
        Flags: func(fs *flag.FlagSet) {
            fs.Bool("verbose", false, "enable verbose output")
        },
        RunContext: func(parent context.Context, run func(ctx context.Context) error) error {
            ctx, cancel := signal.NotifyContext(parent, os.Interrupt)
            defer cancel()
            return run(ctx)
        },
        Subcommands: []*cli.Command{
            {
                Name:  "serve",
                Usage: "Start the server",
                Flags: func(fs *flag.FlagSet) {
                    fs.String("config", "config.yaml", "configuration file (env $MYAPP_CONFIG)")
                    fs.Int("port", 8080, "port to listen on")
                },
                FlagsEnvMap: map[string]string{
                    "config": "MYAPP_CONFIG",
                },
                FlagsRequired: []string{"config"},
                Func: func(ctx context.Context, args []string) error {
                    fmt.Printf("starting server on port %d with config %s (verbose: %v)\n",
                        cli.Get(ctx, "port").(int),
                        cli.Get(ctx, "config").(string),
                        cli.Get(ctx, "verbose").(bool),
                    )
                    return nil
                },
            },
            {
                Name:  "version",
                Usage: "Show version information",
                Func: func(ctx context.Context, args []string) error {
                    fmt.Println("vX.Y.Z")
                    return nil
                },
            },
        },
    }

    err := cmd.Run(context.Background(), os.Args[1:])
    if err != nil && err != flag.ErrHelp {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(2)
    }
}
```
```
$ go run . -h
Usage: /tmp/go-build690470450/b001/exe/main [options] COMMAND [arguments]

A sample application

Options:
  -verbose
    	enable verbose output

Commands:
  serve      Start the server
  version    Show version information
```
```
$ go run . serve -h
Usage: serve [options]

Start the server

Options:
  -config string
    	configuration file (env $MYAPP_CONFIG) (default "config.yaml")
  -port int
    	port to listen on (default 8080)
```

## flag/values

Package `flag/values` provides generic implementations of the standard Go `flag.Value` interface with support for:

- **Generic values**: define flags for any type with custom parse/format functions
- **Basic types**: all Go basic types (int, string, bool, float, ...)
- **Standard library types**: support for time.time, net/url.URL, net/netip.Addr, net/mail.Address, ...
- **Collections**: support for both repeated flags (lists) and delimited values (slices)

```go
func main() {
    flag.Var(values.Basic[int](), "count", "number of items")
    flag.Var(values.BasicList[string](), "tag", "tags (can be specified multiple times)")
    flag.Var(values.BasicSlice[string](","), "regions", "comma-separated list of regions")
    flag.Var(values.Stringer(url.Parse), "endpoint", "API endpoint URL")
    flag.Parse()
    flag.VisitAll(func(f *flag.Flag) { fmt.Printf("%s: %v\n", f.Name, f.Value.(flag.Getter).Get()) })
}
```
```
$ go run . -count 12 -endpoint http://example.com -tag foo -tag bar -regions euw,eune
count: 12
endpoint: http://example.com
regions: [euw eune]
tag: [foo bar]
```

Alternatively, the `Registerer` provides an interface analogous to `flag.FlagSet` simplifying registration for common types:

```go
func main() {
    var (
        reg   = values.FlagSetRegisterer(flag.CommandLine)
        count = reg.Int("count", 10, "number of items")
        email = reg.MailAddr("email", &mail.Address{}, "contact email")
        bind  = reg.IPAddrPort("bind", netip.MustParseAddrPort("0.0.0.0:8080"), "binding address")
    )
    flag.Parse()
    fmt.Printf("Count: %d\n", *count)
    fmt.Printf("Email: %v\n", *email)
    fmt.Printf("Bind: %v\n", *bind)
}
```
```
$ go run . -count 12 -bind 10.0.0.1:80 -email foo@example.com
Count: 12
Email: <foo@example.com>
Bind: 10.0.0.1:80
```
