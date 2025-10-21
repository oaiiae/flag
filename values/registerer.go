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

// RegistererFunc is a function that registers a named [flag.Value] with a usage string.
// It provides an convenient interface analogous to [*flag.FlagSet] for defining flags for common types.
type RegistererFunc func(value flag.Value, name string, usage string)

// FlagSetRegisterer returns a [RegistererFunc] that registers named flags in a [flag.FlagSet].
func FlagSetRegisterer(fs *flag.FlagSet) RegistererFunc { return fs.Var }

// FlagSetEnvRegisterer returns a [RegistererFunc] that registers named flags in a [flag.FlagSet]
// and sets the [flag.Value] with the matching environment variable, ajusting usage accordingly.
//
// The environment variable name is derived from the flag name by:
//
//   - replacing dashes and dots with underscores
//   - transforming to upper case
//   - prepending with given prefix
//
// The environment variable is ignored if it fails to set the flag value.
func FlagSetEnvRegisterer(fs *flag.FlagSet, prefix string) RegistererFunc {
	replacer := strings.NewReplacer("-", "_", ".", "_")
	return func(value flag.Value, name, usage string) {
		envname := prefix + strings.ToUpper(replacer.Replace(name))
		fs.Var(value, name, fmt.Sprintf("%s (env $%s)", usage, envname))
		if val, ok := os.LookupEnv(envname); ok {
			value.Set(val) //nolint: errcheck,gosec // ignore environment then
		}
	}
}

// Bool defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func (f RegistererFunc) Bool(name string, value bool, usage string) *bool {
	f(BasicVar(&value), name, usage)
	return &value
}

// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func (f RegistererFunc) BoolVar(p *bool, name string, value bool, usage string) {
	*p = value
	f(BasicVar(p), name, usage)
}

// BoolList defines a list-style bool flag with specified name, default value, and usage string.
// The return value is the address of a bool slice that stores the values of the flag.
func (f RegistererFunc) BoolList(name string, value []bool, usage string) *[]bool {
	f(BasicListVar(&value), name, usage)
	return &value
}

// BoolListVar defines a list-style bool flag with specified name, default value, and usage string.
// The argument p points to a bool slice variable in which to store the value of the flag.
func (f RegistererFunc) BoolListVar(p *[]bool, name string, value []bool, usage string) {
	*p = value
	f(BasicListVar(p), name, usage)
}

// BoolSlice defines a slice-style bool flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a bool slice that stores the values of the flag.
func (f RegistererFunc) BoolSlice(name string, value []bool, sep string, usage string) *[]bool {
	f(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// BoolSliceVar defines a slice-style bool flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a bool slice variable in which to store the value of the flag.
func (f RegistererFunc) BoolSliceVar(p *[]bool, name string, value []bool, sep string, usage string) {
	*p = value
	f(BasicSliceVar(p, sep), name, usage)
}

// Complex64 defines a complex64 flag with specified name, default value, and usage string.
// The return value is the address of a complex64 variable that stores the value of the flag.
func (f RegistererFunc) Complex64(name string, value complex64, usage string) *complex64 {
	f(BasicVar(&value), name, usage)
	return &value
}

// Complex64Var defines a complex64 flag with specified name, default value, and usage string.
// The argument p points to a complex64 variable in which to store the value of the flag.
func (f RegistererFunc) Complex64Var(p *complex64, name string, value complex64, usage string) {
	*p = value
	f(BasicVar(p), name, usage)
}

// Complex64List defines a list-style complex64 flag with specified name, default value, and usage string.
// The return value is the address of a complex64 slice that stores the values of the flag.
func (f RegistererFunc) Complex64List(name string, value []complex64, usage string) *[]complex64 {
	f(BasicListVar(&value), name, usage)
	return &value
}

// Complex64ListVar defines a list-style complex64 flag with specified name, default value, and usage string.
// The argument p points to a complex64 slice variable in which to store the value of the flag.
func (f RegistererFunc) Complex64ListVar(p *[]complex64, name string, value []complex64, usage string) {
	*p = value
	f(BasicListVar(p), name, usage)
}

// Complex64Slice defines a slice-style complex64 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a complex64 slice that stores the values of the flag.
func (f RegistererFunc) Complex64Slice(name string, value []complex64, sep string, usage string) *[]complex64 {
	f(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Complex64SliceVar defines a slice-style complex64 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a complex64 slice variable in which to store the value of the flag.
func (f RegistererFunc) Complex64SliceVar(p *[]complex64, name string, value []complex64, sep string, usage string) {
	*p = value
	f(BasicSliceVar(p, sep), name, usage)
}

// Complex128 defines a complex128 flag with specified name, default value, and usage string.
// The return value is the address of a complex128 variable that stores the value of the flag.
func (f RegistererFunc) Complex128(name string, value complex128, usage string) *complex128 {
	f(BasicVar(&value), name, usage)
	return &value
}

// Complex128Var defines a complex128 flag with specified name, default value, and usage string.
// The argument p points to a complex128 variable in which to store the value of the flag.
func (f RegistererFunc) Complex128Var(p *complex128, name string, value complex128, usage string) {
	*p = value
	f(BasicVar(p), name, usage)
}

// Complex128List defines a list-style complex128 flag with specified name, default value, and usage string.
// The return value is the address of a complex128 slice that stores the values of the flag.
func (f RegistererFunc) Complex128List(name string, value []complex128, usage string) *[]complex128 {
	f(BasicListVar(&value), name, usage)
	return &value
}

// Complex128ListVar defines a list-style complex128 flag with specified name, default value, and usage string.
// The argument p points to a complex128 slice variable in which to store the value of the flag.
func (f RegistererFunc) Complex128ListVar(p *[]complex128, name string, value []complex128, usage string) {
	*p = value
	f(BasicListVar(p), name, usage)
}

// Complex128Slice defines a slice-style complex128 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a complex128 slice that stores the values of the flag.
func (f RegistererFunc) Complex128Slice(name string, value []complex128, sep string, usage string) *[]complex128 {
	f(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Complex128SliceVar defines a slice-style complex128 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a complex128 slice variable in which to store the value of the flag.
func (f RegistererFunc) Complex128SliceVar(p *[]complex128, name string, value []complex128, sep string, usage string) {
	*p = value
	f(BasicSliceVar(p, sep), name, usage)
}

// Int defines a int flag with specified name, default value, and usage string.
// The return value is the address of a int variable that stores the value of the flag.
func (f RegistererFunc) Int(name string, value int, usage string) *int {
	f(BasicVar(&value), name, usage)
	return &value
}

// IntVar defines a int flag with specified name, default value, and usage string.
// The argument p points to a int variable in which to store the value of the flag.
func (f RegistererFunc) IntVar(p *int, name string, value int, usage string) {
	*p = value
	f(BasicVar(p), name, usage)
}

// IntList defines a list-style int flag with specified name, default value, and usage string.
// The return value is the address of a int slice that stores the values of the flag.
func (f RegistererFunc) IntList(name string, value []int, usage string) *[]int {
	f(BasicListVar(&value), name, usage)
	return &value
}

// IntListVar defines a list-style int flag with specified name, default value, and usage string.
// The argument p points to a int slice variable in which to store the value of the flag.
func (f RegistererFunc) IntListVar(p *[]int, name string, value []int, usage string) {
	*p = value
	f(BasicListVar(p), name, usage)
}

// IntSlice defines a slice-style int flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a int slice that stores the values of the flag.
func (f RegistererFunc) IntSlice(name string, value []int, sep string, usage string) *[]int {
	f(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// IntSliceVar defines a slice-style int flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a int slice variable in which to store the value of the flag.
func (f RegistererFunc) IntSliceVar(p *[]int, name string, value []int, sep string, usage string) {
	*p = value
	f(BasicSliceVar(p, sep), name, usage)
}

// Int8 defines a int8 flag with specified name, default value, and usage string.
// The return value is the address of a int8 variable that stores the value of the flag.
func (f RegistererFunc) Int8(name string, value int8, usage string) *int8 {
	f(BasicVar(&value), name, usage)
	return &value
}

// Int8Var defines a int8 flag with specified name, default value, and usage string.
// The argument p points to a int8 variable in which to store the value of the flag.
func (f RegistererFunc) Int8Var(p *int8, name string, value int8, usage string) {
	*p = value
	f(BasicVar(p), name, usage)
}

// Int8List defines a list-style int8 flag with specified name, default value, and usage string.
// The return value is the address of a int8 slice that stores the values of the flag.
func (f RegistererFunc) Int8List(name string, value []int8, usage string) *[]int8 {
	f(BasicListVar(&value), name, usage)
	return &value
}

// Int8ListVar defines a list-style int8 flag with specified name, default value, and usage string.
// The argument p points to a int8 slice variable in which to store the value of the flag.
func (f RegistererFunc) Int8ListVar(p *[]int8, name string, value []int8, usage string) {
	*p = value
	f(BasicListVar(p), name, usage)
}

// Int8Slice defines a slice-style int8 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a int8 slice that stores the values of the flag.
func (f RegistererFunc) Int8Slice(name string, value []int8, sep string, usage string) *[]int8 {
	f(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Int8SliceVar defines a slice-style int8 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a int8 slice variable in which to store the value of the flag.
func (f RegistererFunc) Int8SliceVar(p *[]int8, name string, value []int8, sep string, usage string) {
	*p = value
	f(BasicSliceVar(p, sep), name, usage)
}

// Int16 defines a int16 flag with specified name, default value, and usage string.
// The return value is the address of a int16 variable that stores the value of the flag.
func (f RegistererFunc) Int16(name string, value int16, usage string) *int16 {
	f(BasicVar(&value), name, usage)
	return &value
}

// Int16Var defines a int16 flag with specified name, default value, and usage string.
// The argument p points to a int16 variable in which to store the value of the flag.
func (f RegistererFunc) Int16Var(p *int16, name string, value int16, usage string) {
	*p = value
	f(BasicVar(p), name, usage)
}

// Int16List defines a list-style int16 flag with specified name, default value, and usage string.
// The return value is the address of a int16 slice that stores the values of the flag.
func (f RegistererFunc) Int16List(name string, value []int16, usage string) *[]int16 {
	f(BasicListVar(&value), name, usage)
	return &value
}

// Int16ListVar defines a list-style int16 flag with specified name, default value, and usage string.
// The argument p points to a int16 slice variable in which to store the value of the flag.
func (f RegistererFunc) Int16ListVar(p *[]int16, name string, value []int16, usage string) {
	*p = value
	f(BasicListVar(p), name, usage)
}

// Int16Slice defines a slice-style int16 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a int16 slice that stores the values of the flag.
func (f RegistererFunc) Int16Slice(name string, value []int16, sep string, usage string) *[]int16 {
	f(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Int16SliceVar defines a slice-style int16 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a int16 slice variable in which to store the value of the flag.
func (f RegistererFunc) Int16SliceVar(p *[]int16, name string, value []int16, sep string, usage string) {
	*p = value
	f(BasicSliceVar(p, sep), name, usage)
}

// Int32 defines a int32 flag with specified name, default value, and usage string.
// The return value is the address of a int32 variable that stores the value of the flag.
func (f RegistererFunc) Int32(name string, value int32, usage string) *int32 {
	f(BasicVar(&value), name, usage)
	return &value
}

// Int32Var defines a int32 flag with specified name, default value, and usage string.
// The argument p points to a int32 variable in which to store the value of the flag.
func (f RegistererFunc) Int32Var(p *int32, name string, value int32, usage string) {
	*p = value
	f(BasicVar(p), name, usage)
}

// Int32List defines a list-style int32 flag with specified name, default value, and usage string.
// The return value is the address of a int32 slice that stores the values of the flag.
func (f RegistererFunc) Int32List(name string, value []int32, usage string) *[]int32 {
	f(BasicListVar(&value), name, usage)
	return &value
}

// Int32ListVar defines a list-style int32 flag with specified name, default value, and usage string.
// The argument p points to a int32 slice variable in which to store the value of the flag.
func (f RegistererFunc) Int32ListVar(p *[]int32, name string, value []int32, usage string) {
	*p = value
	f(BasicListVar(p), name, usage)
}

// Int32Slice defines a slice-style int32 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a int32 slice that stores the values of the flag.
func (f RegistererFunc) Int32Slice(name string, value []int32, sep string, usage string) *[]int32 {
	f(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Int32SliceVar defines a slice-style int32 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a int32 slice variable in which to store the value of the flag.
func (f RegistererFunc) Int32SliceVar(p *[]int32, name string, value []int32, sep string, usage string) {
	*p = value
	f(BasicSliceVar(p, sep), name, usage)
}

// Int64 defines a int64 flag with specified name, default value, and usage string.
// The return value is the address of a int64 variable that stores the value of the flag.
func (f RegistererFunc) Int64(name string, value int64, usage string) *int64 {
	f(BasicVar(&value), name, usage)
	return &value
}

// Int64Var defines a int64 flag with specified name, default value, and usage string.
// The argument p points to a int64 variable in which to store the value of the flag.
func (f RegistererFunc) Int64Var(p *int64, name string, value int64, usage string) {
	*p = value
	f(BasicVar(p), name, usage)
}

// Int64List defines a list-style int64 flag with specified name, default value, and usage string.
// The return value is the address of a int64 slice that stores the values of the flag.
func (f RegistererFunc) Int64List(name string, value []int64, usage string) *[]int64 {
	f(BasicListVar(&value), name, usage)
	return &value
}

// Int64ListVar defines a list-style int64 flag with specified name, default value, and usage string.
// The argument p points to a int64 slice variable in which to store the value of the flag.
func (f RegistererFunc) Int64ListVar(p *[]int64, name string, value []int64, usage string) {
	*p = value
	f(BasicListVar(p), name, usage)
}

// Int64Slice defines a slice-style int64 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a int64 slice that stores the values of the flag.
func (f RegistererFunc) Int64Slice(name string, value []int64, sep string, usage string) *[]int64 {
	f(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Int64SliceVar defines a slice-style int64 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a int64 slice variable in which to store the value of the flag.
func (f RegistererFunc) Int64SliceVar(p *[]int64, name string, value []int64, sep string, usage string) {
	*p = value
	f(BasicSliceVar(p, sep), name, usage)
}

// Uint defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint variable that stores the value of the flag.
func (f RegistererFunc) Uint(name string, value uint, usage string) *uint {
	f(BasicVar(&value), name, usage)
	return &value
}

// UintVar defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func (f RegistererFunc) UintVar(p *uint, name string, value uint, usage string) {
	*p = value
	f(BasicVar(p), name, usage)
}

// UintList defines a list-style uint flag with specified name, default value, and usage string.
// The return value is the address of a uint slice that stores the values of the flag.
func (f RegistererFunc) UintList(name string, value []uint, usage string) *[]uint {
	f(BasicListVar(&value), name, usage)
	return &value
}

// UintListVar defines a list-style uint flag with specified name, default value, and usage string.
// The argument p points to a uint slice variable in which to store the value of the flag.
func (f RegistererFunc) UintListVar(p *[]uint, name string, value []uint, usage string) {
	*p = value
	f(BasicListVar(p), name, usage)
}

// UintSlice defines a slice-style uint flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a uint slice that stores the values of the flag.
func (f RegistererFunc) UintSlice(name string, value []uint, sep string, usage string) *[]uint {
	f(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// UintSliceVar defines a slice-style uint flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a uint slice variable in which to store the value of the flag.
func (f RegistererFunc) UintSliceVar(p *[]uint, name string, value []uint, sep string, usage string) {
	*p = value
	f(BasicSliceVar(p, sep), name, usage)
}

// Uint8 defines a uint8 flag with specified name, default value, and usage string.
// The return value is the address of a uint8 variable that stores the value of the flag.
func (f RegistererFunc) Uint8(name string, value uint8, usage string) *uint8 {
	f(BasicVar(&value), name, usage)
	return &value
}

// Uint8Var defines a uint8 flag with specified name, default value, and usage string.
// The argument p points to a uint8 variable in which to store the value of the flag.
func (f RegistererFunc) Uint8Var(p *uint8, name string, value uint8, usage string) {
	*p = value
	f(BasicVar(p), name, usage)
}

// Uint8List defines a list-style uint8 flag with specified name, default value, and usage string.
// The return value is the address of a uint8 slice that stores the values of the flag.
func (f RegistererFunc) Uint8List(name string, value []uint8, usage string) *[]uint8 {
	f(BasicListVar(&value), name, usage)
	return &value
}

// Uint8ListVar defines a list-style uint8 flag with specified name, default value, and usage string.
// The argument p points to a uint8 slice variable in which to store the value of the flag.
func (f RegistererFunc) Uint8ListVar(p *[]uint8, name string, value []uint8, usage string) {
	*p = value
	f(BasicListVar(p), name, usage)
}

// Uint8Slice defines a slice-style uint8 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a uint8 slice that stores the values of the flag.
func (f RegistererFunc) Uint8Slice(name string, value []uint8, sep string, usage string) *[]uint8 {
	f(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Uint8SliceVar defines a slice-style uint8 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a uint8 slice variable in which to store the value of the flag.
func (f RegistererFunc) Uint8SliceVar(p *[]uint8, name string, value []uint8, sep string, usage string) {
	*p = value
	f(BasicSliceVar(p, sep), name, usage)
}

// Uint16 defines a uint16 flag with specified name, default value, and usage string.
// The return value is the address of a uint16 variable that stores the value of the flag.
func (f RegistererFunc) Uint16(name string, value uint16, usage string) *uint16 {
	f(BasicVar(&value), name, usage)
	return &value
}

// Uint16Var defines a uint16 flag with specified name, default value, and usage string.
// The argument p points to a uint16 variable in which to store the value of the flag.
func (f RegistererFunc) Uint16Var(p *uint16, name string, value uint16, usage string) {
	*p = value
	f(BasicVar(p), name, usage)
}

// Uint16List defines a list-style uint16 flag with specified name, default value, and usage string.
// The return value is the address of a uint16 slice that stores the values of the flag.
func (f RegistererFunc) Uint16List(name string, value []uint16, usage string) *[]uint16 {
	f(BasicListVar(&value), name, usage)
	return &value
}

// Uint16ListVar defines a list-style uint16 flag with specified name, default value, and usage string.
// The argument p points to a uint16 slice variable in which to store the value of the flag.
func (f RegistererFunc) Uint16ListVar(p *[]uint16, name string, value []uint16, usage string) {
	*p = value
	f(BasicListVar(p), name, usage)
}

// Uint16Slice defines a slice-style uint16 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a uint16 slice that stores the values of the flag.
func (f RegistererFunc) Uint16Slice(name string, value []uint16, sep string, usage string) *[]uint16 {
	f(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Uint16SliceVar defines a slice-style uint16 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a uint16 slice variable in which to store the value of the flag.
func (f RegistererFunc) Uint16SliceVar(p *[]uint16, name string, value []uint16, sep string, usage string) {
	*p = value
	f(BasicSliceVar(p, sep), name, usage)
}

// Uint32 defines a uint32 flag with specified name, default value, and usage string.
// The return value is the address of a uint32 variable that stores the value of the flag.
func (f RegistererFunc) Uint32(name string, value uint32, usage string) *uint32 {
	f(BasicVar(&value), name, usage)
	return &value
}

// Uint32Var defines a uint32 flag with specified name, default value, and usage string.
// The argument p points to a uint32 variable in which to store the value of the flag.
func (f RegistererFunc) Uint32Var(p *uint32, name string, value uint32, usage string) {
	*p = value
	f(BasicVar(p), name, usage)
}

// Uint32List defines a list-style uint32 flag with specified name, default value, and usage string.
// The return value is the address of a uint32 slice that stores the values of the flag.
func (f RegistererFunc) Uint32List(name string, value []uint32, usage string) *[]uint32 {
	f(BasicListVar(&value), name, usage)
	return &value
}

// Uint32ListVar defines a list-style uint32 flag with specified name, default value, and usage string.
// The argument p points to a uint32 slice variable in which to store the value of the flag.
func (f RegistererFunc) Uint32ListVar(p *[]uint32, name string, value []uint32, usage string) {
	*p = value
	f(BasicListVar(p), name, usage)
}

// Uint32Slice defines a slice-style uint32 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a uint32 slice that stores the values of the flag.
func (f RegistererFunc) Uint32Slice(name string, value []uint32, sep string, usage string) *[]uint32 {
	f(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Uint32SliceVar defines a slice-style uint32 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a uint32 slice variable in which to store the value of the flag.
func (f RegistererFunc) Uint32SliceVar(p *[]uint32, name string, value []uint32, sep string, usage string) {
	*p = value
	f(BasicSliceVar(p, sep), name, usage)
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
func (f RegistererFunc) Uint64(name string, value uint64, usage string) *uint64 {
	f(BasicVar(&value), name, usage)
	return &value
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func (f RegistererFunc) Uint64Var(p *uint64, name string, value uint64, usage string) {
	*p = value
	f(BasicVar(p), name, usage)
}

// Uint64List defines a list-style uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 slice that stores the values of the flag.
func (f RegistererFunc) Uint64List(name string, value []uint64, usage string) *[]uint64 {
	f(BasicListVar(&value), name, usage)
	return &value
}

// Uint64ListVar defines a list-style uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 slice variable in which to store the value of the flag.
func (f RegistererFunc) Uint64ListVar(p *[]uint64, name string, value []uint64, usage string) {
	*p = value
	f(BasicListVar(p), name, usage)
}

// Uint64Slice defines a slice-style uint64 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a uint64 slice that stores the values of the flag.
func (f RegistererFunc) Uint64Slice(name string, value []uint64, sep string, usage string) *[]uint64 {
	f(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Uint64SliceVar defines a slice-style uint64 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a uint64 slice variable in which to store the value of the flag.
func (f RegistererFunc) Uint64SliceVar(p *[]uint64, name string, value []uint64, sep string, usage string) {
	*p = value
	f(BasicSliceVar(p, sep), name, usage)
}

// Float32 defines a float32 flag with specified name, default value, and usage string.
// The return value is the address of a float32 variable that stores the value of the flag.
func (f RegistererFunc) Float32(name string, value float32, usage string) *float32 {
	f(BasicVar(&value), name, usage)
	return &value
}

// Float32Var defines a float32 flag with specified name, default value, and usage string.
// The argument p points to a float32 variable in which to store the value of the flag.
func (f RegistererFunc) Float32Var(p *float32, name string, value float32, usage string) {
	*p = value
	f(BasicVar(p), name, usage)
}

// Float32List defines a list-style float32 flag with specified name, default value, and usage string.
// The return value is the address of a float32 slice that stores the values of the flag.
func (f RegistererFunc) Float32List(name string, value []float32, usage string) *[]float32 {
	f(BasicListVar(&value), name, usage)
	return &value
}

// Float32ListVar defines a list-style float32 flag with specified name, default value, and usage string.
// The argument p points to a float32 slice variable in which to store the value of the flag.
func (f RegistererFunc) Float32ListVar(p *[]float32, name string, value []float32, usage string) {
	*p = value
	f(BasicListVar(p), name, usage)
}

// Float32Slice defines a slice-style float32 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a float32 slice that stores the values of the flag.
func (f RegistererFunc) Float32Slice(name string, value []float32, sep string, usage string) *[]float32 {
	f(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Float32SliceVar defines a slice-style float32 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a float32 slice variable in which to store the value of the flag.
func (f RegistererFunc) Float32SliceVar(p *[]float32, name string, value []float32, sep string, usage string) {
	*p = value
	f(BasicSliceVar(p, sep), name, usage)
}

// Float64 defines a float64 flag with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag.
func (f RegistererFunc) Float64(name string, value float64, usage string) *float64 {
	f(BasicVar(&value), name, usage)
	return &value
}

// Float64Var defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
func (f RegistererFunc) Float64Var(p *float64, name string, value float64, usage string) {
	*p = value
	f(BasicVar(p), name, usage)
}

// Float64List defines a list-style float64 flag with specified name, default value, and usage string.
// The return value is the address of a float64 slice that stores the values of the flag.
func (f RegistererFunc) Float64List(name string, value []float64, usage string) *[]float64 {
	f(BasicListVar(&value), name, usage)
	return &value
}

// Float64ListVar defines a list-style float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 slice variable in which to store the value of the flag.
func (f RegistererFunc) Float64ListVar(p *[]float64, name string, value []float64, usage string) {
	*p = value
	f(BasicListVar(p), name, usage)
}

// Float64Slice defines a slice-style float64 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a float64 slice that stores the values of the flag.
func (f RegistererFunc) Float64Slice(name string, value []float64, sep string, usage string) *[]float64 {
	f(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// Float64SliceVar defines a slice-style float64 flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a float64 slice variable in which to store the value of the flag.
func (f RegistererFunc) Float64SliceVar(p *[]float64, name string, value []float64, sep string, usage string) {
	*p = value
	f(BasicSliceVar(p, sep), name, usage)
}

// String defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func (f RegistererFunc) String(name string, value string, usage string) *string {
	f(BasicVar(&value), name, usage)
	return &value
}

// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func (f RegistererFunc) StringVar(p *string, name string, value string, usage string) {
	*p = value
	f(BasicVar(p), name, usage)
}

// StringList defines a list-style string flag with specified name, default value, and usage string.
// The return value is the address of a string slice that stores the values of the flag.
func (f RegistererFunc) StringList(name string, value []string, usage string) *[]string {
	f(BasicListVar(&value), name, usage)
	return &value
}

// StringListVar defines a list-style string flag with specified name, default value, and usage string.
// The argument p points to a string slice variable in which to store the value of the flag.
func (f RegistererFunc) StringListVar(p *[]string, name string, value []string, usage string) {
	*p = value
	f(BasicListVar(p), name, usage)
}

// StringSlice defines a slice-style string flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a string slice that stores the values of the flag.
func (f RegistererFunc) StringSlice(name string, value []string, sep string, usage string) *[]string {
	f(BasicSliceVar(&value, sep), name, usage)
	return &value
}

// StringSliceVar defines a slice-style string flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a string slice variable in which to store the value of the flag.
func (f RegistererFunc) StringSliceVar(p *[]string, name string, value []string, sep string, usage string) {
	*p = value
	f(BasicSliceVar(p, sep), name, usage)
}

// Duration defines a [time.Duration] flag with specified name, default value, and usage string.
// The return value is the address of a [time.Duration] variable that stores the value of the flag.
func (f RegistererFunc) Duration(name string, value time.Duration, usage string) *time.Duration {
	f(DurationVar(&value), name, usage)
	return &value
}

// DurationVar defines a [time.Duration] flag with specified name, default value, and usage string.
// The argument p points to a [time.Duration] variable in which to store the value of the flag.
func (f RegistererFunc) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	*p = value
	f(DurationVar(p), name, usage)
}

// DurationList defines a list-style [time.Duration] flag with specified name, default value, and usage string.
// The return value is the address of a [time.Duration] slice that stores the values of the flag.
func (f RegistererFunc) DurationList(name string, value []time.Duration, usage string) *[]time.Duration {
	f(DurationListVar(&value), name, usage)
	return &value
}

// DurationListVar defines a list-style [time.Duration] flag with specified name, default value, and usage string.
// The argument p points to a [time.Duration] slice variable in which to store the value of the flag.
func (f RegistererFunc) DurationListVar(p *[]time.Duration, name string, value []time.Duration, usage string) {
	*p = value
	f(DurationListVar(p), name, usage)
}

// DurationSlice defines a slice-style [time.Duration] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a [time.Duration] slice that stores the values of the flag.
func (f RegistererFunc) DurationSlice(name string, value []time.Duration, sep string, usage string) *[]time.Duration {
	f(DurationSliceVar(&value, sep), name, usage)
	return &value
}

// DurationSliceVar defines a slice-style [time.Duration] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a [time.Duration] slice variable in which to store the value of the flag.
func (f RegistererFunc) DurationSliceVar(p *[]time.Duration, name string, value []time.Duration, sep string, usage string) {
	*p = value
	f(DurationSliceVar(p, sep), name, usage)
}

// IPAddr defines a [netip.Addr] flag with specified name, default value, and usage string.
// The return value is the address of a [netip.Addr] variable that stores the value of the flag.
func (f RegistererFunc) IPAddr(name string, value netip.Addr, usage string) *netip.Addr {
	f(StringerVar(&value, netip.ParseAddr), name, usage)
	return &value
}

// IPAddrVar defines a [netip.Addr] flag with specified name, default value, and usage string.
// The argument p points to a [netip.Addr] variable in which to store the value of the flag.
func (f RegistererFunc) IPAddrVar(p *netip.Addr, name string, value netip.Addr, usage string) {
	*p = value
	f(StringerVar(p, netip.ParseAddr), name, usage)
}

// IPAddrList defines a list-style [netip.Addr] flag with specified name, default value, and usage string.
// The return value is the address of a [netip.Addr] slice that stores the values of the flag.
func (f RegistererFunc) IPAddrList(name string, value []netip.Addr, usage string) *[]netip.Addr {
	f(StringerListVar(&value, netip.ParseAddr), name, usage)
	return &value
}

// IPAddrListVar defines a list-style [netip.Addr] flag with specified name, default value, and usage string.
// The argument p points to a [netip.Addr] slice variable in which to store the value of the flag.
func (f RegistererFunc) IPAddrListVar(p *[]netip.Addr, name string, value []netip.Addr, usage string) {
	*p = value
	f(StringerListVar(p, netip.ParseAddr), name, usage)
}

// IPAddrSlice defines a slice-style [netip.Addr] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a [netip.Addr] slice that stores the values of the flag.
func (f RegistererFunc) IPAddrSlice(name string, value []netip.Addr, sep string, usage string) *[]netip.Addr {
	f(StringerSliceVar(&value, sep, netip.ParseAddr), name, usage)
	return &value
}

// IPAddrSliceVar defines a slice-style [netip.Addr] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a [netip.Addr] slice variable in which to store the value of the flag.
func (f RegistererFunc) IPAddrSliceVar(p *[]netip.Addr, name string, value []netip.Addr, sep string, usage string) {
	*p = value
	f(StringerSliceVar(p, sep, netip.ParseAddr), name, usage)
}

// IPAddrPort defines a [netip.AddrPort] flag with specified name, default value, and usage string.
// The return value is the address of a [netip.AddrPort] variable that stores the value of the flag.
func (f RegistererFunc) IPAddrPort(name string, value netip.AddrPort, usage string) *netip.AddrPort {
	f(StringerVar(&value, netip.ParseAddrPort), name, usage)
	return &value
}

// IPAddrPortVar defines a [netip.AddrPort] flag with specified name, default value, and usage string.
// The argument p points to a [netip.AddrPort] variable in which to store the value of the flag.
func (f RegistererFunc) IPAddrPortVar(p *netip.AddrPort, name string, value netip.AddrPort, usage string) {
	*p = value
	f(StringerVar(p, netip.ParseAddrPort), name, usage)
}

// IPAddrPortList defines a list-style [netip.AddrPort] flag with specified name, default value, and usage string.
// The return value is the address of a [netip.AddrPort] slice that stores the values of the flag.
func (f RegistererFunc) IPAddrPortList(name string, value []netip.AddrPort, usage string) *[]netip.AddrPort {
	f(StringerListVar(&value, netip.ParseAddrPort), name, usage)
	return &value
}

// IPAddrPortListVar defines a list-style [netip.AddrPort] flag with specified name, default value, and usage string.
// The argument p points to a [netip.AddrPort] slice variable in which to store the value of the flag.
func (f RegistererFunc) IPAddrPortListVar(p *[]netip.AddrPort, name string, value []netip.AddrPort, usage string) {
	*p = value
	f(StringerListVar(p, netip.ParseAddrPort), name, usage)
}

// IPAddrPortSlice defines a slice-style [netip.AddrPort] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a [netip.AddrPort] slice that stores the values of the flag.
func (f RegistererFunc) IPAddrPortSlice(name string, value []netip.AddrPort, sep string, usage string) *[]netip.AddrPort {
	f(StringerSliceVar(&value, sep, netip.ParseAddrPort), name, usage)
	return &value
}

// IPAddrPortSliceVar defines a slice-style [netip.AddrPort] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a [netip.AddrPort] slice variable in which to store the value of the flag.
func (f RegistererFunc) IPAddrPortSliceVar(p *[]netip.AddrPort, name string, value []netip.AddrPort, sep string, usage string) { //nolint: golines
	*p = value
	f(StringerSliceVar(p, sep, netip.ParseAddrPort), name, usage)
}

// IPPrefix defines a [netip.Prefix] flag with specified name, default value, and usage string.
// The return value is the address of a [netip.Prefix] variable that stores the value of the flag.
func (f RegistererFunc) IPPrefix(name string, value netip.Prefix, usage string) *netip.Prefix {
	f(StringerVar(&value, netip.ParsePrefix), name, usage)
	return &value
}

// IPPrefixVar defines a [netip.Prefix] flag with specified name, default value, and usage string.
// The argument p points to a [netip.Prefix] variable in which to store the value of the flag.
func (f RegistererFunc) IPPrefixVar(p *netip.Prefix, name string, value netip.Prefix, usage string) {
	*p = value
	f(StringerVar(p, netip.ParsePrefix), name, usage)
}

// IPPrefixList defines a list-style [netip.Prefix] flag with specified name, default value, and usage string.
// The return value is the address of a [netip.Prefix] slice that stores the values of the flag.
func (f RegistererFunc) IPPrefixList(name string, value []netip.Prefix, usage string) *[]netip.Prefix {
	f(StringerListVar(&value, netip.ParsePrefix), name, usage)
	return &value
}

// IPPrefixListVar defines a list-style [netip.Prefix] flag with specified name, default value, and usage string.
// The argument p points to a [netip.Prefix] slice variable in which to store the value of the flag.
func (f RegistererFunc) IPPrefixListVar(p *[]netip.Prefix, name string, value []netip.Prefix, usage string) {
	*p = value
	f(StringerListVar(p, netip.ParsePrefix), name, usage)
}

// IPPrefixSlice defines a slice-style [netip.Prefix] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a [netip.Prefix] slice that stores the values of the flag.
func (f RegistererFunc) IPPrefixSlice(name string, value []netip.Prefix, sep string, usage string) *[]netip.Prefix {
	f(StringerSliceVar(&value, sep, netip.ParsePrefix), name, usage)
	return &value
}

// IPPrefixSliceVar defines a slice-style [netip.Prefix] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a [netip.Prefix] slice variable in which to store the value of the flag.
func (f RegistererFunc) IPPrefixSliceVar(p *[]netip.Prefix, name string, value []netip.Prefix, sep string, usage string) {
	*p = value
	f(StringerSliceVar(p, sep, netip.ParsePrefix), name, usage)
}

// MailAddr defines a [*mail.Address] flag with specified name, default value, and usage string.
// The return value is the address of a [*mail.Address] variable that stores the value of the flag.
func (f RegistererFunc) MailAddr(name string, value *mail.Address, usage string) **mail.Address {
	f(StringerVar(&value, mail.ParseAddress), name, usage)
	return &value
}

// MailAddrVar defines a [*mail.Address] flag with specified name, default value, and usage string.
// The argument p points to a [*mail.Address] variable in which to store the value of the flag.
func (f RegistererFunc) MailAddrVar(p **mail.Address, name string, value *mail.Address, usage string) {
	*p = value
	f(StringerVar(p, mail.ParseAddress), name, usage)
}

// MailAddrList defines a list-style [*mail.Address] flag with specified name, default value, and usage string.
// The return value is the address of a [*mail.Address] slice that stores the values of the flag.
func (f RegistererFunc) MailAddrList(name string, value []*mail.Address, usage string) *[]*mail.Address {
	f(StringerListVar(&value, mail.ParseAddress), name, usage)
	return &value
}

// MailAddrListVar defines a list-style [*mail.Address] flag with specified name, default value, and usage string.
// The argument p points to a [*mail.Address] slice variable in which to store the value of the flag.
func (f RegistererFunc) MailAddrListVar(p *[]*mail.Address, name string, value []*mail.Address, usage string) {
	*p = value
	f(StringerListVar(p, mail.ParseAddress), name, usage)
}

// MailAddrSlice defines a slice-style [*mail.Address] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a [*mail.Address] slice that stores the values of the flag.
func (f RegistererFunc) MailAddrSlice(name string, value []*mail.Address, sep string, usage string) *[]*mail.Address {
	f(StringerSliceVar(&value, sep, mail.ParseAddress), name, usage)
	return &value
}

// MailAddrSliceVar defines a slice-style [*mail.Address] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a [*mail.Address] slice variable in which to store the value of the flag.
func (f RegistererFunc) MailAddrSliceVar(p *[]*mail.Address, name string, value []*mail.Address, sep string, usage string) {
	*p = value
	f(StringerSliceVar(p, sep, mail.ParseAddress), name, usage)
}

// Time defines a [time.Time] flag with specified name, default value, layout format and usage string.
// The return value is the address of a [time.Time] variable that stores the value of the flag.
func (f RegistererFunc) Time(name string, value time.Time, layout string, usage string) *time.Time {
	f(TimeVar(&value, layout), name, usage)
	return &value
}

// TimeVar defines a [time.Time] flag with specified name, default value, layout format and usage string.
// The argument p points to a [time.Time] variable in which to store the value of the flag.
func (f RegistererFunc) TimeVar(p *time.Time, name string, value time.Time, layout string, usage string) {
	*p = value
	f(TimeVar(p, layout), name, usage)
}

// TimeList defines a list-style [time.Time] flag with specified name, default value, layout format and usage string.
// The return value is the address of a [time.Time] slice that stores the values of the flag.
func (f RegistererFunc) TimeList(name string, value []time.Time, layout string, usage string) *[]time.Time {
	f(TimeListVar(&value, layout), name, usage)
	return &value
}

// TimeListVar defines a list-style [time.Time] flag with specified name, default value, layout format and usage string.
// The argument p points to a [time.Time] slice variable in which to store the value of the flag.
func (f RegistererFunc) TimeListVar(p *[]time.Time, name string, value []time.Time, layout string, usage string) {
	*p = value
	f(TimeListVar(p, layout), name, usage)
}

// TimeSlice defines a slice-style [time.Time] flag with specified name, default value, layout format and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a [time.Time] slice that stores the values of the flag.
func (f RegistererFunc) TimeSlice(name string, value []time.Time, sep string, layout string, usage string) *[]time.Time {
	f(TimeSliceVar(&value, sep, layout), name, usage)
	return &value
}

// TimeSliceVar defines a slice-style [time.Time] flag with specified name, default value, layout format and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a [time.Time] slice variable in which to store the value of the flag.
func (f RegistererFunc) TimeSliceVar(p *[]time.Time, name string, value []time.Time, sep string, layout string, usage string) { //nolint: golines
	*p = value
	f(TimeSliceVar(p, sep, layout), name, usage)
}

// URL defines a [*url.URL] flag with specified name, default value, and usage string.
// The return value is the address of a [*url.URL] variable that stores the value of the flag.
func (f RegistererFunc) URL(name string, value *url.URL, usage string) **url.URL {
	f(StringerVar(&value, url.Parse), name, usage)
	return &value
}

// URLVar defines a [*url.URL] flag with specified name, default value, and usage string.
// The argument p points to a [*url.URL] variable in which to store the value of the flag.
func (f RegistererFunc) URLVar(p **url.URL, name string, value *url.URL, usage string) {
	*p = value
	f(StringerVar(p, url.Parse), name, usage)
}

// URLList defines a list-style [*url.URL] flag with specified name, default value, and usage string.
// The return value is the address of a [*url.URL] slice that stores the values of the flag.
func (f RegistererFunc) URLList(name string, value []*url.URL, usage string) *[]*url.URL {
	f(StringerListVar(&value, url.Parse), name, usage)
	return &value
}

// URLListVar defines a list-style [*url.URL] flag with specified name, default value, and usage string.
// The argument p points to a [*url.URL] slice variable in which to store the value of the flag.
func (f RegistererFunc) URLListVar(p *[]*url.URL, name string, value []*url.URL, usage string) {
	*p = value
	f(StringerListVar(p, url.Parse), name, usage)
}

// URLSlice defines a slice-style [*url.URL] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The return value is the address of a [*url.URL] slice that stores the values of the flag.
func (f RegistererFunc) URLSlice(name string, value []*url.URL, sep string, usage string) *[]*url.URL {
	f(StringerSliceVar(&value, sep, url.Parse), name, usage)
	return &value
}

// URLSliceVar defines a slice-style [*url.URL] flag with specified name, default value, and usage string.
// The input strings are split around sep before parsing.
// The argument p points to a [*url.URL] slice variable in which to store the value of the flag.
func (f RegistererFunc) URLSliceVar(p *[]*url.URL, name string, value []*url.URL, sep string, usage string) {
	*p = value
	f(StringerSliceVar(p, sep, url.Parse), name, usage)
}
