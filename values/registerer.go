package values

import (
	"flag"
	"fmt"
	"net/mail"
	"net/netip"
	"net/url"
	"os"
	"strings"
	"time"
)

// Registerer provides an interface analogous to [*flag.FlagSet] for defining
// flags for common types. It trades the flexibility of using plain
// [flag.Value] for convenience and simplicity.
type Registerer struct {
	Varer interface {
		Var(value flag.Value, name string, usage string)
	}
}

type varerFunc func(value flag.Value, name string, usage string)

func (f varerFunc) Var(value flag.Value, name string, usage string) { f(value, name, usage) }

// WithEnvFunc is like [Registerer.WithEnv] but uses the envvar function to map a flag to its environment variable.
func (r Registerer) WithEnvFunc(envvar func(name string) string) Registerer {
	return Registerer{
		Varer: varerFunc(func(value flag.Value, name, usage string) {
			envvar := envvar(name)
			r.Varer.Var(value, name, fmt.Sprintf("%s (env $%s)", usage, envvar))
			if val, ok := os.LookupEnv(envvar); ok {
				value.Set(val) //nolint: errcheck,gosec // ignore environment then
			}
		}),
	}
}

// WithEnv returns a new [Registerer] setting registered flag values from environment and editing usage accordingly.
// The environment variable is mapped from the flag name by:
//
//   - replacing dashes and dots by underscores
//   - transforming to upper case
//   - prefixing
//
// The environment variable is ignored if it fails to set the flag value.
func (r Registerer) WithEnv(prefix string) Registerer {
	replacer := strings.NewReplacer("-", "_", ".", "_")
	return r.WithEnvFunc(func(s string) string {
		s = replacer.Replace(s)
		s = strings.ToUpper(s)
		return prefix + s
	})
}

// FlagSetRegisterer returns a [*flag.FlagSet] based [Registerer].
func FlagSetRegisterer(fs *flag.FlagSet) Registerer { return Registerer{fs} }

// Bool defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func (r Registerer) Bool(name string, value bool, usage string) *bool {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func (r Registerer) BoolVar(p *bool, name string, value bool, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

// BoolList defines a list-style bool flag with specified name, default value, and usage string.
// The return value is the address of a bool slice that stores the values of the flag.
func (r Registerer) BoolList(name string, value []bool, usage string) *[]bool {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

// BoolListVar defines a list-style bool flag with specified name, default value, and usage string.
// The argument p points to a bool slice variable in which to store the value of the flag.
func (r Registerer) BoolListVar(p *[]bool, name string, value []bool, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

// BoolSlice defines a slice-style bool flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a bool slice that stores the values of the flag.
func (r Registerer) BoolSlice(name string, value []bool, sep string, usage string) *[]bool {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// BoolSliceVar defines a slice-style bool flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a bool slice variable in which to store the value of the flag.
func (r Registerer) BoolSliceVar(p *[]bool, name string, value []bool, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

// Complex64 defines a complex64 flag with specified name, default value, and usage string.
// The return value is the address of a complex64 variable that stores the value of the flag.
func (r Registerer) Complex64(name string, value complex64, usage string) *complex64 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

// Complex64Var defines a complex64 flag with specified name, default value, and usage string.
// The argument p points to a complex64 variable in which to store the value of the flag.
func (r Registerer) Complex64Var(p *complex64, name string, value complex64, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

// Complex64List defines a list-style complex64 flag with specified name, default value, and usage string.
// The return value is the address of a complex64 slice that stores the values of the flag.
func (r Registerer) Complex64List(name string, value []complex64, usage string) *[]complex64 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

// Complex64ListVar defines a list-style complex64 flag with specified name, default value, and usage string.
// The argument p points to a complex64 slice variable in which to store the value of the flag.
func (r Registerer) Complex64ListVar(p *[]complex64, name string, value []complex64, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

// Complex64Slice defines a slice-style complex64 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a complex64 slice that stores the values of the flag.
func (r Registerer) Complex64Slice(name string, value []complex64, sep string, usage string) *[]complex64 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Complex64SliceVar defines a slice-style complex64 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a complex64 slice variable in which to store the value of the flag.
func (r Registerer) Complex64SliceVar(p *[]complex64, name string, value []complex64, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

// Complex128 defines a complex128 flag with specified name, default value, and usage string.
// The return value is the address of a complex128 variable that stores the value of the flag.
func (r Registerer) Complex128(name string, value complex128, usage string) *complex128 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

// Complex128Var defines a complex128 flag with specified name, default value, and usage string.
// The argument p points to a complex128 variable in which to store the value of the flag.
func (r Registerer) Complex128Var(p *complex128, name string, value complex128, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

// Complex128List defines a list-style complex128 flag with specified name, default value, and usage string.
// The return value is the address of a complex128 slice that stores the values of the flag.
func (r Registerer) Complex128List(name string, value []complex128, usage string) *[]complex128 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

// Complex128ListVar defines a list-style complex128 flag with specified name, default value, and usage string.
// The argument p points to a complex128 slice variable in which to store the value of the flag.
func (r Registerer) Complex128ListVar(p *[]complex128, name string, value []complex128, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

// Complex128Slice defines a slice-style complex128 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a complex128 slice that stores the values of the flag.
func (r Registerer) Complex128Slice(name string, value []complex128, sep string, usage string) *[]complex128 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Complex128SliceVar defines a slice-style complex128 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a complex128 slice variable in which to store the value of the flag.
func (r Registerer) Complex128SliceVar(p *[]complex128, name string, value []complex128, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

// Int defines a int flag with specified name, default value, and usage string.
// The return value is the address of a int variable that stores the value of the flag.
func (r Registerer) Int(name string, value int, usage string) *int {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

// IntVar defines a int flag with specified name, default value, and usage string.
// The argument p points to a int variable in which to store the value of the flag.
func (r Registerer) IntVar(p *int, name string, value int, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

// IntList defines a list-style int flag with specified name, default value, and usage string.
// The return value is the address of a int slice that stores the values of the flag.
func (r Registerer) IntList(name string, value []int, usage string) *[]int {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

// IntListVar defines a list-style int flag with specified name, default value, and usage string.
// The argument p points to a int slice variable in which to store the value of the flag.
func (r Registerer) IntListVar(p *[]int, name string, value []int, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

// IntSlice defines a slice-style int flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a int slice that stores the values of the flag.
func (r Registerer) IntSlice(name string, value []int, sep string, usage string) *[]int {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// IntSliceVar defines a slice-style int flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a int slice variable in which to store the value of the flag.
func (r Registerer) IntSliceVar(p *[]int, name string, value []int, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

// Int8 defines a int8 flag with specified name, default value, and usage string.
// The return value is the address of a int8 variable that stores the value of the flag.
func (r Registerer) Int8(name string, value int8, usage string) *int8 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

// Int8Var defines a int8 flag with specified name, default value, and usage string.
// The argument p points to a int8 variable in which to store the value of the flag.
func (r Registerer) Int8Var(p *int8, name string, value int8, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

// Int8List defines a list-style int8 flag with specified name, default value, and usage string.
// The return value is the address of a int8 slice that stores the values of the flag.
func (r Registerer) Int8List(name string, value []int8, usage string) *[]int8 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

// Int8ListVar defines a list-style int8 flag with specified name, default value, and usage string.
// The argument p points to a int8 slice variable in which to store the value of the flag.
func (r Registerer) Int8ListVar(p *[]int8, name string, value []int8, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

// Int8Slice defines a slice-style int8 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a int8 slice that stores the values of the flag.
func (r Registerer) Int8Slice(name string, value []int8, sep string, usage string) *[]int8 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Int8SliceVar defines a slice-style int8 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a int8 slice variable in which to store the value of the flag.
func (r Registerer) Int8SliceVar(p *[]int8, name string, value []int8, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

// Int16 defines a int16 flag with specified name, default value, and usage string.
// The return value is the address of a int16 variable that stores the value of the flag.
func (r Registerer) Int16(name string, value int16, usage string) *int16 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

// Int16Var defines a int16 flag with specified name, default value, and usage string.
// The argument p points to a int16 variable in which to store the value of the flag.
func (r Registerer) Int16Var(p *int16, name string, value int16, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

// Int16List defines a list-style int16 flag with specified name, default value, and usage string.
// The return value is the address of a int16 slice that stores the values of the flag.
func (r Registerer) Int16List(name string, value []int16, usage string) *[]int16 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

// Int16ListVar defines a list-style int16 flag with specified name, default value, and usage string.
// The argument p points to a int16 slice variable in which to store the value of the flag.
func (r Registerer) Int16ListVar(p *[]int16, name string, value []int16, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

// Int16Slice defines a slice-style int16 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a int16 slice that stores the values of the flag.
func (r Registerer) Int16Slice(name string, value []int16, sep string, usage string) *[]int16 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Int16SliceVar defines a slice-style int16 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a int16 slice variable in which to store the value of the flag.
func (r Registerer) Int16SliceVar(p *[]int16, name string, value []int16, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

// Int32 defines a int32 flag with specified name, default value, and usage string.
// The return value is the address of a int32 variable that stores the value of the flag.
func (r Registerer) Int32(name string, value int32, usage string) *int32 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

// Int32Var defines a int32 flag with specified name, default value, and usage string.
// The argument p points to a int32 variable in which to store the value of the flag.
func (r Registerer) Int32Var(p *int32, name string, value int32, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

// Int32List defines a list-style int32 flag with specified name, default value, and usage string.
// The return value is the address of a int32 slice that stores the values of the flag.
func (r Registerer) Int32List(name string, value []int32, usage string) *[]int32 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

// Int32ListVar defines a list-style int32 flag with specified name, default value, and usage string.
// The argument p points to a int32 slice variable in which to store the value of the flag.
func (r Registerer) Int32ListVar(p *[]int32, name string, value []int32, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

// Int32Slice defines a slice-style int32 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a int32 slice that stores the values of the flag.
func (r Registerer) Int32Slice(name string, value []int32, sep string, usage string) *[]int32 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Int32SliceVar defines a slice-style int32 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a int32 slice variable in which to store the value of the flag.
func (r Registerer) Int32SliceVar(p *[]int32, name string, value []int32, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

// Int64 defines a int64 flag with specified name, default value, and usage string.
// The return value is the address of a int64 variable that stores the value of the flag.
func (r Registerer) Int64(name string, value int64, usage string) *int64 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

// Int64Var defines a int64 flag with specified name, default value, and usage string.
// The argument p points to a int64 variable in which to store the value of the flag.
func (r Registerer) Int64Var(p *int64, name string, value int64, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

// Int64List defines a list-style int64 flag with specified name, default value, and usage string.
// The return value is the address of a int64 slice that stores the values of the flag.
func (r Registerer) Int64List(name string, value []int64, usage string) *[]int64 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

// Int64ListVar defines a list-style int64 flag with specified name, default value, and usage string.
// The argument p points to a int64 slice variable in which to store the value of the flag.
func (r Registerer) Int64ListVar(p *[]int64, name string, value []int64, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

// Int64Slice defines a slice-style int64 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a int64 slice that stores the values of the flag.
func (r Registerer) Int64Slice(name string, value []int64, sep string, usage string) *[]int64 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Int64SliceVar defines a slice-style int64 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a int64 slice variable in which to store the value of the flag.
func (r Registerer) Int64SliceVar(p *[]int64, name string, value []int64, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

// Uint defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint variable that stores the value of the flag.
func (r Registerer) Uint(name string, value uint, usage string) *uint {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

// UintVar defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func (r Registerer) UintVar(p *uint, name string, value uint, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

// UintList defines a list-style uint flag with specified name, default value, and usage string.
// The return value is the address of a uint slice that stores the values of the flag.
func (r Registerer) UintList(name string, value []uint, usage string) *[]uint {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

// UintListVar defines a list-style uint flag with specified name, default value, and usage string.
// The argument p points to a uint slice variable in which to store the value of the flag.
func (r Registerer) UintListVar(p *[]uint, name string, value []uint, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

// UintSlice defines a slice-style uint flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a uint slice that stores the values of the flag.
func (r Registerer) UintSlice(name string, value []uint, sep string, usage string) *[]uint {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// UintSliceVar defines a slice-style uint flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a uint slice variable in which to store the value of the flag.
func (r Registerer) UintSliceVar(p *[]uint, name string, value []uint, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

// Uint8 defines a uint8 flag with specified name, default value, and usage string.
// The return value is the address of a uint8 variable that stores the value of the flag.
func (r Registerer) Uint8(name string, value uint8, usage string) *uint8 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

// Uint8Var defines a uint8 flag with specified name, default value, and usage string.
// The argument p points to a uint8 variable in which to store the value of the flag.
func (r Registerer) Uint8Var(p *uint8, name string, value uint8, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

// Uint8List defines a list-style uint8 flag with specified name, default value, and usage string.
// The return value is the address of a uint8 slice that stores the values of the flag.
func (r Registerer) Uint8List(name string, value []uint8, usage string) *[]uint8 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

// Uint8ListVar defines a list-style uint8 flag with specified name, default value, and usage string.
// The argument p points to a uint8 slice variable in which to store the value of the flag.
func (r Registerer) Uint8ListVar(p *[]uint8, name string, value []uint8, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

// Uint8Slice defines a slice-style uint8 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a uint8 slice that stores the values of the flag.
func (r Registerer) Uint8Slice(name string, value []uint8, sep string, usage string) *[]uint8 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Uint8SliceVar defines a slice-style uint8 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a uint8 slice variable in which to store the value of the flag.
func (r Registerer) Uint8SliceVar(p *[]uint8, name string, value []uint8, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

// Uint16 defines a uint16 flag with specified name, default value, and usage string.
// The return value is the address of a uint16 variable that stores the value of the flag.
func (r Registerer) Uint16(name string, value uint16, usage string) *uint16 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

// Uint16Var defines a uint16 flag with specified name, default value, and usage string.
// The argument p points to a uint16 variable in which to store the value of the flag.
func (r Registerer) Uint16Var(p *uint16, name string, value uint16, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

// Uint16List defines a list-style uint16 flag with specified name, default value, and usage string.
// The return value is the address of a uint16 slice that stores the values of the flag.
func (r Registerer) Uint16List(name string, value []uint16, usage string) *[]uint16 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

// Uint16ListVar defines a list-style uint16 flag with specified name, default value, and usage string.
// The argument p points to a uint16 slice variable in which to store the value of the flag.
func (r Registerer) Uint16ListVar(p *[]uint16, name string, value []uint16, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

// Uint16Slice defines a slice-style uint16 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a uint16 slice that stores the values of the flag.
func (r Registerer) Uint16Slice(name string, value []uint16, sep string, usage string) *[]uint16 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Uint16SliceVar defines a slice-style uint16 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a uint16 slice variable in which to store the value of the flag.
func (r Registerer) Uint16SliceVar(p *[]uint16, name string, value []uint16, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

// Uint32 defines a uint32 flag with specified name, default value, and usage string.
// The return value is the address of a uint32 variable that stores the value of the flag.
func (r Registerer) Uint32(name string, value uint32, usage string) *uint32 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

// Uint32Var defines a uint32 flag with specified name, default value, and usage string.
// The argument p points to a uint32 variable in which to store the value of the flag.
func (r Registerer) Uint32Var(p *uint32, name string, value uint32, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

// Uint32List defines a list-style uint32 flag with specified name, default value, and usage string.
// The return value is the address of a uint32 slice that stores the values of the flag.
func (r Registerer) Uint32List(name string, value []uint32, usage string) *[]uint32 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

// Uint32ListVar defines a list-style uint32 flag with specified name, default value, and usage string.
// The argument p points to a uint32 slice variable in which to store the value of the flag.
func (r Registerer) Uint32ListVar(p *[]uint32, name string, value []uint32, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

// Uint32Slice defines a slice-style uint32 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a uint32 slice that stores the values of the flag.
func (r Registerer) Uint32Slice(name string, value []uint32, sep string, usage string) *[]uint32 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Uint32SliceVar defines a slice-style uint32 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a uint32 slice variable in which to store the value of the flag.
func (r Registerer) Uint32SliceVar(p *[]uint32, name string, value []uint32, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
func (r Registerer) Uint64(name string, value uint64, usage string) *uint64 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func (r Registerer) Uint64Var(p *uint64, name string, value uint64, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

// Uint64List defines a list-style uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 slice that stores the values of the flag.
func (r Registerer) Uint64List(name string, value []uint64, usage string) *[]uint64 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

// Uint64ListVar defines a list-style uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 slice variable in which to store the value of the flag.
func (r Registerer) Uint64ListVar(p *[]uint64, name string, value []uint64, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

// Uint64Slice defines a slice-style uint64 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a uint64 slice that stores the values of the flag.
func (r Registerer) Uint64Slice(name string, value []uint64, sep string, usage string) *[]uint64 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Uint64SliceVar defines a slice-style uint64 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a uint64 slice variable in which to store the value of the flag.
func (r Registerer) Uint64SliceVar(p *[]uint64, name string, value []uint64, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

// Float32 defines a float32 flag with specified name, default value, and usage string.
// The return value is the address of a float32 variable that stores the value of the flag.
func (r Registerer) Float32(name string, value float32, usage string) *float32 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

// Float32Var defines a float32 flag with specified name, default value, and usage string.
// The argument p points to a float32 variable in which to store the value of the flag.
func (r Registerer) Float32Var(p *float32, name string, value float32, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

// Float32List defines a list-style float32 flag with specified name, default value, and usage string.
// The return value is the address of a float32 slice that stores the values of the flag.
func (r Registerer) Float32List(name string, value []float32, usage string) *[]float32 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

// Float32ListVar defines a list-style float32 flag with specified name, default value, and usage string.
// The argument p points to a float32 slice variable in which to store the value of the flag.
func (r Registerer) Float32ListVar(p *[]float32, name string, value []float32, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

// Float32Slice defines a slice-style float32 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a float32 slice that stores the values of the flag.
func (r Registerer) Float32Slice(name string, value []float32, sep string, usage string) *[]float32 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Float32SliceVar defines a slice-style float32 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a float32 slice variable in which to store the value of the flag.
func (r Registerer) Float32SliceVar(p *[]float32, name string, value []float32, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

// Float64 defines a float64 flag with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag.
func (r Registerer) Float64(name string, value float64, usage string) *float64 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

// Float64Var defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
func (r Registerer) Float64Var(p *float64, name string, value float64, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

// Float64List defines a list-style float64 flag with specified name, default value, and usage string.
// The return value is the address of a float64 slice that stores the values of the flag.
func (r Registerer) Float64List(name string, value []float64, usage string) *[]float64 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

// Float64ListVar defines a list-style float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 slice variable in which to store the value of the flag.
func (r Registerer) Float64ListVar(p *[]float64, name string, value []float64, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

// Float64Slice defines a slice-style float64 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a float64 slice that stores the values of the flag.
func (r Registerer) Float64Slice(name string, value []float64, sep string, usage string) *[]float64 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Float64SliceVar defines a slice-style float64 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a float64 slice variable in which to store the value of the flag.
func (r Registerer) Float64SliceVar(p *[]float64, name string, value []float64, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

// String defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func (r Registerer) String(name string, value string, usage string) *string {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func (r Registerer) StringVar(p *string, name string, value string, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

// StringList defines a list-style string flag with specified name, default value, and usage string.
// The return value is the address of a string slice that stores the values of the flag.
func (r Registerer) StringList(name string, value []string, usage string) *[]string {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

// StringListVar defines a list-style string flag with specified name, default value, and usage string.
// The argument p points to a string slice variable in which to store the value of the flag.
func (r Registerer) StringListVar(p *[]string, name string, value []string, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

// StringSlice defines a slice-style string flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a string slice that stores the values of the flag.
func (r Registerer) StringSlice(name string, value []string, sep string, usage string) *[]string {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// StringSliceVar defines a slice-style string flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a string slice variable in which to store the value of the flag.
func (r Registerer) StringSliceVar(p *[]string, name string, value []string, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

// Duration defines a [time.Duration] flag with specified name, default value, and usage string.
// The return value is the address of a [time.Duration] variable that stores the value of the flag.
func (r Registerer) Duration(name string, value time.Duration, usage string) *time.Duration {
	r.Varer.Var(DurationVar(&value), name, usage)
	return &value
}

// DurationVar defines a [time.Duration] flag with specified name, default value, and usage string.
// The argument p points to a [time.Duration] variable in which to store the value of the flag.
func (r Registerer) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	*p = value
	r.Varer.Var(DurationVar(p), name, usage)
}

// DurationList defines a list-style [time.Duration] flag with specified name, default value, and usage string.
// The return value is the address of a [time.Duration] slice that stores the values of the flag.
func (r Registerer) DurationList(name string, value []time.Duration, usage string) *[]time.Duration {
	r.Varer.Var(DurationListVar(&value), name, usage)
	return &value
}

// DurationListVar defines a list-style [time.Duration] flag with specified name, default value, and usage string.
// The argument p points to a [time.Duration] slice variable in which to store the value of the flag.
func (r Registerer) DurationListVar(p *[]time.Duration, name string, value []time.Duration, usage string) {
	*p = value
	r.Varer.Var(DurationListVar(p), name, usage)
}

// DurationSlice defines a slice-style [time.Duration] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a [time.Duration] slice that stores the values of the flag.
func (r Registerer) DurationSlice(name string, value []time.Duration, sep string, usage string) *[]time.Duration {
	r.Varer.Var(DurationSliceVar(&value, sep), name, usage)
	return &value
}

// DurationSliceVar defines a slice-style [time.Duration] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a [time.Duration] slice variable in which to store the value of the flag.
func (r Registerer) DurationSliceVar(p *[]time.Duration, name string, value []time.Duration, sep string, usage string) {
	*p = value
	r.Varer.Var(DurationSliceVar(p, sep), name, usage)
}

// IPAddr defines a [netip.Addr] flag with specified name, default value, and usage string.
// The return value is the address of a [netip.Addr] variable that stores the value of the flag.
func (r Registerer) IPAddr(name string, value netip.Addr, usage string) *netip.Addr {
	r.Varer.Var(StringerVar(&value, netip.ParseAddr), name, usage)
	return &value
}

// IPAddrVar defines a [netip.Addr] flag with specified name, default value, and usage string.
// The argument p points to a [netip.Addr] variable in which to store the value of the flag.
func (r Registerer) IPAddrVar(p *netip.Addr, name string, value netip.Addr, usage string) {
	*p = value
	r.Varer.Var(StringerVar(p, netip.ParseAddr), name, usage)
}

// IPAddrList defines a list-style [netip.Addr] flag with specified name, default value, and usage string.
// The return value is the address of a [netip.Addr] slice that stores the values of the flag.
func (r Registerer) IPAddrList(name string, value []netip.Addr, usage string) *[]netip.Addr {
	r.Varer.Var(StringerListVar(&value, netip.ParseAddr), name, usage)
	return &value
}

// IPAddrListVar defines a list-style [netip.Addr] flag with specified name, default value, and usage string.
// The argument p points to a [netip.Addr] slice variable in which to store the value of the flag.
func (r Registerer) IPAddrListVar(p *[]netip.Addr, name string, value []netip.Addr, usage string) {
	*p = value
	r.Varer.Var(StringerListVar(p, netip.ParseAddr), name, usage)
}

// IPAddrSlice defines a slice-style [netip.Addr] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a [netip.Addr] slice that stores the values of the flag.
func (r Registerer) IPAddrSlice(name string, value []netip.Addr, sep string, usage string) *[]netip.Addr {
	r.Varer.Var(StringerSliceVar(&value, sep, netip.ParseAddr), name, usage)
	return &value
}

// IPAddrSliceVar defines a slice-style [netip.Addr] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a [netip.Addr] slice variable in which to store the value of the flag.
func (r Registerer) IPAddrSliceVar(p *[]netip.Addr, name string, value []netip.Addr, sep string, usage string) {
	*p = value
	r.Varer.Var(StringerSliceVar(p, sep, netip.ParseAddr), name, usage)
}

// IPAddrPort defines a [netip.AddrPort] flag with specified name, default value, and usage string.
// The return value is the address of a [netip.AddrPort] variable that stores the value of the flag.
func (r Registerer) IPAddrPort(name string, value netip.AddrPort, usage string) *netip.AddrPort {
	r.Varer.Var(StringerVar(&value, netip.ParseAddrPort), name, usage)
	return &value
}

// IPAddrPortVar defines a [netip.AddrPort] flag with specified name, default value, and usage string.
// The argument p points to a [netip.AddrPort] variable in which to store the value of the flag.
func (r Registerer) IPAddrPortVar(p *netip.AddrPort, name string, value netip.AddrPort, usage string) {
	*p = value
	r.Varer.Var(StringerVar(p, netip.ParseAddrPort), name, usage)
}

// IPAddrPortList defines a list-style [netip.AddrPort] flag with specified name, default value, and usage string.
// The return value is the address of a [netip.AddrPort] slice that stores the values of the flag.
func (r Registerer) IPAddrPortList(name string, value []netip.AddrPort, usage string) *[]netip.AddrPort {
	r.Varer.Var(StringerListVar(&value, netip.ParseAddrPort), name, usage)
	return &value
}

// IPAddrPortListVar defines a list-style [netip.AddrPort] flag with specified name, default value, and usage string.
// The argument p points to a [netip.AddrPort] slice variable in which to store the value of the flag.
func (r Registerer) IPAddrPortListVar(p *[]netip.AddrPort, name string, value []netip.AddrPort, usage string) {
	*p = value
	r.Varer.Var(StringerListVar(p, netip.ParseAddrPort), name, usage)
}

// IPAddrPortSlice defines a slice-style [netip.AddrPort] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a [netip.AddrPort] slice that stores the values of the flag.
func (r Registerer) IPAddrPortSlice(name string, value []netip.AddrPort, sep string, usage string) *[]netip.AddrPort {
	r.Varer.Var(StringerSliceVar(&value, sep, netip.ParseAddrPort), name, usage)
	return &value
}

// IPAddrPortSliceVar defines a slice-style [netip.AddrPort] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a [netip.AddrPort] slice variable in which to store the value of the flag.
func (r Registerer) IPAddrPortSliceVar(p *[]netip.AddrPort, name string, value []netip.AddrPort, sep string, usage string) { //nolint: golines
	*p = value
	r.Varer.Var(StringerSliceVar(p, sep, netip.ParseAddrPort), name, usage)
}

// IPPrefix defines a [netip.Prefix] flag with specified name, default value, and usage string.
// The return value is the address of a [netip.Prefix] variable that stores the value of the flag.
func (r Registerer) IPPrefix(name string, value netip.Prefix, usage string) *netip.Prefix {
	r.Varer.Var(StringerVar(&value, netip.ParsePrefix), name, usage)
	return &value
}

// IPPrefixVar defines a [netip.Prefix] flag with specified name, default value, and usage string.
// The argument p points to a [netip.Prefix] variable in which to store the value of the flag.
func (r Registerer) IPPrefixVar(p *netip.Prefix, name string, value netip.Prefix, usage string) {
	*p = value
	r.Varer.Var(StringerVar(p, netip.ParsePrefix), name, usage)
}

// IPPrefixList defines a list-style [netip.Prefix] flag with specified name, default value, and usage string.
// The return value is the address of a [netip.Prefix] slice that stores the values of the flag.
func (r Registerer) IPPrefixList(name string, value []netip.Prefix, usage string) *[]netip.Prefix {
	r.Varer.Var(StringerListVar(&value, netip.ParsePrefix), name, usage)
	return &value
}

// IPPrefixListVar defines a list-style [netip.Prefix] flag with specified name, default value, and usage string.
// The argument p points to a [netip.Prefix] slice variable in which to store the value of the flag.
func (r Registerer) IPPrefixListVar(p *[]netip.Prefix, name string, value []netip.Prefix, usage string) {
	*p = value
	r.Varer.Var(StringerListVar(p, netip.ParsePrefix), name, usage)
}

// IPPrefixSlice defines a slice-style [netip.Prefix] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a [netip.Prefix] slice that stores the values of the flag.
func (r Registerer) IPPrefixSlice(name string, value []netip.Prefix, sep string, usage string) *[]netip.Prefix {
	r.Varer.Var(StringerSliceVar(&value, sep, netip.ParsePrefix), name, usage)
	return &value
}

// IPPrefixSliceVar defines a slice-style [netip.Prefix] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a [netip.Prefix] slice variable in which to store the value of the flag.
func (r Registerer) IPPrefixSliceVar(p *[]netip.Prefix, name string, value []netip.Prefix, sep string, usage string) {
	*p = value
	r.Varer.Var(StringerSliceVar(p, sep, netip.ParsePrefix), name, usage)
}

// MailAddr defines a [*mail.Address] flag with specified name, default value, and usage string.
// The return value is the address of a [*mail.Address] variable that stores the value of the flag.
func (r Registerer) MailAddr(name string, value *mail.Address, usage string) **mail.Address {
	r.Varer.Var(StringerVar(&value, mail.ParseAddress), name, usage)
	return &value
}

// MailAddrVar defines a [*mail.Address] flag with specified name, default value, and usage string.
// The argument p points to a [*mail.Address] variable in which to store the value of the flag.
func (r Registerer) MailAddrVar(p **mail.Address, name string, value *mail.Address, usage string) {
	*p = value
	r.Varer.Var(StringerVar(p, mail.ParseAddress), name, usage)
}

// MailAddrList defines a list-style [*mail.Address] flag with specified name, default value, and usage string.
// The return value is the address of a [*mail.Address] slice that stores the values of the flag.
func (r Registerer) MailAddrList(name string, value []*mail.Address, usage string) *[]*mail.Address {
	r.Varer.Var(StringerListVar(&value, mail.ParseAddress), name, usage)
	return &value
}

// MailAddrListVar defines a list-style [*mail.Address] flag with specified name, default value, and usage string.
// The argument p points to a [*mail.Address] slice variable in which to store the value of the flag.
func (r Registerer) MailAddrListVar(p *[]*mail.Address, name string, value []*mail.Address, usage string) {
	*p = value
	r.Varer.Var(StringerListVar(p, mail.ParseAddress), name, usage)
}

// MailAddrSlice defines a slice-style [*mail.Address] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a [*mail.Address] slice that stores the values of the flag.
func (r Registerer) MailAddrSlice(name string, value []*mail.Address, sep string, usage string) *[]*mail.Address {
	r.Varer.Var(StringerSliceVar(&value, sep, mail.ParseAddress), name, usage)
	return &value
}

// MailAddrSliceVar defines a slice-style [*mail.Address] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a [*mail.Address] slice variable in which to store the value of the flag.
func (r Registerer) MailAddrSliceVar(p *[]*mail.Address, name string, value []*mail.Address, sep string, usage string) {
	*p = value
	r.Varer.Var(StringerSliceVar(p, sep, mail.ParseAddress), name, usage)
}

// Time defines a [time.Time] flag with specified name, default value, layout format and usage string.
// The return value is the address of a [time.Time] variable that stores the value of the flag.
func (r Registerer) Time(name string, value time.Time, layout string, usage string) *time.Time {
	r.Varer.Var(TimeVar(&value, layout), name, usage)
	return &value
}

// TimeVar defines a [time.Time] flag with specified name, default value, layout format and usage string.
// The argument p points to a [time.Time] variable in which to store the value of the flag.
func (r Registerer) TimeVar(p *time.Time, name string, value time.Time, layout string, usage string) {
	*p = value
	r.Varer.Var(TimeVar(p, layout), name, usage)
}

// TimeList defines a list-style [time.Time] flag with specified name, default value, layout format and usage string.
// The return value is the address of a [time.Time] slice that stores the values of the flag.
func (r Registerer) TimeList(name string, value []time.Time, layout string, usage string) *[]time.Time {
	r.Varer.Var(TimeListVar(&value, layout), name, usage)
	return &value
}

// TimeListVar defines a list-style [time.Time] flag with specified name, default value, layout format and usage string.
// The argument p points to a [time.Time] slice variable in which to store the value of the flag.
func (r Registerer) TimeListVar(p *[]time.Time, name string, value []time.Time, layout string, usage string) {
	*p = value
	r.Varer.Var(TimeListVar(p, layout), name, usage)
}

// TimeSlice defines a slice-style [time.Time] flag with specified name, default value, layout format and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a [time.Time] slice that stores the values of the flag.
func (r Registerer) TimeSlice(name string, value []time.Time, sep string, layout string, usage string) *[]time.Time {
	r.Varer.Var(TimeSliceVar(&value, sep, layout), name, usage)
	return &value
}

// TimeSliceVar defines a slice-style [time.Time] flag with specified name, default value, layout format and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a [time.Time] slice variable in which to store the value of the flag.
func (r Registerer) TimeSliceVar(p *[]time.Time, name string, value []time.Time, sep string, layout string, usage string) { //nolint: golines
	*p = value
	r.Varer.Var(TimeSliceVar(p, sep, layout), name, usage)
}

// URL defines a [*url.URL] flag with specified name, default value, and usage string.
// The return value is the address of a [*url.URL] variable that stores the value of the flag.
func (r Registerer) URL(name string, value *url.URL, usage string) **url.URL {
	r.Varer.Var(StringerVar(&value, url.Parse), name, usage)
	return &value
}

// URLVar defines a [*url.URL] flag with specified name, default value, and usage string.
// The argument p points to a [*url.URL] variable in which to store the value of the flag.
func (r Registerer) URLVar(p **url.URL, name string, value *url.URL, usage string) {
	*p = value
	r.Varer.Var(StringerVar(p, url.Parse), name, usage)
}

// URLList defines a list-style [*url.URL] flag with specified name, default value, and usage string.
// The return value is the address of a [*url.URL] slice that stores the values of the flag.
func (r Registerer) URLList(name string, value []*url.URL, usage string) *[]*url.URL {
	r.Varer.Var(StringerListVar(&value, url.Parse), name, usage)
	return &value
}

// URLListVar defines a list-style [*url.URL] flag with specified name, default value, and usage string.
// The argument p points to a [*url.URL] slice variable in which to store the value of the flag.
func (r Registerer) URLListVar(p *[]*url.URL, name string, value []*url.URL, usage string) {
	*p = value
	r.Varer.Var(StringerListVar(p, url.Parse), name, usage)
}

// URLSlice defines a slice-style [*url.URL] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a [*url.URL] slice that stores the values of the flag.
func (r Registerer) URLSlice(name string, value []*url.URL, sep string, usage string) *[]*url.URL {
	r.Varer.Var(StringerSliceVar(&value, sep, url.Parse), name, usage)
	return &value
}

// URLSliceVar defines a slice-style [*url.URL] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a [*url.URL] slice variable in which to store the value of the flag.
func (r Registerer) URLSliceVar(p *[]*url.URL, name string, value []*url.URL, sep string, usage string) {
	*p = value
	r.Varer.Var(StringerSliceVar(p, sep, url.Parse), name, usage)
}
