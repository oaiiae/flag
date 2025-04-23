package values_test

import (
	"flag"
	"fmt"
	"net/netip"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"gitlab.com/oaiiae/flag/values"
)

func TestFlagValues(t *testing.T) {
	type pair struct{ a, b string }
	parse := func(s string) (pair, error) { a, b, _ := strings.Cut(s, ":"); return pair{a, b}, nil }
	format := func(p pair) string { return p.a + ":" + p.b }

	t.Run("generic", func(t *testing.T) {
		v := values.Generic(parse, format)
		require.NoError(t, v.Set("foo:bar"))
		require.NoError(t, v.Set("bar:baz"))
		require.Equal(t, "bar:baz", v.String())
		require.Equal(t, pair{"bar", "baz"}, v.(flag.Getter).Get())
	})

	t.Run("generic var", func(t *testing.T) {
		var p pair
		v := values.GenericVar(&p, parse, format)
		require.NoError(t, v.Set("foo:bar"))
		require.NoError(t, v.Set("bar:baz"))
		require.Equal(t, "bar:baz", v.String())
		require.Equal(t, pair{"bar", "baz"}, v.(flag.Getter).Get())
		require.Equal(t, pair{"bar", "baz"}, p)
	})

	t.Run("generic slice", func(t *testing.T) {
		v := values.GenericSlice(",", parse, format)
		require.NoError(t, v.Set("quu:quux"))
		require.NoError(t, v.Set("foo:bar,bar:baz"))
		require.Equal(t, "foo:bar,bar:baz", v.String())
		require.Equal(t, []pair{{"foo", "bar"}, {"bar", "baz"}}, v.(flag.Getter).Get())
	})

	t.Run("generic slice var", func(t *testing.T) {
		var p []pair
		v := values.GenericSliceVar(&p, ",", parse, format)
		require.NoError(t, v.Set("quu:quux"))
		require.NoError(t, v.Set("foo:bar,bar:baz"))
		require.Equal(t, "foo:bar,bar:baz", v.String())
		require.Equal(t, []pair{{"foo", "bar"}, {"bar", "baz"}}, v.(flag.Getter).Get())
		require.Equal(t, []pair{{"foo", "bar"}, {"bar", "baz"}}, p)
	})

	t.Run("generic list", func(t *testing.T) {
		v := values.GenericList(parse, format)
		require.NoError(t, v.Set("foo:bar"))
		require.NoError(t, v.Set("bar:baz"))
		require.Equal(t, "[foo:bar bar:baz]", v.String())
		require.Equal(t, []pair{{"foo", "bar"}, {"bar", "baz"}}, v.(flag.Getter).Get())
	})

	t.Run("generic list var", func(t *testing.T) {
		var p []pair
		v := values.GenericListVar(&p, parse, format)
		require.NoError(t, v.Set("foo:bar"))
		require.NoError(t, v.Set("bar:baz"))
		require.Equal(t, "[foo:bar bar:baz]", v.String())
		require.Equal(t, []pair{{"foo", "bar"}, {"bar", "baz"}}, v.(flag.Getter).Get())
		require.Equal(t, []pair{{"foo", "bar"}, {"bar", "baz"}}, p)
	})

	t.Run("basic", func(t *testing.T) {
		v := values.Basic[complex64]()
		require.NoError(t, v.Set("3+4i"))
		require.Equal(t, "(3+4i)", v.String())
		require.Equal(t, complex64(3+4i), v.(flag.Getter).Get())
	})

	t.Run("basic var", func(t *testing.T) {
		var p complex64
		v := values.BasicVar(&p)
		require.NoError(t, v.Set("3+4i"))
		require.Equal(t, "(3+4i)", v.String())
		require.Equal(t, complex64(3+4i), v.(flag.Getter).Get())
		require.Equal(t, complex64(3+4i), p)
	})

	t.Run("basic slice", func(t *testing.T) {
		v := values.BasicSlice[complex64](",")
		require.NoError(t, v.Set("3+4i,5+6i"))
		require.Equal(t, "(3+4i),(5+6i)", v.String())
		require.Equal(t, []complex64{3 + 4i, 5 + 6i}, v.(flag.Getter).Get())
	})

	t.Run("basic slice var", func(t *testing.T) {
		var p []complex64
		v := values.BasicSliceVar(&p, ",")
		require.NoError(t, v.Set("3+4i,5+6i"))
		require.Equal(t, "(3+4i),(5+6i)", v.String())
		require.Equal(t, []complex64{3 + 4i, 5 + 6i}, v.(flag.Getter).Get())
		require.Equal(t, []complex64{3 + 4i, 5 + 6i}, p)
	})

	t.Run("basic list", func(t *testing.T) {
		v := values.BasicList[complex64]()
		require.NoError(t, v.Set("3+4i"))
		require.NoError(t, v.Set("5+6i"))
		require.Equal(t, "[(3+4i) (5+6i)]", v.String())
		require.Equal(t, []complex64{3 + 4i, 5 + 6i}, v.(flag.Getter).Get())
	})

	t.Run("basic list var", func(t *testing.T) {
		var p []complex64
		v := values.BasicListVar(&p)
		require.NoError(t, v.Set("3+4i"))
		require.NoError(t, v.Set("5+6i"))
		require.Equal(t, "[(3+4i) (5+6i)]", v.String())
		require.Equal(t, []complex64{3 + 4i, 5 + 6i}, v.(flag.Getter).Get())
		require.Equal(t, []complex64{3 + 4i, 5 + 6i}, p)
	})

	t.Run("stringer", func(t *testing.T) {
		v := values.Stringer(netip.ParseAddr)
		require.NoError(t, v.Set("1.2.3.4"))
		require.Equal(t, "1.2.3.4", v.String())
		require.Equal(t, netip.AddrFrom4([4]byte{1, 2, 3, 4}), v.(flag.Getter).Get())
	})

	t.Run("stringer var", func(t *testing.T) {
		var p netip.Addr
		v := values.StringerVar(&p, netip.ParseAddr)
		require.NoError(t, v.Set("1.2.3.4"))
		require.Equal(t, "1.2.3.4", v.String())
		require.Equal(t, netip.AddrFrom4([4]byte{1, 2, 3, 4}), v.(flag.Getter).Get())
		require.Equal(t, netip.AddrFrom4([4]byte{1, 2, 3, 4}), p)
	})

	t.Run("stringer slice", func(t *testing.T) {
		v := values.StringerSlice(",", netip.ParseAddr)
		require.NoError(t, v.Set("1.2.3.4,5.6.7.8"))
		require.Equal(t, "1.2.3.4,5.6.7.8", v.String())
		require.Equal(t, []netip.Addr{
			netip.AddrFrom4([4]byte{1, 2, 3, 4}),
			netip.AddrFrom4([4]byte{5, 6, 7, 8}),
		}, v.(flag.Getter).Get())
	})

	t.Run("stringer slice var", func(t *testing.T) {
		var p []netip.Addr
		v := values.StringerSliceVar(&p, ",", netip.ParseAddr)
		require.NoError(t, v.Set("1.2.3.4,5.6.7.8"))
		require.Equal(t, "1.2.3.4,5.6.7.8", v.String())
		require.Equal(t, []netip.Addr{
			netip.AddrFrom4([4]byte{1, 2, 3, 4}),
			netip.AddrFrom4([4]byte{5, 6, 7, 8}),
		}, v.(flag.Getter).Get())
		require.Equal(t, []netip.Addr{
			netip.AddrFrom4([4]byte{1, 2, 3, 4}),
			netip.AddrFrom4([4]byte{5, 6, 7, 8}),
		}, p)
	})

	t.Run("stringer list", func(t *testing.T) {
		v := values.StringerList(netip.ParseAddr)
		require.NoError(t, v.Set("1.2.3.4"))
		require.NoError(t, v.Set("5.6.7.8"))
		require.Equal(t, "[1.2.3.4 5.6.7.8]", v.String())
		require.Equal(t, []netip.Addr{
			netip.AddrFrom4([4]byte{1, 2, 3, 4}),
			netip.AddrFrom4([4]byte{5, 6, 7, 8}),
		}, v.(flag.Getter).Get())
	})

	t.Run("stringer list var", func(t *testing.T) {
		var p []netip.Addr
		v := values.StringerListVar(&p, netip.ParseAddr)
		require.NoError(t, v.Set("1.2.3.4"))
		require.NoError(t, v.Set("5.6.7.8"))
		require.Equal(t, "[1.2.3.4 5.6.7.8]", v.String())
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
		require.NoError(t, v.Set("2025-05-07T06:06:06Z"))
		require.Equal(t, "2025-05-07T06:06:06Z", v.String())
		require.Equal(t, time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC), v.(flag.Getter).Get())
	})

	t.Run("time var", func(t *testing.T) {
		var p time.Time
		v := values.TimeVar(&p, time.RFC3339)
		require.NoError(t, v.Set("2025-05-07T06:06:06Z"))
		require.Equal(t, "2025-05-07T06:06:06Z", v.String())
		require.Equal(t, time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC), v.(flag.Getter).Get())
		require.Equal(t, time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC), p)
	})

	t.Run("time slice", func(t *testing.T) {
		v := values.TimeSlice(",", time.RFC3339)
		require.NoError(t, v.Set("2025-05-07T06:06:06Z,2025-05-07T09:09:09Z"))
		require.Equal(t, "2025-05-07T06:06:06Z,2025-05-07T09:09:09Z", v.String())
		require.Equal(t, []time.Time{
			time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC),
			time.Date(2025, time.May, 7, 9, 9, 9, 0, time.UTC),
		}, v.(flag.Getter).Get())
	})

	t.Run("time slice var", func(t *testing.T) {
		var p []time.Time
		v := values.TimeSliceVar(&p, ",", time.RFC3339)
		require.NoError(t, v.Set("2025-05-07T06:06:06Z,2025-05-07T09:09:09Z"))
		require.Equal(t, "2025-05-07T06:06:06Z,2025-05-07T09:09:09Z", v.String())
		require.Equal(t, []time.Time{
			time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC),
			time.Date(2025, time.May, 7, 9, 9, 9, 0, time.UTC),
		}, v.(flag.Getter).Get())
		require.Equal(t, []time.Time{
			time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC),
			time.Date(2025, time.May, 7, 9, 9, 9, 0, time.UTC),
		}, p)
	})

	t.Run("time list", func(t *testing.T) {
		v := values.TimeList(time.RFC3339)
		require.NoError(t, v.Set("2025-05-07T06:06:06Z"))
		require.NoError(t, v.Set("2025-05-07T09:09:09Z"))
		require.Equal(t, "[2025-05-07T06:06:06Z 2025-05-07T09:09:09Z]", v.String())
		require.Equal(t, []time.Time{
			time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC),
			time.Date(2025, time.May, 7, 9, 9, 9, 0, time.UTC),
		}, v.(flag.Getter).Get())
	})

	t.Run("time list var", func(t *testing.T) {
		var p []time.Time
		v := values.TimeListVar(&p, time.RFC3339)
		require.NoError(t, v.Set("2025-05-07T06:06:06Z"))
		require.NoError(t, v.Set("2025-05-07T09:09:09Z"))
		require.Equal(t, "[2025-05-07T06:06:06Z 2025-05-07T09:09:09Z]", v.String())
		require.Equal(t, []time.Time{
			time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC),
			time.Date(2025, time.May, 7, 9, 9, 9, 0, time.UTC),
		}, v.(flag.Getter).Get())
		require.Equal(t, []time.Time{
			time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC),
			time.Date(2025, time.May, 7, 9, 9, 9, 0, time.UTC),
		}, p)
	})

	t.Run("duration", func(t *testing.T) {
		v := values.Duration()
		require.NoError(t, v.Set("5h30m"))
		require.Equal(t, "5h30m0s", v.String())
		require.Equal(t, 5*time.Hour+30*time.Minute, v.(flag.Getter).Get())
	})

	t.Run("duration var", func(t *testing.T) {
		var p time.Duration
		v := values.DurationVar(&p)
		require.NoError(t, v.Set("5h30m"))
		require.Equal(t, "5h30m0s", v.String())
		require.Equal(t, 5*time.Hour+30*time.Minute, v.(flag.Getter).Get())
		require.Equal(t, 5*time.Hour+30*time.Minute, p)
	})

	t.Run("duration slice", func(t *testing.T) {
		v := values.DurationSlice(",")
		require.NoError(t, v.Set("5h30m,1h15m"))
		require.Equal(t, "5h30m0s,1h15m0s", v.String())
		require.Equal(t, []time.Duration{330 * time.Minute, 75 * time.Minute}, v.(flag.Getter).Get())
	})

	t.Run("duration slice var", func(t *testing.T) {
		var p []time.Duration
		v := values.DurationSliceVar(&p, ",")
		require.NoError(t, v.Set("5h30m,1h15m"))
		require.Equal(t, "5h30m0s,1h15m0s", v.String())
		require.Equal(t, []time.Duration{330 * time.Minute, 75 * time.Minute}, v.(flag.Getter).Get())
		require.Equal(t, []time.Duration{330 * time.Minute, 75 * time.Minute}, p)
	})

	t.Run("duration list", func(t *testing.T) {
		v := values.DurationList()
		require.NoError(t, v.Set("5h30m"))
		require.NoError(t, v.Set("1h15m"))
		require.Equal(t, "[5h30m0s 1h15m0s]", v.String())
		require.Equal(t, []time.Duration{330 * time.Minute, 75 * time.Minute}, v.(flag.Getter).Get())
	})

	t.Run("duration list var", func(t *testing.T) {
		var p []time.Duration
		v := values.DurationListVar(&p)
		require.NoError(t, v.Set("5h30m"))
		require.NoError(t, v.Set("1h15m"))
		require.Equal(t, "[5h30m0s 1h15m0s]", v.String())
		require.Equal(t, []time.Duration{330 * time.Minute, 75 * time.Minute}, v.(flag.Getter).Get())
		require.Equal(t, []time.Duration{330 * time.Minute, 75 * time.Minute}, p)
	})
}

func TestBasicParsing(t *testing.T) {
	testCases := []struct {
		value     flag.Value
		input     string
		expectVal any
		expectStr string
	}{
		{
			value:     values.Basic[bool](),
			input:     "true",
			expectVal: true,
			expectStr: "true",
		},
		{
			value:     values.Basic[complex64](),
			input:     "3+4i",
			expectVal: complex64(3 + 4i),
			expectStr: "(3+4i)",
		},
		{
			value:     values.Basic[complex128](),
			input:     "5+6i",
			expectVal: complex128(5 + 6i),
			expectStr: "(5+6i)",
		},
		{
			value:     values.Basic[int](),
			input:     "-42",
			expectVal: int(-42),
			expectStr: "-42",
		},
		{
			value:     values.Basic[int8](),
			input:     "-42",
			expectVal: int8(-42),
			expectStr: "-42",
		},
		{
			value:     values.Basic[int16](),
			input:     "-42",
			expectVal: int16(-42),
			expectStr: "-42",
		},
		{
			value:     values.Basic[int32](),
			input:     "-42",
			expectVal: int32(-42),
			expectStr: "-42",
		},
		{
			value:     values.Basic[int64](),
			input:     "-42",
			expectVal: int64(-42),
			expectStr: "-42",
		},
		{
			value:     values.Basic[uint](),
			input:     "42",
			expectVal: uint(42),
			expectStr: "42",
		},
		{
			value:     values.Basic[uint8](),
			input:     "42",
			expectVal: uint8(42),
			expectStr: "42",
		},
		{
			value:     values.Basic[uint16](),
			input:     "42",
			expectVal: uint16(42),
			expectStr: "42",
		},
		{
			value:     values.Basic[uint32](),
			input:     "42",
			expectVal: uint32(42),
			expectStr: "42",
		},
		{
			value:     values.Basic[uint64](),
			input:     "42",
			expectVal: uint64(42),
			expectStr: "42",
		},
		{
			value:     values.Basic[float32](),
			input:     "3.14",
			expectVal: float32(3.14),
			expectStr: "3.14",
		},
		{
			value:     values.Basic[float64](),
			input:     "3.14159",
			expectVal: float64(3.14159),
			expectStr: "3.14159",
		},
		{
			value:     values.Basic[string](),
			input:     "hello world",
			expectVal: "hello world",
			expectStr: "hello world",
		},
		// FIXME: check fails because of []byte and reflect.DeepEqual?
		// {
		// 	value:     values.Basic[[]byte](),
		// 	input:     "hello world",
		// 	expectVal: []byte("hello world"),
		// 	expectStr: "hello world",
		// },
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%T", tc.value), func(t *testing.T) {
			require.NoError(t, tc.value.Set(tc.input))
			require.Equal(t, tc.expectStr, tc.value.String())
			require.Equal(t, tc.expectVal, tc.value.(flag.Getter).Get())
		})
	}
}
