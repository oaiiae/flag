package values

import (
	"flag"
	"fmt"
)

func formatStringer[T fmt.Stringer](t T) string { return t.String() }

func Stringer[T fmt.Stringer](parse func(string) (T, error)) flag.Value {
	return Generic(parse, formatStringer[T])
}

func StringerVar[T fmt.Stringer](p *T, parse func(string) (T, error)) flag.Value {
	return GenericVar(p, parse, formatStringer[T])
}

func StringerList[T fmt.Stringer](parse func(string) (T, error)) flag.Value {
	return GenericList(parse, formatStringer[T])
}

func StringerListVar[T fmt.Stringer](p *[]T, parse func(string) (T, error)) flag.Value {
	return GenericListVar(p, parse, formatStringer[T])
}

func StringerSlice[T fmt.Stringer](sep string, parse func(string) (T, error)) flag.Value {
	return GenericSlice(sep, parse, formatStringer[T])
}

func StringerSliceVar[T fmt.Stringer](p *[]T, sep string, parse func(string) (T, error)) flag.Value {
	return GenericSliceVar(p, sep, parse, formatStringer[T])
}
