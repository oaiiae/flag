package cli_test

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/oaiiae/flag/cli"
	"github.com/stretchr/testify/require"
)

func ExampleUsage() {
	c := cli.Command{
		Name:      "foo",
		Usage:     "A foo-lish command",
		UsageArgs: "quu quux ...",
		Flags: func(fs *flag.FlagSet) {
			fs.Bool("bool", false, "a bool flag")
			fs.Int("int", 12, "an int flag")
		},
		Subcommands: []*cli.Command{
			{
				Name:  "sub1",
				Usage: "subcommand 1",
			},
			{
				Name:  "sub2",
				Usage: "subcommand 2",
			},
		},
	}

	fs := flag.NewFlagSet("", flag.PanicOnError)
	fs.SetOutput(os.Stdout)
	c.Flags(fs)

	cli.Usage(&c, fs)

	// Output:
	// Usage: foo [options] COMMAND quu quux ...
	//
	// A foo-lish command
	//
	// Options:
	//   -bool
	//     	a bool flag
	//   -int int
	//     	an int flag (default 12)
	//
	// Commands:
	//   sub1    subcommand 1
	//   sub2    subcommand 2
}

func TestCommandRun(t *testing.T) {
	c := cli.Command{
		Flags: func(fs *flag.FlagSet) {
			fs.Bool("bool", false, "a bool flag")
			fs.Int("int", 12, "an int flag")
		},
		FlagsRequired: []string{"bool"},
		RunContext: func(parent context.Context, run func(context.Context) error) error {
			return errors.Join(run(parent), errors.New("runcontext terminated"))
		},
		Subcommands: []*cli.Command{
			{
				Name: "sub",
				Func: func(ctx context.Context, args []string) error {
					return fmt.Errorf("sub terminated: %v %v %v", cli.Get(ctx, "bool"), cli.Get(ctx, "int"), args)
				},
			},
		},
		Func: func(ctx context.Context, args []string) error {
			return fmt.Errorf("foo terminated: %v %v %v", cli.Get(ctx, "bool"), cli.Get(ctx, "int"), args)
		},
	}

	t.Run("missing required flag", func(t *testing.T) {
		err := c.Run(context.Background(), []string{})
		require.ErrorContains(t, err, "missing required flag -bool")
	})

	t.Run("parses bool", func(t *testing.T) {
		err := c.Run(context.Background(), []string{"-bool"})
		require.ErrorContains(t, err, "foo terminated: true 12 []")
	})

	t.Run("reports parsing errors", func(t *testing.T) {
		err := c.Run(context.Background(), []string{"-bool=notbool"})
		require.ErrorContains(t, err, "invalid boolean value \"notbool\"")
	})

	t.Run("runs within context", func(t *testing.T) {
		err := c.Run(context.Background(), []string{"-bool", "-int", "42"})
		require.ErrorContains(t, err, "foo terminated: true 42 []\nruncontext terminated")
	})

	t.Run("runs subcommand", func(t *testing.T) {
		err := c.Run(context.Background(), []string{"-bool", "-int", "42", "sub", "foo", "bar", "baz"})
		require.ErrorContains(t, err, "sub terminated: true 42 [foo bar baz]\nruncontext terminated")
	})
}

func TestCommand_unable(t *testing.T) {
	c := cli.Command{}
	err := c.Run(context.Background(), []string{})
	require.ErrorContains(t, err, "cli cannot proceed with arguments []")
}

func TestGetValueNotFound(t *testing.T) {
	c := cli.Command{
		Func: func(ctx context.Context, _ []string) error {
			require.Nil(t, cli.Get(ctx, "foo"))
			return errors.New("command terminated")
		},
	}
	err := c.Run(context.Background(), []string{})
	require.ErrorContains(t, err, "command terminated")
}

func TestGetNotCliContext(t *testing.T) {
	require.Nil(t, cli.Get(context.Background(), "foo"))
}
