// Package cli provides a simple way to create tree-like command-line interfaces
// while staying as close as possible to the standard [flag] package.
package cli

import (
	"context"
	"flag"
	"fmt"
	"os"
	"slices"
)

// Command is the basic building block of command-line interfaces.
type Command struct {
	// Name of the command.
	Name string
	// Usage description of the command.
	Usage string
	// Usage command argument placeholders.
	UsageArgs string
	// Flags definition function for this command.
	Flags func(fs *flag.FlagSet)
	// Flag to environment variable mappings.
	// Allows users to define flags that may be set through environment as well.
	// Environment is parsed before command-line arguments.
	FlagsEnvMap map[string]string
	// Flags marked as required, enabling early failure.
	FlagsRequired []string
	// Function for adding custom code and passing values around the execution
	// of the actual [Command]. Any error returned here is reported by the
	// [Command.Run] method.
	//
	// For instance, this can be useful for handling shared resources:
	//
	// func(parent context.Context, run func(ctx context.Context) error) error {
	// 	db, err := sql.Open("postgres", cli.Get(parent, "dsn").(string))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	defer db.Close()
	// 	return run(context.WithValue(parent, dbKey{}, db))
	// }
	//
	// This opens a database handler and stores it in the context, making it
	// available from this node to the remaining of the tree. This approach
	// supports deferred statements, keeping cleanup code idiomatic.
	RunContext func(parent context.Context, run func(ctx context.Context) error) error
	// Subcommands definitions.
	Subcommands []*Command
	// Command function to run.
	Func func(ctx context.Context, args []string) error
}

// Usage is the function called when an error occurs when parsing flags or when help is requested.
// It may be customized by the user.
var Usage = func(c *Command, fs *flag.FlagSet) { //nolint: gochecknoglobals // mimicking [flag.Usage] global
	w := fs.Output()

	usage := []any{"Usage:", c.Name}
	if c.Flags != nil {
		usage = append(usage, "[options]")
	}
	if len(c.Subcommands) > 0 {
		usage = append(usage, "COMMAND")
	}
	usage = append(usage, c.UsageArgs)
	fmt.Fprintln(w, usage...)

	if c.Usage != "" {
		fmt.Fprintln(w)
		fmt.Fprintln(w, c.Usage)
	}

	if c.Flags != nil {
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Options:")
		fs.PrintDefaults()
	}

	if len(c.Subcommands) > 0 {
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Commands:")

		lines := []fmt.Stringer{}
		width := 0
		for _, c := range c.Subcommands {
			lines = append(lines, stringerFunc(func() string { return fmt.Sprintf("  %-*s    %s", width, c.Name, c.Usage) }))
			width = max(width, len(c.Name))
		}
		for _, line := range lines {
			fmt.Fprintln(w, line)
		}
	}
}

type stringerFunc func() string

func (f stringerFunc) String() string { return f() }

// Run runs the command tree by parsing environment & flag arguments into [flag.Value] and store them in the context.
// If a subcommand can be run using the remaining non-flag arguments, then it is run, otherwise it runs the [Command]'s function.
// If there is no function to run, it prints usage and returns.
func (c *Command) Run(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet(c.Name, flag.ContinueOnError)
	fs.Usage = func() { Usage(c, fs) }

	if c.Flags != nil {
		c.Flags(fs)
	}

	for name, envname := range c.FlagsEnvMap {
		if env, ok := os.LookupEnv(envname); ok {
			if err := fs.Set(name, env); err != nil {
				return err
			}
		}
	}

	err := fs.Parse(args)
	if err != nil {
		return err
	} else { //nolint: revive // keeps code of required-flag checks within a block
		placed := make([]string, 0, fs.NFlag())
		fs.Visit(func(f *flag.Flag) { placed = append(placed, f.Name) })
		for _, name := range c.FlagsRequired {
			if !slices.Contains(placed, name) {
				return fmt.Errorf("missing required flag -%s", name)
			}
		}
	}
	args = fs.Args()

	flags, _ := ctx.Value(ctxFlags{}).(map[string]*flag.Flag)
	if flags == nil {
		flags = make(map[string]*flag.Flag)
		ctx = context.WithValue(ctx, ctxFlags{}, flags)
	}
	fs.VisitAll(func(f *flag.Flag) { flags[f.Name] = f })

	runContext := defaultRunContext
	if c.RunContext != nil {
		runContext = c.RunContext
	}

	return runContext(ctx, func(child context.Context) error {
		i := slices.IndexFunc(c.Subcommands, func(c *Command) bool { return len(args) > 0 && args[0] == c.Name })
		switch {
		case i != -1: // the remaining arguments matched a subcommand
			return c.Subcommands[i].Run(child, args[1:])
		case c.Func != nil: // no subcommand could be run, fallback to this command action
			return c.Func(child, args)
		default: // nothing could be done, print usage
			fs.Usage()
			return fmt.Errorf("cli cannot proceed with arguments %v", args)
		}
	})
}

// defaultRunContext is the default implementation of [Command.RunContext].
// It simply runs the callback without modifying anything.
func defaultRunContext(parent context.Context, run func(ctx context.Context) error) error {
	return run(parent)
}

type ctxFlags struct{}

// Get looks for the named flag and returns its value.
// It returns nil if:
//   - the specified [flag.Flag] was not found
//   - its [flag.Value] does not implement [flag.Getter]
//   - the [flag.Getter] itself returns nil
func Get(ctx context.Context, name string) any {
	flags, _ := ctx.Value(ctxFlags{}).(map[string]*flag.Flag)
	if flags == nil {
		return nil
	}

	f := flags[name]
	if f == nil {
		return nil
	}

	g, ok := f.Value.(flag.Getter)
	if !ok {
		return nil
	}

	return g.Get()
}
