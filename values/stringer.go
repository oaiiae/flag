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

func Stringers[T fmt.Stringer](parse func(string) (T, error), split func(string) []string) flag.Value {
	return Generics(parse, formatStringer[T], split)
}

func StringersVar[T fmt.Stringer](p *[]T, parse func(string) (T, error), split func(string) []string) flag.Value {
	return GenericsVar(p, parse, formatStringer[T], split)
}
