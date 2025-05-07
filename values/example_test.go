package values_test

import (
	"flag"
	"net/netip"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/oaiiae/flag/values"
)

func ExampleGeneric_usage() {
	type pair struct{ a, b string }
	parse := func(s string) (pair, error) { a, b, _ := strings.Cut(s, ":"); return pair{a, b}, nil }
	format := func(p pair) string { return p.a + ":" + p.b }

	fs := flag.NewFlagSet("test", flag.ContinueOnError)

	fs.Var(values.Generic(parse, format), "generic", "usage")
	fs.Var(values.GenericVar(&pair{"foo", "bar"}, parse, format), "generic-var", "usage")
	fs.Var(values.GenericList(parse, format), "generic-list", "usage")
	fs.Var(values.GenericListVar(&[]pair{{"foo", "bar"}, {"quu", "quux"}}, parse, format), "generic-list-var", "usage")
	fs.Var(values.GenericSlice(",", parse, format), "generic-slice", "usage")
	fs.Var(values.GenericSliceVar(&[]pair{{"foo", "bar"}, {"quu", "quux"}}, ",", parse, format), "generic-slice-var", "usage")

	fs.SetOutput(os.Stdout)
	fs.PrintDefaults()

	// Output:
	//   -generic value
	//     	usage
	//   -generic-list value
	//     	usage
	//   -generic-list-var value
	//     	usage (default [foo:bar quu:quux])
	//   -generic-slice value
	//     	usage
	//   -generic-slice-var value
	//     	usage (default foo:bar,quu:quux)
	//   -generic-var value
	//     	usage (default foo:bar)
}

func ExampleBasic_usage() {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)

	c := 12 + 42i
	fs.Var(values.Basic[complex128](), "complex", "usage")
	fs.Var(values.BasicVar(&c), "complex-var", "usage")
	fs.Var(values.BasicList[complex128](), "complex-list", "usage")
	fs.Var(values.BasicListVar(&[]complex128{c, c}), "complex-list-var", "usage")
	fs.Var(values.BasicSlice[complex128](","), "complex-slice", "usage")
	fs.Var(values.BasicSliceVar(&[]complex128{c, c}, ","), "complex-slice-var", "usage")

	fs.SetOutput(os.Stdout)
	fs.PrintDefaults()

	// Output:
	//   -complex value
	//     	usage
	//   -complex-list value
	//     	usage
	//   -complex-list-var value
	//     	usage (default [(12+42i) (12+42i)])
	//   -complex-slice value
	//     	usage
	//   -complex-slice-var value
	//     	usage (default (12+42i),(12+42i))
	//   -complex-var value
	//     	usage (default (12+42i))
}

func ExampleStringer_usage() {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)

	ip := netip.MustParseAddr("1.2.3.4")
	fs.Var(values.Stringer(netip.ParseAddr), "ip", "usage")
	fs.Var(values.StringerVar(&ip, netip.ParseAddr), "ip-var", "usage")
	fs.Var(values.StringerList(netip.ParseAddr), "ip-list", "usage")
	fs.Var(values.StringerListVar(&[]netip.Addr{ip, ip}, netip.ParseAddr), "ip-list-var", "usage")
	fs.Var(values.StringerSlice(",", netip.ParseAddr), "ip-slice", "usage")
	fs.Var(values.StringerSliceVar(&[]netip.Addr{ip, ip}, ",", netip.ParseAddr), "ip-slice-var", "usage")

	fs.SetOutput(os.Stdout)
	fs.PrintDefaults()

	// Output:
	//   -ip value
	//     	usage
	//   -ip-list value
	//     	usage
	//   -ip-list-var value
	//     	usage (default [1.2.3.4 1.2.3.4])
	//   -ip-slice value
	//     	usage
	//   -ip-slice-var value
	//     	usage (default 1.2.3.4,1.2.3.4)
	//   -ip-var value
	//     	usage (default 1.2.3.4)
}

func ExampleStringer_usage2() {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)

	u := &url.URL{Scheme: "foo", Path: "bar"}
	fs.Var(values.Stringer(url.Parse), "url", "usage")
	fs.Var(values.StringerVar(&u, url.Parse), "url-var", "usage")
	fs.Var(values.StringerList(url.Parse), "url-list", "usage")
	fs.Var(values.StringerListVar(&[]*url.URL{u, u}, url.Parse), "url-list-var", "usage")
	fs.Var(values.StringerSlice(",", url.Parse), "url-slice", "usage")
	fs.Var(values.StringerSliceVar(&[]*url.URL{u, u}, ",", url.Parse), "url-slice-var", "usage")

	fs.SetOutput(os.Stdout)
	fs.PrintDefaults()

	// Output:
	//   -url value
	//     	usage
	//   -url-list value
	//     	usage
	//   -url-list-var value
	//     	usage (default [foo://bar foo://bar])
	//   -url-slice value
	//     	usage
	//   -url-slice-var value
	//     	usage (default foo://bar,foo://bar)
	//   -url-var value
	//     	usage (default foo://bar)
}

func ExampleTime_usage() {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)

	t := time.Date(2025, 2, 1, 12, 34, 56, 0, time.UTC)
	fs.Var(values.Time(time.RFC3339), "time", "usage")
	fs.Var(values.TimeVar(&t, time.RFC3339), "time-var", "usage")
	fs.Var(values.TimeList(time.RFC3339), "time-list", "usage")
	fs.Var(values.TimeListVar(&[]time.Time{t, t}, time.RFC3339), "time-list-var", "usage")
	fs.Var(values.TimeSlice(",", time.RFC3339), "time-slice", "usage")
	fs.Var(values.TimeSliceVar(&[]time.Time{t, t}, ",", time.RFC3339), "time-slice-var", "usage")

	fs.SetOutput(os.Stdout)
	fs.PrintDefaults()

	// Output:
	//   -time value
	//     	usage
	//   -time-list value
	//     	usage
	//   -time-list-var value
	//     	usage (default [2025-02-01T12:34:56Z 2025-02-01T12:34:56Z])
	//   -time-slice value
	//     	usage
	//   -time-slice-var value
	//     	usage (default 2025-02-01T12:34:56Z,2025-02-01T12:34:56Z)
	//   -time-var value
	//     	usage (default 2025-02-01T12:34:56Z)
}

func ExampleDuration_usage() {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)

	d := 1234 * time.Second
	fs.Var(values.Duration(), "duration", "usage")
	fs.Var(values.DurationVar(&d), "duration-var", "usage")
	fs.Var(values.DurationList(), "duration-list", "usage")
	fs.Var(values.DurationListVar(&[]time.Duration{d, d}), "duration-list-var", "usage")
	fs.Var(values.DurationSlice(","), "duration-slice", "usage")
	fs.Var(values.DurationSliceVar(&[]time.Duration{d, d}, ","), "duration-slice-var", "usage")

	fs.SetOutput(os.Stdout)
	fs.PrintDefaults()

	// Output:
	//   -duration value
	//     	usage
	//   -duration-list value
	//     	usage
	//   -duration-list-var value
	//     	usage (default [20m34s 20m34s])
	//   -duration-slice value
	//     	usage
	//   -duration-slice-var value
	//     	usage (default 20m34s,20m34s)
	//   -duration-var value
	//     	usage (default 20m34s)
}
