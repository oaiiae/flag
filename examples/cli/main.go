package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/oaiiae/flag/cli"
)

func main() {
	cli := &cli.Command{
		Name:  os.Args[0],
		Usage: "My super CLI",
		Flags: func(fs *flag.FlagSet) {
			fs.BoolFunc("version", "show version & exit", func(string) error {
				fmt.Println("version number")
				os.Exit(0)
				return nil
			})
			fs.Int("val", 42, "a configurable value")
			fs.Bool("v", false, "verbose switch")
			fs.Duration("dur", 0, "a duration")
		},
		RunContext: func(parent context.Context, run func(ctx context.Context) error) error {
			ctx, cancel := signal.NotifyContext(parent, os.Interrupt)
			defer cancel()
			return run(ctx)
		},
		Subcommands: []*cli.Command{
			{
				Name:      "dump",
				Usage:     "Dump CLI options & arguments",
				UsageArgs: "arguments...",
				Flags: func(fs *flag.FlagSet) {
					fs.String("foo", "", "a foo-lish option")
				},
				FlagsRequired: []string{"foo"},
				Func: func(ctx context.Context, args []string) error {
					fmt.Println("val", cli.Get(ctx, "val").(int))
					fmt.Println("dur", cli.Get(ctx, "dur").(time.Duration))
					fmt.Println("foo", cli.Get(ctx, "foo").(string))
					fmt.Println("arguments", args)
					return nil
				},
			},
			{
				Name:  "wait",
				Usage: "Wait until context is done",
				Flags: func(fs *flag.FlagSet) {
					fs.Duration("timeout", 10*time.Second, "wait up to this duration")
				},
				Func: func(ctx context.Context, _ []string) error {
					ctx, cancel := context.WithTimeout(ctx, cli.Get(ctx, "timeout").(time.Duration))
					defer cancel()
					fmt.Println("waiting...")
					<-ctx.Done()
					return ctx.Err()
				},
			},
		},
	}

	err := cli.Run(context.TODO(), os.Args[1:])
	if err != nil && err != flag.ErrHelp { //nolint: errorlint // explicit check
		fmt.Println(err)
		os.Exit(2) // exit immediately (skips deferred statements)
	}
}
