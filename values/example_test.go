package values_test

import (
	"flag"
	"net/netip"
	"net/url"
	"os"
	"strings"
	"time"

	"gitlab.com/oaiiae/flag/values"
)

func ExampleGeneric_usage() {
	type pair struct{ a, b string }
	parse := func(s string) (pair, error) { a, b, _ := strings.Cut(s, ":"); return pair{a, b}, nil }
	format := func(p pair) string { return p.a + ":" + p.b }

	fs := flag.NewFlagSet("test", flag.ContinueOnError)

	fs.Var(values.Generic(parse, format), "generic", "usage")
	fs.Var(values.GenericVar(&pair{"foo", "bar"}, parse, format), "generic-var", "usage")
	fs.Var(values.Generics(parse, format, values.Unsplit), "generics", "usage")
	fs.Var(values.GenericsVar(&[]pair{{"foo", "bar"}, {"quu", "quux"}}, parse, format, values.Unsplit), "generics-var", "usage")

	ip := netip.MustParseAddr("1.2.3.4")
	fs.Var(values.Stringer(netip.ParseAddr), "ip", "usage")
	fs.Var(values.StringerVar(&ip, netip.ParseAddr), "ip-var", "usage")
	fs.Var(values.Stringers(netip.ParseAddr, values.Split(",")), "ips", "usage")
	fs.Var(values.StringersVar(&[]netip.Addr{ip, ip}, netip.ParseAddr, values.Split(",")), "ips-var", "usage")

	u := &url.URL{Scheme: "foo", Path: "bar"}
	fs.Var(values.Stringer(url.Parse), "url", "usage")
	fs.Var(values.StringerVar(&u, url.Parse), "url-var", "usage")
	fs.Var(values.Stringers(url.Parse, values.Unsplit), "urls", "usage")
	fs.Var(values.StringersVar(&[]*url.URL{u, u}, url.Parse, values.Unsplit), "urls-var", "usage")

	t := time.Date(2025, 2, 1, 12, 34, 56, 0, time.UTC)
	fs.Var(values.Time(time.RFC3339), "time", "usage")
	fs.Var(values.TimeVar(&t, time.RFC3339), "time-var", "usage")
	fs.Var(values.Times(time.RFC3339, values.Unsplit), "times", "usage")
	fs.Var(values.TimesVar(&[]time.Time{t, t}, time.RFC3339, values.Unsplit), "times-var", "usage")

	fs.SetOutput(os.Stdout)
	fs.PrintDefaults()

	// Output:
	//   -generic value
	//     	usage
	//   -generic-var value
	//     	usage (default foo:bar)
	//   -generics value
	//     	usage
	//   -generics-var value
	//     	usage (default [foo:bar quu:quux])
	//   -ip value
	//     	usage
	//   -ip-var value
	//     	usage (default 1.2.3.4)
	//   -ips value
	//     	usage
	//   -ips-var value
	//     	usage (default [1.2.3.4 1.2.3.4])
	//   -time value
	//     	usage
	//   -time-var value
	//     	usage (default 2025-02-01T12:34:56Z)
	//   -times value
	//     	usage
	//   -times-var value
	//     	usage (default [2025-02-01T12:34:56Z 2025-02-01T12:34:56Z])
	//   -url value
	//     	usage
	//   -url-var value
	//     	usage (default foo://bar)
	//   -urls value
	//     	usage
	//   -urls-var value
	//     	usage (default [foo://bar foo://bar])
}
