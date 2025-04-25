package values_test

import (
	"flag"
	"net/mail"
	"net/netip"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"gitlab.com/oaiiae/flag/values"
)

func TestRegisterer_values(t *testing.T) {
	testCases := []struct {
		name     string
		setup    func(values.Registerer)
		defValue string
		isType   any
		input    string
		output   any
	}{
		{
			name:     "bool",
			setup:    func(r values.Registerer) { r.Bool("f", false, "usg") },
			defValue: "false",
			isType:   values.Basic[bool](),
			input:    "true",
			output:   true,
		},
		{
			name:     "complex64",
			setup:    func(r values.Registerer) { r.Complex64("f", 0, "usg") },
			defValue: "(0+0i)",
			isType:   values.Basic[complex64](),
			input:    "12+42i",
			output:   complex64(12 + 42i),
		},
		{
			name:     "complex128",
			setup:    func(r values.Registerer) { r.Complex128("f", 0, "usg") },
			defValue: "(0+0i)",
			isType:   values.Basic[complex128](),
			input:    "15+62i",
			output:   complex128(15 + 62i),
		},
		{
			name:     "int",
			setup:    func(r values.Registerer) { r.Int("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[int](),
			input:    "-42",
			output:   int(-42),
		},
		{
			name:     "int8",
			setup:    func(r values.Registerer) { r.Int8("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[int8](),
			input:    "-42",
			output:   int8(-42),
		},
		{
			name:     "int16",
			setup:    func(r values.Registerer) { r.Int16("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[int16](),
			input:    "-42",
			output:   int16(-42),
		},
		{
			name:     "int32",
			setup:    func(r values.Registerer) { r.Int32("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[int32](),
			input:    "-42",
			output:   int32(-42),
		},
		{
			name:     "int64",
			setup:    func(r values.Registerer) { r.Int64("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[int64](),
			input:    "-42",
			output:   int64(-42),
		},
		{
			name:     "uint",
			setup:    func(r values.Registerer) { r.Uint("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[uint](),
			input:    "42",
			output:   uint(42),
		},
		{
			name:     "uint8",
			setup:    func(r values.Registerer) { r.Uint8("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[uint8](),
			input:    "42",
			output:   uint8(42),
		},
		{
			name:     "uint16",
			setup:    func(r values.Registerer) { r.Uint16("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[uint16](),
			input:    "42",
			output:   uint16(42),
		},
		{
			name:     "uint32",
			setup:    func(r values.Registerer) { r.Uint32("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[uint32](),
			input:    "42",
			output:   uint32(42),
		},
		{
			name:     "uint64",
			setup:    func(r values.Registerer) { r.Uint64("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[uint64](),
			input:    "42",
			output:   uint64(42),
		},
		{
			name:     "float32",
			setup:    func(r values.Registerer) { r.Float32("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[float32](),
			input:    "3.14",
			output:   float32(3.14),
		},
		{
			name:     "float64",
			setup:    func(r values.Registerer) { r.Float64("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[float64](),
			input:    "3.14159",
			output:   float64(3.14159),
		},
		{
			name:     "string",
			setup:    func(r values.Registerer) { r.String("f", "", "usg") },
			defValue: "",
			isType:   values.Basic[string](),
			input:    "hello world",
			output:   "hello world",
		},
		{
			name:     "duration",
			setup:    func(r values.Registerer) { r.Duration("f", 5*time.Minute, "usg") },
			defValue: "5m0s",
			isType:   values.DurationVar(nil),
			input:    "3h30m",
			output:   3*time.Hour + 30*time.Minute,
		},
		{
			name:     "time",
			setup:    func(r values.Registerer) { r.Time("f", time.Time{}, time.RFC3339, "usg") },
			defValue: "0001-01-01T00:00:00Z",
			isType:   values.TimeVar(nil, ""),
			input:    "2025-05-07T06:06:06Z",
			output:   time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC),
		},
		{
			name:     "ip address",
			setup:    func(r values.Registerer) { r.IPAddr("f", netip.Addr{}, "usg") },
			defValue: "invalid IP",
			isType:   values.StringerVar[netip.Addr](nil, nil),
			input:    "192.168.1.1",
			output:   netip.MustParseAddr("192.168.1.1"),
		},
		{
			name:     "ip address and port",
			setup:    func(r values.Registerer) { r.IPAddrPort("f", netip.AddrPort{}, "usg") },
			defValue: "invalid AddrPort",
			isType:   values.StringerVar[netip.AddrPort](nil, nil),
			input:    "192.168.1.1:8080",
			output:   netip.MustParseAddrPort("192.168.1.1:8080"),
		},
		{
			name:     "ip prefix",
			setup:    func(r values.Registerer) { r.IPPrefix("f", netip.Prefix{}, "usg") },
			defValue: "invalid Prefix",
			isType:   values.StringerVar[netip.Prefix](nil, nil),
			input:    "192.168.1.0/24",
			output:   netip.MustParsePrefix("192.168.1.0/24"),
		},
		{
			name:     "mail addr",
			setup:    func(r values.Registerer) { r.MailAddr("f", &mail.Address{}, "usg") },
			defValue: "<@>",
			isType:   values.StringerVar[*mail.Address](nil, nil),
			input:    "foo@bar.com",
			output:   &mail.Address{Address: "foo@bar.com"},
		},
		{
			name:     "url",
			setup:    func(r values.Registerer) { r.URL("f", &url.URL{}, "usg") },
			defValue: "",
			isType:   values.StringerVar[*url.URL](nil, nil),
			input:    "foo://bar",
			output:   &url.URL{Scheme: "foo", Host: "bar"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fs := flag.NewFlagSet("", flag.ContinueOnError)
			tc.setup(values.FlagSetRegisterer(fs))

			f := fs.Lookup("f")
			require.NotNil(t, f)
			require.Equal(t, "f", f.Name)
			require.Equal(t, "usg", f.Usage)
			require.Equal(t, tc.defValue, f.DefValue)
			require.IsType(t, tc.isType, f.Value)
			require.NoError(t, fs.Parse([]string{"-f", tc.input}))
			require.Equal(t, tc.output, f.Value.(flag.Getter).Get())
		})
	}
}
