package values

import (
	"flag"
	"fmt"
)

func formatStringer[T fmt.Stringer](t T) string { return t.String() }

// Stringer declares a [flag.Value] implemented using the parse function and String method of a [fmt.Stringer].
// The actual value type is T.
func Stringer[T fmt.Stringer](parse func(string) (T, error)) flag.Value {
	return Generic(parse, formatStringer[T])
}

// StringerVar is like [Stringer] but store the value in p.
func StringerVar[T fmt.Stringer](p *T, parse func(string) (T, error)) flag.Value {
	return GenericVar(p, parse, formatStringer[T])
}

// StringerList declares a list-style [flag.Value] implemented using the parse function and String method of a [fmt.Stringer].
// The actual value type is []T.
func StringerList[T fmt.Stringer](parse func(string) (T, error)) flag.Value {
	return GenericList(parse, formatStringer[T])
}

// StringerListVar is like [StringerList] but stores the values in p.
func StringerListVar[T fmt.Stringer](p *[]T, parse func(string) (T, error)) flag.Value {
	return GenericListVar(p, parse, formatStringer[T])
}

// StringerSlice declares a slice-style [flag.Value] implemented using the parse function and String method of a [fmt.Stringer].
// The input strings are split around sep before parsing.
// The actual value type is []T.
func StringerSlice[T fmt.Stringer](sep string, parse func(string) (T, error)) flag.Value {
	return GenericSlice(sep, parse, formatStringer[T])
}

// StringerSliceVar is like [StringerSlice] but stores the values in p.
func StringerSliceVar[T fmt.Stringer](p *[]T, sep string, parse func(string) (T, error)) flag.Value {
	return GenericSliceVar(p, sep, parse, formatStringer[T])
}
