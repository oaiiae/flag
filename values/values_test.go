package values_test

import (
	"flag"
	"net/netip"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"gitlab.com/oaiiae/flag/values"
)

func TestParse(t *testing.T) {
	type pair struct{ a, b string }
	parse := func(s string) (pair, error) { a, b, _ := strings.Cut(s, ":"); return pair{a, b}, nil }
	format := func(p pair) string { return p.a + ":" + p.b }

	t.Run("generic", func(t *testing.T) {
		fs := flag.NewFlagSet("", flag.PanicOnError)
		v := values.Generic(fs, "v", "", parse, format)
		fs.Parse([]string{"-v", "foo:bar", "-v", "bar:baz"})
		require.Equal(t, &pair{"bar", "baz"}, v)
	})

	t.Run("generic var", func(t *testing.T) {
		fs := flag.NewFlagSet("", flag.PanicOnError)
		var v pair
		values.GenericVar(fs, &v, "v", "", parse, format)
		fs.Parse([]string{"-v", "foo:bar", "-v", "bar:baz"})
		require.Equal(t, pair{"bar", "baz"}, v)
	})

	t.Run("generics", func(t *testing.T) {
		fs := flag.NewFlagSet("", flag.PanicOnError)
		v := values.Generics(fs, "v", "", parse, format, values.Unsplit)
		fs.Parse([]string{"-v", "foo:bar", "-v", "bar:baz", "-v", "foo:bar,bar:baz"})
		require.Equal(t, &[]pair{{"foo", "bar"}, {"bar", "baz"}, {"foo", "bar,bar:baz"}}, v)
	})

	t.Run("generics var", func(t *testing.T) {
		fs := flag.NewFlagSet("", flag.PanicOnError)
		var v []pair
		values.GenericsVar(fs, &v, "v", "", parse, format, values.Unsplit)
		fs.Parse([]string{"-v", "foo:bar", "-v", "bar:baz", "-v", "foo:bar,bar:baz"})
		require.Equal(t, []pair{{"foo", "bar"}, {"bar", "baz"}, {"foo", "bar,bar:baz"}}, v)
	})

	t.Run("generics split", func(t *testing.T) {
		fs := flag.NewFlagSet("", flag.PanicOnError)
		v := values.Generics(fs, "v", "", parse, format, values.Split(","))
		fs.Parse([]string{"-v", "foo:bar", "-v", "bar:baz", "-v", "foo:bar,bar:baz"})
		require.Equal(t, &[]pair{{"foo", "bar"}, {"bar", "baz"}, {"foo", "bar"}, {"bar", "baz"}}, v)
	})

	t.Run("stringer", func(t *testing.T) {
		fs := flag.NewFlagSet("", flag.PanicOnError)
		v := values.Stringer(fs, "v", "", netip.ParseAddr)
		fs.Parse([]string{"-v", "1.2.3.4"})
		require.Equal(t, netip.AddrFrom4([4]byte{1, 2, 3, 4}), *v)
	})

	t.Run("stringer var", func(t *testing.T) {
		fs := flag.NewFlagSet("", flag.PanicOnError)
		var v netip.Addr
		values.StringerVar(fs, &v, "v", "", netip.ParseAddr)
		fs.Parse([]string{"-v", "1.2.3.4"})
		require.Equal(t, netip.AddrFrom4([4]byte{1, 2, 3, 4}), v)
	})

	t.Run("stringers", func(t *testing.T) {
		fs := flag.NewFlagSet("", flag.PanicOnError)
		v := values.Stringers(fs, "v", "", netip.ParseAddr, values.Unsplit)
		fs.Parse([]string{"-v", "1.2.3.4", "-v", "5.6.7.8"})
		require.Equal(t, []netip.Addr{netip.AddrFrom4([4]byte{1, 2, 3, 4}), netip.AddrFrom4([4]byte{5, 6, 7, 8})}, *v)
	})

	t.Run("stringers var", func(t *testing.T) {
		fs := flag.NewFlagSet("", flag.PanicOnError)
		var v []netip.Addr
		values.StringersVar(fs, &v, "v", "", netip.ParseAddr, values.Unsplit)
		fs.Parse([]string{"-v", "1.2.3.4", "-v", "5.6.7.8"})
		require.Equal(t, []netip.Addr{netip.AddrFrom4([4]byte{1, 2, 3, 4}), netip.AddrFrom4([4]byte{5, 6, 7, 8})}, v)
	})

	t.Run("time", func(t *testing.T) {
		fs := flag.NewFlagSet("", flag.PanicOnError)
		v := values.Time(fs, "v", "", time.RFC3339)
		fs.Parse([]string{"-v", "2025-05-07T06:06:06Z"})
		require.Equal(t, time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC), *v)
	})

	t.Run("time var", func(t *testing.T) {
		fs := flag.NewFlagSet("", flag.PanicOnError)
		var v time.Time
		values.TimeVar(fs, &v, "v", "", time.RFC3339)
		fs.Parse([]string{"-v", "2025-05-07T06:06:06Z"})
		require.Equal(t, time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC), v)
	})

	t.Run("times", func(t *testing.T) {
		fs := flag.NewFlagSet("", flag.PanicOnError)
		v := values.Times(fs, "v", "", time.RFC3339, values.Unsplit)
		fs.Parse([]string{"-v", "2025-05-07T06:06:06Z", "-v", "2025-05-07T09:09:09Z"})
		require.Equal(t, []time.Time{time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC), time.Date(2025, time.May, 7, 9, 9, 9, 0, time.UTC)}, *v)
	})

	t.Run("times var", func(t *testing.T) {
		fs := flag.NewFlagSet("", flag.PanicOnError)
		var v []time.Time
		values.TimesVar(fs, &v, "v", "", time.RFC3339, values.Unsplit)
		fs.Parse([]string{"-v", "2025-05-07T06:06:06Z", "-v", "2025-05-07T09:09:09Z"})
		require.Equal(t, []time.Time{time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC), time.Date(2025, time.May, 7, 9, 9, 9, 0, time.UTC)}, v)
	})
}
