package values

import (
	"flag"
	"fmt"
	"strconv"
)

type basic = interface {
	bool |
		complex64 | complex128 |
		int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 |
		string | []byte
}

func parseBasic[T basic](s string) (T, error) { //nolint: funlen // long but still readable
	var v T
	var err error

	switch p := any(&v).(type) {
	case *bool:
		*p, err = strconv.ParseBool(s)

	case *complex64:
		var tmp complex128
		tmp, err = strconv.ParseComplex(s, 64)
		*p = complex64(tmp)

	case *complex128:
		*p, err = strconv.ParseComplex(s, 128)

	case *int:
		var i int64
		i, err = strconv.ParseInt(s, 0, 0)
		*p = int(i)

	case *int8:
		var i int64
		i, err = strconv.ParseInt(s, 0, 8)
		*p = int8(i)

	case *int16:
		var i int64
		i, err = strconv.ParseInt(s, 0, 16)
		*p = int16(i)

	case *int32:
		var i int64
		i, err = strconv.ParseInt(s, 0, 32)
		*p = int32(i)

	case *int64:
		*p, err = strconv.ParseInt(s, 0, 64)

	case *uint:
		var u uint64
		u, err = strconv.ParseUint(s, 0, 0)
		*p = uint(u)

	case *uint8:
		var u uint64
		u, err = strconv.ParseUint(s, 0, 8)
		*p = uint8(u)

	case *uint16:
		var u uint64
		u, err = strconv.ParseUint(s, 0, 16)
		*p = uint16(u)

	case *uint32:
		var u uint64
		u, err = strconv.ParseUint(s, 0, 32)
		*p = uint32(u)

	case *uint64:
		*p, err = strconv.ParseUint(s, 0, 64)

	case *float32:
		var f float64
		f, err = strconv.ParseFloat(s, 32)
		*p = float32(f)

	case *float64:
		*p, err = strconv.ParseFloat(s, 64)

	case *string:
		*p = s

	case *[]byte:
		*p = []byte(s)

	default:
		panic("values: unsupported type")
	}

	return v, err
}

func formatBasic[T basic](v T) string {
	return fmt.Sprint(v)
}

func Basic[T basic]() flag.Value {
	return Generic[T](parseBasic, formatBasic)
}

func BasicVar[T basic](p *T) flag.Value {
	return GenericVar(p, parseBasic, formatBasic)
}

func BasicList[T basic]() flag.Value {
	return GenericList[T](parseBasic, formatBasic)
}

func BasicListVar[T basic](p *[]T) flag.Value {
	return GenericListVar(p, parseBasic, formatBasic)
}

func BasicSlice[T basic](sep string) flag.Value {
	return GenericSlice[T](sep, parseBasic, formatBasic)
}

func BasicSliceVar[T basic](p *[]T, sep string) flag.Value {
	return GenericSliceVar(p, sep, parseBasic, formatBasic)
}
