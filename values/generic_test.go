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
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)

	type pair struct{ a, b string }
	parse := func(s string) (pair, error) { a, b, _ := strings.Cut(s, ":"); return pair{a, b}, nil }
	format := func(p pair) string { return p.a + ":" + p.b }

	values.Generic(fs, "generic", "usage", parse, format)
	values.GenericVar(fs, &pair{"foo", "bar"}, "generic-var", "usage", parse, format)
	values.Generics(fs, "generics", "usage", parse, format, values.Unsplit)
	values.GenericsVar(fs, &[]pair{{"foo", "bar"}, {"quu", "quux"}}, "generics-var", "usage", parse, format, values.Unsplit)

	ip := netip.MustParseAddr("1.2.3.4")
	values.Stringer(fs, "ip", "usage", netip.ParseAddr)
	values.StringerVar(fs, &ip, "ip-var", "usage", netip.ParseAddr)
	values.Stringers(fs, "ips", "usage", netip.ParseAddr, values.Split(","))
	values.StringersVar(fs, &[]netip.Addr{ip, ip}, "ips-var", "usage", netip.ParseAddr, values.Split(","))

	u := &url.URL{Scheme: "foo", Path: "bar"}
	values.Stringer(fs, "url", "usage", url.Parse)
	values.StringerVar(fs, &u, "url-var", "usage", url.Parse)
	values.Stringers(fs, "urls", "usage", url.Parse, values.Unsplit)
	values.StringersVar(fs, &[]*url.URL{u, u}, "urls-var", "usage", url.Parse, values.Unsplit)

	t := time.Date(2025, 2, 1, 12, 34, 56, 0, time.UTC)
	values.Time(fs, "time", "usage", time.RFC3339)
	values.TimeVar(fs, &t, "time-var", "usage", time.RFC3339)
	values.Times(fs, "times", "usage", time.RFC3339, values.Unsplit)
	values.TimesVar(fs, &[]time.Time{t, t}, "times-var", "usage", time.RFC3339, values.Unsplit)

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
