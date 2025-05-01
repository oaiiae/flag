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
		Name:      "myapp",
		Usage:     "A sample application",
		UsageArgs: "[arguments]",
		Flags: func(fs *flag.FlagSet) {
			fs.String("config", "config.yaml", "configuration file (env $MYAPP_CONFIG)")
			fs.Bool("verbose", false, "enable verbose output")
		},
		FlagsEnvMap: map[string]string{
			"config": "MYAPP_CONFIG",
		},
		FlagsRequired: []string{"config"},
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
					fs.Int("port", 8080, "port to listen on")
				},
				Func: func(ctx context.Context, args []string) error {
					// Access the flag values
					config := cli.Get(ctx, "config").(string)
					port := cli.Get(ctx, "port").(int)
					fmt.Printf("Starting server on port %d with config %s\n", port, config)
					return nil
				},
			},
			{
				Name:  "version",
				Usage: "Show version information",
				Func: func(ctx context.Context, args []string) error {
					fmt.Println("v1.0.0")
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
```
