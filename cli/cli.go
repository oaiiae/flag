// Package cli provides a simple way to create tree-like command-line interfaces
// while staying as close as possible to the standard [flag] package.
package cli

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"slices"
)

// Command is the basic building block of command-line interfaces.
type Command struct {
	// Name of the command
	Name string

	// A simple description of what the command does.
	// Used when printing usage.
	Usage string

	// Arguments placeholders when printing command usage.
	UsageArgs string

	// Function provided by the user to define flags of the current command.
	FlagSet func(fs *flag.FlagSet)

	// Flag to environment variable mappings.
	// It allows users to create flag arguments that may be set through environment as well.
	// The environment is parsed before command-line arguments.
	// May be nil.
	FlagEnvironment map[string]string

	// List of flags required when parsing arguments.
	// Allows users to early fail if a critical flag is not set.
	FlagRequired []string

	// Optional invocation function. May be useful for executing code before and/or after every subcommand.
	//
	// func(ctx context.Context, sub *cli.Command, args []string) error {
	// 	db, err := sql.Open("postgres", cli.Get(ctx, "dsn").(string))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	defer db.Close()
	// 	return sub.Run(context.WithValue(ctx, dbKey{}, db), args)
	// }
	Invoke func(ctx context.Context, sub *Command, args []string) error

	// Subcommands definitions.
	Subcommands []Command

	// Command function to run.
	Func func(ctx context.Context, args []string) error
}

// Usage is the function called when an error occurs when parsing flags or when help is requested.
// It may be customized by the user.
var Usage = func(c *Command, fs *flag.FlagSet) { //nolint: gochecknoglobals // mimicking [flag.Usage] global
	w := fs.Output()

	usage := []any{"Usage:", c.Name}
	if c.FlagSet != nil {
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

	if c.FlagSet != nil {
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Options:")
		fs.PrintDefaults()
	}

	if len(c.Subcommands) > 0 {
		fmt.Fprintln(w, "")
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

	if c.FlagSet != nil {
		c.FlagSet(fs)
	}

	for name, envname := range c.FlagEnvironment {
		if env, ok := os.LookupEnv(envname); ok {
			if err := fs.Set(name, env); err != nil {
				return err
			}
		}
	}

	err := fs.Parse(args)
	if err != nil {
		return err
	}
	args = fs.Args()

	actualFlags := []string{}
	fs.Visit(func(f *flag.Flag) { actualFlags = append(actualFlags, f.Name) })
	for _, name := range c.FlagRequired {
		if !slices.Contains(actualFlags, name) {
			return errors.New("missing required flag -" + name)
		}
	}

	flags, _ := ctx.Value(ctxFlags{}).(map[string]*flag.Flag)
	if flags == nil {
		flags = make(map[string]*flag.Flag)
		ctx = context.WithValue(ctx, ctxFlags{}, flags)
	}
	fs.VisitAll(func(f *flag.Flag) { flags[f.Name] = f })

	if len(args) > 0 {
		i := slices.IndexFunc(c.Subcommands, func(c Command) bool { return c.Name == args[0] })
		if i != -1 {
			return c.invoke(ctx, i, args[1:])
		}
	}

	if c.Func != nil {
		return c.Func(ctx, args)
	}

	fs.Usage()
	return nil
}

// invoke runs the i-th subcommand, using the [Command]'s Invoke function if set.
func (c *Command) invoke(ctx context.Context, i int, args []string) error {
	if c.Invoke != nil {
		return c.Invoke(ctx, &c.Subcommands[i], args)
	}
	return c.Subcommands[i].Run(ctx, args)
}

type ctxFlags struct{}

// Flag returns the named [*flag.Flag] from the context.
func Flag(ctx context.Context, name string) *flag.Flag {
	flags, _ := ctx.Value(ctxFlags{}).(map[string]*flag.Flag)
	if flags == nil {
		return nil
	}
	return flags[name]
}

// Get looks for the named [*flag.Flag] and returns the result of calling Get on its [flag.Value].
// It returns nil if:
//   - the specified [*flag.Flag] was not found
//   - its [flag.Value] does not implement [flag.Getter]
//   - the [flag.Getter] itself returns nil
func Get(ctx context.Context, name string) any {
	f := Flag(ctx, name)
	if f == nil {
		return nil
	}

	g, ok := f.Value.(flag.Getter)
	if !ok {
		return nil
	}

	return g.Get()
}
