# flag/cli

A minimal package providing a simple and flexible CLI framework for building command-line applications with:

- **Command Trees**: Build hierarchical command structures with subcommands
- **Flag Compatibility**: Works with Go's standard `flag` package
- **Environment Support**: Map flags to environment variables
- **Required Flags**: Enforce required flags for validation
- **Context Integration**: Pass values through context between commands
- **Custom Invocation**: Execute code before/after subcommands

## Usage

Create a command tree and run it:

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
					fmt.Printf("Starting server on port %d with config %s (verbose: %v)\n",
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
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
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

$ go run . serve -h
  Usage: serve [options]

  Start the server

  Options:
    -config string
      	configuration file (env $MYAPP_CONFIG) (default "config.yaml")
    -port int
      	port to listen on (default 8080)
```
