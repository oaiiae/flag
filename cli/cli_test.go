package cli_test

import (
	"flag"
	"os"

	"gitlab.com/oaiiae/flag/cli"
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
