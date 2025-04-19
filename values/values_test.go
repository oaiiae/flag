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

func TestGenericParse(t *testing.T) {
	type pair struct{ a, b string }
	parse := func(s string) (pair, error) { a, b, _ := strings.Cut(s, ":"); return pair{a, b}, nil }
	format := func(p pair) string { return p.a + ":" + p.b }

	t.Run("generic", func(t *testing.T) {
		v := values.Generic(parse, format)
		v.Set("foo:bar")
		v.Set("bar:baz")
		require.Equal(t, pair{"bar", "baz"}, v.(flag.Getter).Get())
	})

	t.Run("generic var", func(t *testing.T) {
		var p pair
		v := values.GenericVar(&p, parse, format)
		v.Set("foo:bar")
		v.Set("bar:baz")
		require.Equal(t, pair{"bar", "baz"}, v.(flag.Getter).Get())
		require.Equal(t, pair{"bar", "baz"}, p)
	})

	t.Run("generics", func(t *testing.T) {
		v := values.Generics(parse, format, values.Unsplit)
		v.Set("foo:bar")
		v.Set("bar:baz")
		v.Set("foo:bar,bar:baz")
		require.Equal(t, []pair{{"foo", "bar"}, {"bar", "baz"}, {"foo", "bar,bar:baz"}}, v.(flag.Getter).Get())
	})

	t.Run("generics var", func(t *testing.T) {
		var p []pair
		v := values.GenericsVar(&p, parse, format, values.Unsplit)
		v.Set("foo:bar")
		v.Set("bar:baz")
		v.Set("foo:bar,bar:baz")
		require.Equal(t, []pair{{"foo", "bar"}, {"bar", "baz"}, {"foo", "bar,bar:baz"}}, v.(flag.Getter).Get())
		require.Equal(t, []pair{{"foo", "bar"}, {"bar", "baz"}, {"foo", "bar,bar:baz"}}, p)
	})

	t.Run("generics split", func(t *testing.T) {
		var p []pair
		v := values.GenericsVar(&p, parse, format, values.Split(","))
		v.Set("foo:bar")
		v.Set("bar:baz")
		v.Set("foo:bar,bar:baz")
		require.Equal(t, []pair{{"foo", "bar"}, {"bar", "baz"}, {"foo", "bar"}, {"bar", "baz"}}, v.(flag.Getter).Get())
		require.Equal(t, []pair{{"foo", "bar"}, {"bar", "baz"}, {"foo", "bar"}, {"bar", "baz"}}, p)
	})
}

func TestOthersParse(t *testing.T) {
	t.Run("stringer", func(t *testing.T) {
		v := values.Stringer(netip.ParseAddr)
		v.Set("1.2.3.4")
		require.Equal(t, netip.AddrFrom4([4]byte{1, 2, 3, 4}), v.(flag.Getter).Get())
	})

	t.Run("stringer var", func(t *testing.T) {
		var p netip.Addr
		v := values.StringerVar(&p, netip.ParseAddr)
		v.Set("1.2.3.4")
		require.Equal(t, netip.AddrFrom4([4]byte{1, 2, 3, 4}), v.(flag.Getter).Get())
		require.Equal(t, netip.AddrFrom4([4]byte{1, 2, 3, 4}), p)
	})

	t.Run("stringers", func(t *testing.T) {
		v := values.Stringers(netip.ParseAddr, values.Unsplit)
		v.Set("1.2.3.4")
		v.Set("5.6.7.8")
		require.Equal(t, []netip.Addr{
			netip.AddrFrom4([4]byte{1, 2, 3, 4}),
			netip.AddrFrom4([4]byte{5, 6, 7, 8}),
		}, v.(flag.Getter).Get())
	})

	t.Run("stringers var", func(t *testing.T) {
		var p []netip.Addr
		v := values.StringersVar(&p, netip.ParseAddr, values.Unsplit)
		v.Set("1.2.3.4")
		v.Set("5.6.7.8")
		require.Equal(t, []netip.Addr{
			netip.AddrFrom4([4]byte{1, 2, 3, 4}),
			netip.AddrFrom4([4]byte{5, 6, 7, 8}),
		}, v.(flag.Getter).Get())
		require.Equal(t, []netip.Addr{
			netip.AddrFrom4([4]byte{1, 2, 3, 4}),
			netip.AddrFrom4([4]byte{5, 6, 7, 8}),
		}, p)
	})

	t.Run("time", func(t *testing.T) {
		v := values.Time(time.RFC3339)
		v.Set("2025-05-07T06:06:06Z")
		require.Equal(t, time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC), v.(flag.Getter).Get())
	})

	t.Run("time var", func(t *testing.T) {
		var p time.Time
		v := values.TimeVar(&p, time.RFC3339)
		v.Set("2025-05-07T06:06:06Z")
		require.Equal(t, time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC), v.(flag.Getter).Get())
		require.Equal(t, time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC), p)
	})

	t.Run("times", func(t *testing.T) {
		v := values.Times(time.RFC3339, values.Unsplit)
		v.Set("2025-05-07T06:06:06Z")
		v.Set("2025-05-07T09:09:09Z")
		require.Equal(t, []time.Time{
			time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC),
			time.Date(2025, time.May, 7, 9, 9, 9, 0, time.UTC),
		}, v.(flag.Getter).Get())
	})

	t.Run("times var", func(t *testing.T) {
		var p []time.Time
		v := values.TimesVar(&p, time.RFC3339, values.Unsplit)
		v.Set("2025-05-07T06:06:06Z")
		v.Set("2025-05-07T09:09:09Z")
		require.Equal(t, []time.Time{
			time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC),
			time.Date(2025, time.May, 7, 9, 9, 9, 0, time.UTC),
		}, v.(flag.Getter).Get())
		require.Equal(t, []time.Time{
			time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC),
			time.Date(2025, time.May, 7, 9, 9, 9, 0, time.UTC),
		}, p)
	})
}
