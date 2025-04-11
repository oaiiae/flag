package values

import (
	"flag"
	"fmt"
)

func formatStringer[T fmt.Stringer](t T) string { return t.String() }

func Stringer[T fmt.Stringer](fs *flag.FlagSet, name, usage string, parse func(string) (T, error)) *T {
	return Generic(fs, name, usage, parse, formatStringer[T])
}

func StringerVar[T fmt.Stringer](fs *flag.FlagSet, p *T, name, usage string, parse func(string) (T, error)) {
	GenericVar(fs, p, name, usage, parse, formatStringer[T])
}

func Stringers[T fmt.Stringer](fs *flag.FlagSet, name, usage string, parse func(string) (T, error)) *[]T {
	return Generics(fs, name, usage, parse, formatStringer[T])
}

func StringersVar[T fmt.Stringer](fs *flag.FlagSet, p *[]T, name, usage string, parse func(string) (T, error)) {
	GenericsVar(fs, p, name, usage, parse, formatStringer[T])
}
