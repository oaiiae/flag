package values_test

import (
	"flag"
	"fmt"
	"net/mail"
	"net/netip"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/rlibaert/flag/values"
)

func ExampleFlagSetEnvRegisterer() {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	os.Setenv("FOO_INT", "42")
	values.FlagSetEnvRegisterer(fs, "FOO_").Int("int", 12, "an int")
	fmt.Println(fs.Lookup("int").DefValue, fs.Lookup("int").Value)
	fs.Usage()

	// Output:
	// 12 42
	// Usage:
	//   -int value
	//     	an int (env $FOO_INT) (default 12)
}

func TestRegisterer_values(t *testing.T) {
	testCases := []struct {
		name     string
		setup    func(values.RegistererFunc)
		defValue string
		isType   any
		input    string
		output   any
	}{
		{
			name:     "bool",
			setup:    func(r values.RegistererFunc) { r.Bool("f", false, "usg") },
			defValue: "false",
			isType:   values.Basic[bool](),
			input:    "true",
			output:   true,
		},
		{
			name:     "bool list",
			setup:    func(r values.RegistererFunc) { r.BoolList("f", nil, "usg") },
			defValue: "",
			isType:   values.BasicList[bool](),
			input:    "true",
			output:   []bool{true},
		},
		{
			name:     "bool slice",
			setup:    func(r values.RegistererFunc) { r.BoolSlice("f", nil, ",", "usg") },
			defValue: "",
			isType:   values.BasicSlice[bool](""),
			input:    "true,false,true",
			output:   []bool{true, false, true},
		},
		{
			name:     "complex64",
			setup:    func(r values.RegistererFunc) { r.Complex64("f", 0, "usg") },
			defValue: "(0+0i)",
			isType:   values.Basic[complex64](),
			input:    "12+42i",
			output:   complex64(12 + 42i),
		},
		{
			name:     "complex64 list",
			setup:    func(r values.RegistererFunc) { r.Complex64List("f", nil, "usg") },
			defValue: "",
			isType:   values.BasicList[complex64](),
			input:    "12+42i",
			output:   []complex64{complex64(12 + 42i)},
		},
		{
			name:     "complex64 slice",
			setup:    func(r values.RegistererFunc) { r.Complex64Slice("f", nil, ",", "usg") },
			defValue: "",
			isType:   values.BasicSlice[complex64](""),
			input:    "12+42i,13+43i,14+44i",
			output:   []complex64{complex64(12 + 42i), complex64(13 + 43i), complex64(14 + 44i)},
		},
		{
			name:     "complex128",
			setup:    func(r values.RegistererFunc) { r.Complex128("f", 0, "usg") },
			defValue: "(0+0i)",
			isType:   values.Basic[complex128](),
			input:    "15+62i",
			output:   complex128(15 + 62i),
		},
		{
			name:     "complex128 list",
			setup:    func(r values.RegistererFunc) { r.Complex128List("f", nil, "usg") },
			defValue: "",
			isType:   values.BasicList[complex128](),
			input:    "15+62i",
			output:   []complex128{complex128(15 + 62i)},
		},
		{
			name:     "complex128 slice",
			setup:    func(r values.RegistererFunc) { r.Complex128Slice("f", nil, ",", "usg") },
			defValue: "",
			isType:   values.BasicSlice[complex128](""),
			input:    "15+62i,16+63i,17+64i",
			output:   []complex128{complex128(15 + 62i), complex128(16 + 63i), complex128(17 + 64i)},
		},
		{
			name:     "int",
			setup:    func(r values.RegistererFunc) { r.Int("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[int](),
			input:    "-42",
			output:   int(-42),
		},
		{
			name:     "int list",
			setup:    func(r values.RegistererFunc) { r.IntList("f", nil, "usg") },
			defValue: "",
			isType:   values.BasicList[int](),
			input:    "-42",
			output:   []int{int(-42)},
		},
		{
			name:     "int slice",
			setup:    func(r values.RegistererFunc) { r.IntSlice("f", nil, ",", "usg") },
			defValue: "",
			isType:   values.BasicSlice[int](""),
			input:    "-42,-43,-44",
			output:   []int{int(-42), int(-43), int(-44)},
		},
		{
			name:     "int8",
			setup:    func(r values.RegistererFunc) { r.Int8("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[int8](),
			input:    "-42",
			output:   int8(-42),
		},
		{
			name:     "int8 list",
			setup:    func(r values.RegistererFunc) { r.Int8List("f", nil, "usg") },
			defValue: "",
			isType:   values.BasicList[int8](),
			input:    "-42",
			output:   []int8{int8(-42)},
		},
		{
			name:     "int8 slice",
			setup:    func(r values.RegistererFunc) { r.Int8Slice("f", nil, ",", "usg") },
			defValue: "",
			isType:   values.BasicSlice[int8](""),
			input:    "-42,-43,-44",
			output:   []int8{int8(-42), int8(-43), int8(-44)},
		},
		{
			name:     "int16",
			setup:    func(r values.RegistererFunc) { r.Int16("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[int16](),
			input:    "-42",
			output:   int16(-42),
		},
		{
			name:     "int16 list",
			setup:    func(r values.RegistererFunc) { r.Int16List("f", nil, "usg") },
			defValue: "",
			isType:   values.BasicList[int16](),
			input:    "-42",
			output:   []int16{int16(-42)},
		},
		{
			name:     "int16 slice",
			setup:    func(r values.RegistererFunc) { r.Int16Slice("f", nil, ",", "usg") },
			defValue: "",
			isType:   values.BasicSlice[int16](""),
			input:    "-42,-43,-44",
			output:   []int16{int16(-42), int16(-43), int16(-44)},
		},
		{
			name:     "int32",
			setup:    func(r values.RegistererFunc) { r.Int32("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[int32](),
			input:    "-42",
			output:   int32(-42),
		},
		{
			name:     "int32 list",
			setup:    func(r values.RegistererFunc) { r.Int32List("f", nil, "usg") },
			defValue: "",
			isType:   values.BasicList[int32](),
			input:    "-42",
			output:   []int32{int32(-42)},
		},
		{
			name:     "int32 slice",
			setup:    func(r values.RegistererFunc) { r.Int32Slice("f", nil, ",", "usg") },
			defValue: "",
			isType:   values.BasicSlice[int32](""),
			input:    "-42,-43,-44",
			output:   []int32{int32(-42), int32(-43), int32(-44)},
		},
		{
			name:     "int64",
			setup:    func(r values.RegistererFunc) { r.Int64("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[int64](),
			input:    "-42",
			output:   int64(-42),
		},
		{
			name:     "int64 list",
			setup:    func(r values.RegistererFunc) { r.Int64List("f", nil, "usg") },
			defValue: "",
			isType:   values.BasicList[int64](),
			input:    "-42",
			output:   []int64{int64(-42)},
		},
		{
			name:     "int64 slice",
			setup:    func(r values.RegistererFunc) { r.Int64Slice("f", nil, ",", "usg") },
			defValue: "",
			isType:   values.BasicSlice[int64](""),
			input:    "-42,-43,-44",
			output:   []int64{int64(-42), int64(-43), int64(-44)},
		},
		{
			name:     "uint",
			setup:    func(r values.RegistererFunc) { r.Uint("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[uint](),
			input:    "42",
			output:   uint(42),
		},
		{
			name:     "uint list",
			setup:    func(r values.RegistererFunc) { r.UintList("f", nil, "usg") },
			defValue: "",
			isType:   values.BasicList[uint](),
			input:    "42",
			output:   []uint{uint(42)},
		},
		{
			name:     "uint slice",
			setup:    func(r values.RegistererFunc) { r.UintSlice("f", nil, ",", "usg") },
			defValue: "",
			isType:   values.BasicSlice[uint](""),
			input:    "42,43,44",
			output:   []uint{uint(42), uint(43), uint(44)},
		},
		{
			name:     "uint8",
			setup:    func(r values.RegistererFunc) { r.Uint8("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[uint8](),
			input:    "42",
			output:   uint8(42),
		},
		{
			name:     "uint8 list",
			setup:    func(r values.RegistererFunc) { r.Uint8List("f", nil, "usg") },
			defValue: "",
			isType:   values.BasicList[uint8](),
			input:    "42",
			output:   []uint8{uint8(42)},
		},
		{
			name:     "uint8 slice",
			setup:    func(r values.RegistererFunc) { r.Uint8Slice("f", nil, ",", "usg") },
			defValue: "",
			isType:   values.BasicSlice[uint8](""),
			input:    "42,43,44",
			output:   []uint8{uint8(42), uint8(43), uint8(44)},
		},
		{
			name:     "uint16",
			setup:    func(r values.RegistererFunc) { r.Uint16("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[uint16](),
			input:    "42",
			output:   uint16(42),
		},
		{
			name:     "uint16 list",
			setup:    func(r values.RegistererFunc) { r.Uint16List("f", nil, "usg") },
			defValue: "",
			isType:   values.BasicList[uint16](),
			input:    "42",
			output:   []uint16{uint16(42)},
		},
		{
			name:     "uint16 slice",
			setup:    func(r values.RegistererFunc) { r.Uint16Slice("f", nil, ",", "usg") },
			defValue: "",
			isType:   values.BasicSlice[uint16](""),
			input:    "42,43,44",
			output:   []uint16{uint16(42), uint16(43), uint16(44)},
		},
		{
			name:     "uint32",
			setup:    func(r values.RegistererFunc) { r.Uint32("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[uint32](),
			input:    "42",
			output:   uint32(42),
		},
		{
			name:     "uint32 list",
			setup:    func(r values.RegistererFunc) { r.Uint32List("f", nil, "usg") },
			defValue: "",
			isType:   values.BasicList[uint32](),
			input:    "42",
			output:   []uint32{uint32(42)},
		},
		{
			name:     "uint32 slice",
			setup:    func(r values.RegistererFunc) { r.Uint32Slice("f", nil, ",", "usg") },
			defValue: "",
			isType:   values.BasicSlice[uint32](""),
			input:    "42,43,44",
			output:   []uint32{uint32(42), uint32(43), uint32(44)},
		},
		{
			name:     "uint64",
			setup:    func(r values.RegistererFunc) { r.Uint64("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[uint64](),
			input:    "42",
			output:   uint64(42),
		},
		{
			name:     "uint64 list",
			setup:    func(r values.RegistererFunc) { r.Uint64List("f", nil, "usg") },
			defValue: "",
			isType:   values.BasicList[uint64](),
			input:    "42",
			output:   []uint64{uint64(42)},
		},
		{
			name:     "uint64 slice",
			setup:    func(r values.RegistererFunc) { r.Uint64Slice("f", nil, ",", "usg") },
			defValue: "",
			isType:   values.BasicSlice[uint64](""),
			input:    "42,43,44",
			output:   []uint64{uint64(42), uint64(43), uint64(44)},
		},
		{
			name:     "float32",
			setup:    func(r values.RegistererFunc) { r.Float32("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[float32](),
			input:    "3.14",
			output:   float32(3.14),
		},
		{
			name:     "float32 list",
			setup:    func(r values.RegistererFunc) { r.Float32List("f", nil, "usg") },
			defValue: "",
			isType:   values.BasicList[float32](),
			input:    "3.14",
			output:   []float32{float32(3.14)},
		},
		{
			name:     "float32 slice",
			setup:    func(r values.RegistererFunc) { r.Float32Slice("f", nil, ",", "usg") },
			defValue: "",
			isType:   values.BasicSlice[float32](""),
			input:    "3.14,2.71,1.41",
			output:   []float32{float32(3.14), float32(2.71), float32(1.41)},
		},
		{
			name:     "float64",
			setup:    func(r values.RegistererFunc) { r.Float64("f", 0, "usg") },
			defValue: "0",
			isType:   values.Basic[float64](),
			input:    "3.14159",
			output:   float64(3.14159),
		},
		{
			name:     "float64 list",
			setup:    func(r values.RegistererFunc) { r.Float64List("f", nil, "usg") },
			defValue: "",
			isType:   values.BasicList[float64](),
			input:    "3.14159",
			output:   []float64{float64(3.14159)},
		},
		{
			name:     "float64 slice",
			setup:    func(r values.RegistererFunc) { r.Float64Slice("f", nil, ",", "usg") },
			defValue: "",
			isType:   values.BasicSlice[float64](""),
			input:    "3.14159,2.71828,1.41421",
			output:   []float64{float64(3.14159), float64(2.71828), float64(1.41421)},
		},
		{
			name:     "string",
			setup:    func(r values.RegistererFunc) { r.String("f", "", "usg") },
			defValue: "",
			isType:   values.Basic[string](),
			input:    "hello world",
			output:   "hello world",
		},
		{
			name:     "string list",
			setup:    func(r values.RegistererFunc) { r.StringList("f", nil, "usg") },
			defValue: "",
			isType:   values.BasicList[string](),
			input:    "hello world",
			output:   []string{"hello world"},
		},
		{
			name:     "string slice",
			setup:    func(r values.RegistererFunc) { r.StringSlice("f", nil, ",", "usg") },
			defValue: "",
			isType:   values.BasicSlice[string](""),
			input:    "hello,world,foo",
			output:   []string{"hello", "world", "foo"},
		},
		{
			name:     "duration",
			setup:    func(r values.RegistererFunc) { r.Duration("f", 5*time.Minute, "usg") },
			defValue: "5m0s",
			isType:   values.DurationVar(nil),
			input:    "3h30m",
			output:   3*time.Hour + 30*time.Minute,
		},
		{
			name:     "duration list",
			setup:    func(r values.RegistererFunc) { r.DurationList("f", nil, "usg") },
			defValue: "",
			isType:   values.DurationList(),
			input:    "3h30m",
			output:   []time.Duration{3*time.Hour + 30*time.Minute},
		},
		{
			name:     "duration slice",
			setup:    func(r values.RegistererFunc) { r.DurationSlice("f", nil, ",", "usg") },
			defValue: "",
			isType:   values.DurationSlice(""),
			input:    "3h30m,1h15m,45s",
			output:   []time.Duration{3*time.Hour + 30*time.Minute, 1*time.Hour + 15*time.Minute, 45 * time.Second},
		},
		{
			name:     "time",
			setup:    func(r values.RegistererFunc) { r.Time("f", time.Time{}, time.RFC3339, "usg") },
			defValue: "0001-01-01T00:00:00Z",
			isType:   values.TimeVar(nil, ""),
			input:    "2025-05-07T06:06:06Z",
			output:   time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC),
		},
		{
			name:     "time list",
			setup:    func(r values.RegistererFunc) { r.TimeList("f", nil, time.RFC3339, "usg") },
			defValue: "",
			isType:   values.TimeList(""),
			input:    "2025-05-07T06:06:06Z",
			output:   []time.Time{time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC)},
		},
		{
			name:     "time slice",
			setup:    func(r values.RegistererFunc) { r.TimeSlice("f", nil, ",", time.RFC3339, "usg") },
			defValue: "",
			isType:   values.TimeSlice("", ""),
			input:    "2025-05-07T06:06:06Z,2025-05-08T07:07:07Z,2025-05-09T08:08:08Z",
			output: []time.Time{
				time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC),
				time.Date(2025, time.May, 8, 7, 7, 7, 0, time.UTC),
				time.Date(2025, time.May, 9, 8, 8, 8, 0, time.UTC),
			},
		},
		{
			name:     "ip address",
			setup:    func(r values.RegistererFunc) { r.IPAddr("f", netip.Addr{}, "usg") },
			defValue: "invalid IP",
			isType:   values.StringerVar[netip.Addr](nil, nil),
			input:    "192.168.1.1",
			output:   netip.MustParseAddr("192.168.1.1"),
		},
		{
			name:     "ip address list",
			setup:    func(r values.RegistererFunc) { r.IPAddrList("f", nil, "usg") },
			defValue: "",
			isType:   values.StringerList[netip.Addr](nil),
			input:    "192.168.1.1",
			output:   []netip.Addr{netip.MustParseAddr("192.168.1.1")},
		},
		{
			name:     "ip address slice",
			setup:    func(r values.RegistererFunc) { r.IPAddrSlice("f", nil, ",", "usg") },
			defValue: "",
			isType:   values.StringerSlice[netip.Addr]("", nil),
			input:    "192.168.1.1,10.0.0.1,172.16.0.1",
			output: []netip.Addr{
				netip.MustParseAddr("192.168.1.1"),
				netip.MustParseAddr("10.0.0.1"),
				netip.MustParseAddr("172.16.0.1"),
			},
		},
		{
			name:     "ip address and port",
			setup:    func(r values.RegistererFunc) { r.IPAddrPort("f", netip.AddrPort{}, "usg") },
			defValue: "invalid AddrPort",
			isType:   values.StringerVar[netip.AddrPort](nil, nil),
			input:    "192.168.1.1:8080",
			output:   netip.MustParseAddrPort("192.168.1.1:8080"),
		},
		{
			name:     "ip address and port list",
			setup:    func(r values.RegistererFunc) { r.IPAddrPortList("f", nil, "usg") },
			defValue: "",
			isType:   values.StringerList[netip.AddrPort](nil),
			input:    "192.168.1.1:8080",
			output:   []netip.AddrPort{netip.MustParseAddrPort("192.168.1.1:8080")},
		},
		{
			name:     "ip address and port slice",
			setup:    func(r values.RegistererFunc) { r.IPAddrPortSlice("f", nil, ",", "usg") },
			defValue: "",
			isType:   values.StringerSlice[netip.AddrPort]("", nil),
			input:    "192.168.1.1:8080,10.0.0.1:9090,172.16.0.1:7070",
			output: []netip.AddrPort{
				netip.MustParseAddrPort("192.168.1.1:8080"),
				netip.MustParseAddrPort("10.0.0.1:9090"),
				netip.MustParseAddrPort("172.16.0.1:7070"),
			},
		},
		{
			name:     "ip prefix",
			setup:    func(r values.RegistererFunc) { r.IPPrefix("f", netip.Prefix{}, "usg") },
			defValue: "invalid Prefix",
			isType:   values.StringerVar[netip.Prefix](nil, nil),
			input:    "192.168.1.0/24",
			output:   netip.MustParsePrefix("192.168.1.0/24"),
		},
		{
			name:     "ip prefix list",
			setup:    func(r values.RegistererFunc) { r.IPPrefixList("f", nil, "usg") },
			defValue: "",
			isType:   values.StringerList[netip.Prefix](nil),
			input:    "192.168.1.0/24",
			output:   []netip.Prefix{netip.MustParsePrefix("192.168.1.0/24")},
		},
		{
			name:     "ip prefix slice",
			setup:    func(r values.RegistererFunc) { r.IPPrefixSlice("f", nil, ",", "usg") },
			defValue: "",
			isType:   values.StringerSlice[netip.Prefix]("", nil),
			input:    "192.168.1.0/24,10.0.0.0/8,172.16.0.0/16",
			output: []netip.Prefix{
				netip.MustParsePrefix("192.168.1.0/24"),
				netip.MustParsePrefix("10.0.0.0/8"),
				netip.MustParsePrefix("172.16.0.0/16"),
			},
		},
		{
			name:     "mail addr",
			setup:    func(r values.RegistererFunc) { r.MailAddr("f", &mail.Address{}, "usg") },
			defValue: "<@>",
			isType:   values.StringerVar[*mail.Address](nil, nil),
			input:    "foo@bar.com",
			output:   &mail.Address{Address: "foo@bar.com"},
		},
		{
			name:     "mail addr list",
			setup:    func(r values.RegistererFunc) { r.MailAddrList("f", nil, "usg") },
			defValue: "",
			isType:   values.StringerList[*mail.Address](nil),
			input:    "foo@bar.com",
			output:   []*mail.Address{{Address: "foo@bar.com"}},
		},
		{
			name:     "mail addr slice",
			setup:    func(r values.RegistererFunc) { r.MailAddrSlice("f", nil, ",", "usg") },
			defValue: "",
			isType:   values.StringerSlice[*mail.Address]("", nil),
			input:    "foo@bar.com,baz@qux.com,quux@corge.com",
			output: []*mail.Address{
				{Address: "foo@bar.com"},
				{Address: "baz@qux.com"},
				{Address: "quux@corge.com"},
			},
		},
		{
			name:     "url",
			setup:    func(r values.RegistererFunc) { r.URL("f", &url.URL{}, "usg") },
			defValue: "",
			isType:   values.StringerVar[*url.URL](nil, nil),
			input:    "foo://bar",
			output:   &url.URL{Scheme: "foo", Host: "bar"},
		},
		{
			name:     "url list",
			setup:    func(r values.RegistererFunc) { r.URLList("f", nil, "usg") },
			defValue: "",
			isType:   values.StringerList[*url.URL](nil),
			input:    "foo://bar",
			output:   []*url.URL{{Scheme: "foo", Host: "bar"}},
		},
		{
			name:     "url slice",
			setup:    func(r values.RegistererFunc) { r.URLSlice("f", nil, ",", "usg") },
			defValue: "",
			isType:   values.StringerSlice[*url.URL]("", nil),
			input:    "foo://bar,baz://qux,quux://corge",
			output: []*url.URL{
				{Scheme: "foo", Host: "bar"},
				{Scheme: "baz", Host: "qux"},
				{Scheme: "quux", Host: "corge"},
			},
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

func TestRegisterer_valuesVar(t *testing.T) {
	testCases := []struct {
		name     string
		setup    func(values.RegistererFunc) func() any
		defValue string
		isType   any
		input    string
		output   any
	}{
		{
			name: "bool",
			setup: func(r values.RegistererFunc) func() any {
				p := new(bool)
				r.BoolVar(p, "f", false, "usg")
				return func() any { return *p }
			},
			defValue: "false",
			isType:   values.Basic[bool](),
			input:    "true",
			output:   true,
		},
		{
			name: "bool list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]bool)
				r.BoolListVar(p, "f", nil, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicList[bool](),
			input:    "true",
			output:   []bool{true},
		},
		{
			name: "bool slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]bool)
				r.BoolSliceVar(p, "f", nil, ",", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicSlice[bool](""),
			input:    "true,false,true",
			output:   []bool{true, false, true},
		},
		{
			name: "complex64",
			setup: func(r values.RegistererFunc) func() any {
				p := new(complex64)
				r.Complex64Var(p, "f", 0, "usg")
				return func() any { return *p }
			},
			defValue: "(0+0i)",
			isType:   values.Basic[complex64](),
			input:    "12+42i",
			output:   complex64(12 + 42i),
		},
		{
			name: "complex64 list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]complex64)
				r.Complex64ListVar(p, "f", nil, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicList[complex64](),
			input:    "12+42i",
			output:   []complex64{complex64(12 + 42i)},
		},
		{
			name: "complex64 slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]complex64)
				r.Complex64SliceVar(p, "f", nil, ",", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicSlice[complex64](""),
			input:    "12+42i,13+43i,14+44i",
			output:   []complex64{complex64(12 + 42i), complex64(13 + 43i), complex64(14 + 44i)},
		},
		{
			name: "complex128",
			setup: func(r values.RegistererFunc) func() any {
				p := new(complex128)
				r.Complex128Var(p, "f", 0, "usg")
				return func() any { return *p }
			},
			defValue: "(0+0i)",
			isType:   values.Basic[complex128](),
			input:    "15+62i",
			output:   complex128(15 + 62i),
		},
		{
			name: "complex128 list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]complex128)
				r.Complex128ListVar(p, "f", nil, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicList[complex128](),
			input:    "15+62i",
			output:   []complex128{complex128(15 + 62i)},
		},
		{
			name: "complex128 slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]complex128)
				r.Complex128SliceVar(p, "f", nil, ",", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicSlice[complex128](""),
			input:    "15+62i,16+63i,17+64i",
			output:   []complex128{complex128(15 + 62i), complex128(16 + 63i), complex128(17 + 64i)},
		},
		{
			name: "int",
			setup: func(r values.RegistererFunc) func() any {
				p := new(int)
				r.IntVar(p, "f", 0, "usg")
				return func() any { return *p }
			},
			defValue: "0",
			isType:   values.Basic[int](),
			input:    "-42",
			output:   int(-42),
		},
		{
			name: "int list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]int)
				r.IntListVar(p, "f", nil, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicList[int](),
			input:    "-42",
			output:   []int{int(-42)},
		},
		{
			name: "int slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]int)
				r.IntSliceVar(p, "f", nil, ",", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicSlice[int](""),
			input:    "-42,-43,-44",
			output:   []int{int(-42), int(-43), int(-44)},
		},
		{
			name: "int8",
			setup: func(r values.RegistererFunc) func() any {
				p := new(int8)
				r.Int8Var(p, "f", 0, "usg")
				return func() any { return *p }
			},
			defValue: "0",
			isType:   values.Basic[int8](),
			input:    "-42",
			output:   int8(-42),
		},
		{
			name: "int8 list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]int8)
				r.Int8ListVar(p, "f", nil, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicList[int8](),
			input:    "-42",
			output:   []int8{int8(-42)},
		},
		{
			name: "int8 slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]int8)
				r.Int8SliceVar(p, "f", nil, ",", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicSlice[int8](""),
			input:    "-42,-43,-44",
			output:   []int8{int8(-42), int8(-43), int8(-44)},
		},
		{
			name: "int16",
			setup: func(r values.RegistererFunc) func() any {
				p := new(int16)
				r.Int16Var(p, "f", 0, "usg")
				return func() any { return *p }
			},
			defValue: "0",
			isType:   values.Basic[int16](),
			input:    "-42",
			output:   int16(-42),
		},
		{
			name: "int16 list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]int16)
				r.Int16ListVar(p, "f", nil, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicList[int16](),
			input:    "-42",
			output:   []int16{int16(-42)},
		},
		{
			name: "int16 slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]int16)
				r.Int16SliceVar(p, "f", nil, ",", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicSlice[int16](""),
			input:    "-42,-43,-44",
			output:   []int16{int16(-42), int16(-43), int16(-44)},
		},
		{
			name: "int32",
			setup: func(r values.RegistererFunc) func() any {
				p := new(int32)
				r.Int32Var(p, "f", 0, "usg")
				return func() any { return *p }
			},
			defValue: "0",
			isType:   values.Basic[int32](),
			input:    "-42",
			output:   int32(-42),
		},
		{
			name: "int32 list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]int32)
				r.Int32ListVar(p, "f", nil, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicList[int32](),
			input:    "-42",
			output:   []int32{int32(-42)},
		},
		{
			name: "int32 slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]int32)
				r.Int32SliceVar(p, "f", nil, ",", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicSlice[int32](""),
			input:    "-42,-43,-44",
			output:   []int32{int32(-42), int32(-43), int32(-44)},
		},
		{
			name: "int64",
			setup: func(r values.RegistererFunc) func() any {
				p := new(int64)
				r.Int64Var(p, "f", 0, "usg")
				return func() any { return *p }
			},
			defValue: "0",
			isType:   values.Basic[int64](),
			input:    "-42",
			output:   int64(-42),
		},
		{
			name: "int64 list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]int64)
				r.Int64ListVar(p, "f", nil, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicList[int64](),
			input:    "-42",
			output:   []int64{int64(-42)},
		},
		{
			name: "int64 slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]int64)
				r.Int64SliceVar(p, "f", nil, ",", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicSlice[int64](""),
			input:    "-42,-43,-44",
			output:   []int64{int64(-42), int64(-43), int64(-44)},
		},
		{
			name: "uint",
			setup: func(r values.RegistererFunc) func() any {
				p := new(uint)
				r.UintVar(p, "f", 0, "usg")
				return func() any { return *p }
			},
			defValue: "0",
			isType:   values.Basic[uint](),
			input:    "42",
			output:   uint(42),
		},
		{
			name: "uint list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]uint)
				r.UintListVar(p, "f", nil, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicList[uint](),
			input:    "42",
			output:   []uint{uint(42)},
		},
		{
			name: "uint slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]uint)
				r.UintSliceVar(p, "f", nil, ",", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicSlice[uint](""),
			input:    "42,43,44",
			output:   []uint{uint(42), uint(43), uint(44)},
		},
		{
			name: "uint8",
			setup: func(r values.RegistererFunc) func() any {
				p := new(uint8)
				r.Uint8Var(p, "f", 0, "usg")
				return func() any { return *p }
			},
			defValue: "0",
			isType:   values.Basic[uint8](),
			input:    "42",
			output:   uint8(42),
		},
		{
			name: "uint8 list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]uint8)
				r.Uint8ListVar(p, "f", nil, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicList[uint8](),
			input:    "42",
			output:   []uint8{uint8(42)},
		},
		{
			name: "uint8 slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]uint8)
				r.Uint8SliceVar(p, "f", nil, ",", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicSlice[uint8](""),
			input:    "42,43,44",
			output:   []uint8{uint8(42), uint8(43), uint8(44)},
		},
		{
			name: "uint16",
			setup: func(r values.RegistererFunc) func() any {
				p := new(uint16)
				r.Uint16Var(p, "f", 0, "usg")
				return func() any { return *p }
			},
			defValue: "0",
			isType:   values.Basic[uint16](),
			input:    "42",
			output:   uint16(42),
		},
		{
			name: "uint16 list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]uint16)
				r.Uint16ListVar(p, "f", nil, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicList[uint16](),
			input:    "42",
			output:   []uint16{uint16(42)},
		},
		{
			name: "uint16 slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]uint16)
				r.Uint16SliceVar(p, "f", nil, ",", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicSlice[uint16](""),
			input:    "42,43,44",
			output:   []uint16{uint16(42), uint16(43), uint16(44)},
		},
		{
			name: "uint32",
			setup: func(r values.RegistererFunc) func() any {
				p := new(uint32)
				r.Uint32Var(p, "f", 0, "usg")
				return func() any { return *p }
			},
			defValue: "0",
			isType:   values.Basic[uint32](),
			input:    "42",
			output:   uint32(42),
		},
		{
			name: "uint32 list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]uint32)
				r.Uint32ListVar(p, "f", nil, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicList[uint32](),
			input:    "42",
			output:   []uint32{uint32(42)},
		},
		{
			name: "uint32 slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]uint32)
				r.Uint32SliceVar(p, "f", nil, ",", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicSlice[uint32](""),
			input:    "42,43,44",
			output:   []uint32{uint32(42), uint32(43), uint32(44)},
		},
		{
			name: "uint64",
			setup: func(r values.RegistererFunc) func() any {
				p := new(uint64)
				r.Uint64Var(p, "f", 0, "usg")
				return func() any { return *p }
			},
			defValue: "0",
			isType:   values.Basic[uint64](),
			input:    "42",
			output:   uint64(42),
		},
		{
			name: "uint64 list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]uint64)
				r.Uint64ListVar(p, "f", nil, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicList[uint64](),
			input:    "42",
			output:   []uint64{uint64(42)},
		},
		{
			name: "uint64 slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]uint64)
				r.Uint64SliceVar(p, "f", nil, ",", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicSlice[uint64](""),
			input:    "42,43,44",
			output:   []uint64{uint64(42), uint64(43), uint64(44)},
		},
		{
			name: "float32",
			setup: func(r values.RegistererFunc) func() any {
				p := new(float32)
				r.Float32Var(p, "f", 0, "usg")
				return func() any { return *p }
			},
			defValue: "0",
			isType:   values.Basic[float32](),
			input:    "3.14",
			output:   float32(3.14),
		},
		{
			name: "float32 list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]float32)
				r.Float32ListVar(p, "f", nil, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicList[float32](),
			input:    "3.14",
			output:   []float32{float32(3.14)},
		},
		{
			name: "float32 slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]float32)
				r.Float32SliceVar(p, "f", nil, ",", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicSlice[float32](""),
			input:    "3.14,2.71,1.41",
			output:   []float32{float32(3.14), float32(2.71), float32(1.41)},
		},
		{
			name: "float64",
			setup: func(r values.RegistererFunc) func() any {
				p := new(float64)
				r.Float64Var(p, "f", 0, "usg")
				return func() any { return *p }
			},
			defValue: "0",
			isType:   values.Basic[float64](),
			input:    "3.14159",
			output:   float64(3.14159),
		},
		{
			name: "float64 list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]float64)
				r.Float64ListVar(p, "f", nil, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicList[float64](),
			input:    "3.14159",
			output:   []float64{float64(3.14159)},
		},
		{
			name: "float64 slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]float64)
				r.Float64SliceVar(p, "f", nil, ",", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicSlice[float64](""),
			input:    "3.14159,2.71828,1.41421",
			output:   []float64{float64(3.14159), float64(2.71828), float64(1.41421)},
		},
		{
			name: "string",
			setup: func(r values.RegistererFunc) func() any {
				p := new(string)
				r.StringVar(p, "f", "", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.Basic[string](),
			input:    "hello world",
			output:   "hello world",
		},
		{
			name: "string list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]string)
				r.StringListVar(p, "f", nil, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicList[string](),
			input:    "hello world",
			output:   []string{"hello world"},
		},
		{
			name: "string slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]string)
				r.StringSliceVar(p, "f", nil, ",", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.BasicSlice[string](""),
			input:    "hello,world,foo",
			output:   []string{"hello", "world", "foo"},
		},
		{
			name: "duration",
			setup: func(r values.RegistererFunc) func() any {
				p := new(time.Duration)
				r.DurationVar(p, "f", 5*time.Minute, "usg")
				return func() any { return *p }
			},
			defValue: "5m0s",
			isType:   values.DurationVar(nil),
			input:    "3h30m",
			output:   3*time.Hour + 30*time.Minute,
		},
		{
			name: "duration list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]time.Duration)
				r.DurationListVar(p, "f", nil, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.DurationList(),
			input:    "3h30m",
			output:   []time.Duration{3*time.Hour + 30*time.Minute},
		},
		{
			name: "duration slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]time.Duration)
				r.DurationSliceVar(p, "f", nil, ",", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.DurationSlice(""),
			input:    "3h30m,1h15m,45s",
			output:   []time.Duration{3*time.Hour + 30*time.Minute, 1*time.Hour + 15*time.Minute, 45 * time.Second},
		},
		{
			name: "time",
			setup: func(r values.RegistererFunc) func() any {
				p := new(time.Time)
				r.TimeVar(p, "f", time.Time{}, time.RFC3339, "usg")
				return func() any { return *p }
			},
			defValue: "0001-01-01T00:00:00Z",
			isType:   values.TimeVar(nil, ""),
			input:    "2025-05-07T06:06:06Z",
			output:   time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC),
		},
		{
			name: "time list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]time.Time)
				r.TimeListVar(p, "f", nil, time.RFC3339, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.TimeList(""),
			input:    "2025-05-07T06:06:06Z",
			output:   []time.Time{time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC)},
		},
		{
			name: "time slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]time.Time)
				r.TimeSliceVar(p, "f", nil, ",", time.RFC3339, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.TimeSlice("", ""),
			input:    "2025-05-07T06:06:06Z,2025-05-08T07:07:07Z,2025-05-09T08:08:08Z",
			output: []time.Time{
				time.Date(2025, time.May, 7, 6, 6, 6, 0, time.UTC),
				time.Date(2025, time.May, 8, 7, 7, 7, 0, time.UTC),
				time.Date(2025, time.May, 9, 8, 8, 8, 0, time.UTC),
			},
		},
		{
			name: "ip address",
			setup: func(r values.RegistererFunc) func() any {
				p := new(netip.Addr)
				r.IPAddrVar(p, "f", netip.Addr{}, "usg")
				return func() any { return *p }
			},
			defValue: "invalid IP",
			isType:   values.StringerVar[netip.Addr](nil, nil),
			input:    "192.168.1.1",
			output:   netip.MustParseAddr("192.168.1.1"),
		},
		{
			name: "ip address list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]netip.Addr)
				r.IPAddrListVar(p, "f", nil, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.StringerList[netip.Addr](nil),
			input:    "192.168.1.1",
			output:   []netip.Addr{netip.MustParseAddr("192.168.1.1")},
		},
		{
			name: "ip address slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]netip.Addr)
				r.IPAddrSliceVar(p, "f", nil, ",", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.StringerSlice[netip.Addr]("", nil),
			input:    "192.168.1.1,10.0.0.1,172.16.0.1",
			output: []netip.Addr{
				netip.MustParseAddr("192.168.1.1"),
				netip.MustParseAddr("10.0.0.1"),
				netip.MustParseAddr("172.16.0.1"),
			},
		},
		{
			name: "ip address and port",
			setup: func(r values.RegistererFunc) func() any {
				p := new(netip.AddrPort)
				r.IPAddrPortVar(p, "f", netip.AddrPort{}, "usg")
				return func() any { return *p }
			},
			defValue: "invalid AddrPort",
			isType:   values.StringerVar[netip.AddrPort](nil, nil),
			input:    "192.168.1.1:8080",
			output:   netip.MustParseAddrPort("192.168.1.1:8080"),
		},
		{
			name: "ip address and port list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]netip.AddrPort)
				r.IPAddrPortListVar(p, "f", nil, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.StringerList[netip.AddrPort](nil),
			input:    "192.168.1.1:8080",
			output:   []netip.AddrPort{netip.MustParseAddrPort("192.168.1.1:8080")},
		},
		{
			name: "ip address and port slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]netip.AddrPort)
				r.IPAddrPortSliceVar(p, "f", nil, ",", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.StringerSlice[netip.AddrPort]("", nil),
			input:    "192.168.1.1:8080,10.0.0.1:9090,172.16.0.1:7070",
			output: []netip.AddrPort{
				netip.MustParseAddrPort("192.168.1.1:8080"),
				netip.MustParseAddrPort("10.0.0.1:9090"),
				netip.MustParseAddrPort("172.16.0.1:7070"),
			},
		},
		{
			name: "ip prefix",
			setup: func(r values.RegistererFunc) func() any {
				p := new(netip.Prefix)
				r.IPPrefixVar(p, "f", netip.Prefix{}, "usg")
				return func() any { return *p }
			},
			defValue: "invalid Prefix",
			isType:   values.StringerVar[netip.Prefix](nil, nil),
			input:    "192.168.1.0/24",
			output:   netip.MustParsePrefix("192.168.1.0/24"),
		},
		{
			name: "ip prefix list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]netip.Prefix)
				r.IPPrefixListVar(p, "f", nil, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.StringerList[netip.Prefix](nil),
			input:    "192.168.1.0/24",
			output:   []netip.Prefix{netip.MustParsePrefix("192.168.1.0/24")},
		},
		{
			name: "ip prefix slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]netip.Prefix)
				r.IPPrefixSliceVar(p, "f", nil, ",", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.StringerSlice[netip.Prefix]("", nil),
			input:    "192.168.1.0/24,10.0.0.0/8,172.16.0.0/16",
			output: []netip.Prefix{
				netip.MustParsePrefix("192.168.1.0/24"),
				netip.MustParsePrefix("10.0.0.0/8"),
				netip.MustParsePrefix("172.16.0.0/16"),
			},
		},
		{
			name: "mail addr",
			setup: func(r values.RegistererFunc) func() any {
				p := new(*mail.Address)
				r.MailAddrVar(p, "f", &mail.Address{}, "usg")
				return func() any { return *p }
			},
			defValue: "<@>",
			isType:   values.StringerVar[*mail.Address](nil, nil),
			input:    "foo@bar.com",
			output:   &mail.Address{Address: "foo@bar.com"},
		},
		{
			name: "mail addr list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]*mail.Address)
				r.MailAddrListVar(p, "f", nil, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.StringerList[*mail.Address](nil),
			input:    "foo@bar.com",
			output:   []*mail.Address{{Address: "foo@bar.com"}},
		},
		{
			name: "mail addr slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]*mail.Address)
				r.MailAddrSliceVar(p, "f", nil, ",", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.StringerSlice[*mail.Address]("", nil),
			input:    "foo@bar.com,baz@qux.com,quux@corge.com",
			output: []*mail.Address{
				{Address: "foo@bar.com"},
				{Address: "baz@qux.com"},
				{Address: "quux@corge.com"},
			},
		},
		{
			name: "url",
			setup: func(r values.RegistererFunc) func() any {
				p := new(*url.URL)
				r.URLVar(p, "f", &url.URL{}, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.StringerVar[*url.URL](nil, nil),
			input:    "foo://bar",
			output:   &url.URL{Scheme: "foo", Host: "bar"},
		},
		{
			name: "url list",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]*url.URL)
				r.URLListVar(p, "f", nil, "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.StringerList[*url.URL](nil),
			input:    "foo://bar",
			output:   []*url.URL{{Scheme: "foo", Host: "bar"}},
		},
		{
			name: "url slice",
			setup: func(r values.RegistererFunc) func() any {
				p := new([]*url.URL)
				r.URLSliceVar(p, "f", nil, ",", "usg")
				return func() any { return *p }
			},
			defValue: "",
			isType:   values.StringerSlice[*url.URL]("", nil),
			input:    "foo://bar,baz://qux,quux://corge",
			output: []*url.URL{
				{Scheme: "foo", Host: "bar"},
				{Scheme: "baz", Host: "qux"},
				{Scheme: "quux", Host: "corge"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fs := flag.NewFlagSet("", flag.ContinueOnError)
			get := tc.setup(values.FlagSetRegisterer(fs))

			f := fs.Lookup("f")
			require.NotNil(t, f)
			require.Equal(t, "f", f.Name)
			require.Equal(t, "usg", f.Usage)
			require.Equal(t, tc.defValue, f.DefValue)
			require.IsType(t, tc.isType, f.Value)
			require.NoError(t, fs.Parse([]string{"-f", tc.input}))
			require.Equal(t, tc.output, f.Value.(flag.Getter).Get())
			require.Equal(t, tc.output, get())
		})
	}
}
