package values

import (
	"flag"
	"fmt"
)

type generic[T any] struct {
	parse  func(string) (T, error)
	format func(T) string
	value  *T
	isset  bool
}

func (v *generic[T]) Set(s string) error {
	val, err := v.parse(s)
	if err != nil {
		return err
	}
	*v.value, v.isset = val, true
	return nil
}

func (v *generic[T]) String() string {
	if v.value == nil || !v.isset {
		return ""
	}
	return v.format(*v.value)
}

func (v *generic[T]) Get() any {
	return *v.value
}

func Generic[T any](fs *flag.FlagSet, name, usage string, parse func(string) (T, error), format func(T) string) *T {
	g := generic[T]{parse, format, new(T), false}
	fs.Var(&g, name, usage)
	return g.value
}

func GenericVar[T any](fs *flag.FlagSet, p *T, name, usage string, parse func(string) (T, error), format func(T) string) {
	g := generic[T]{parse, format, p, true}
	fs.Var(&g, name, usage)
}

type generics[T any] struct {
	parse  func(string) (T, error)
	format func(T) string
	values *[]T
}

func (v *generics[T]) Set(s string) error {
	val, err := v.parse(s)
	if err != nil {
		return err
	}
	*v.values = append(*v.values, val)
	return nil
}

func (v *generics[T]) String() string {
	if v.values == nil || len(*v.values) == 0 {
		return ""
	}
	a := make([]string, 0, len(*v.values))
	for i := range *v.values {
		a = append(a, v.format((*v.values)[i]))
	}
	return fmt.Sprint(a)
}

func (v *generics[T]) Get() any {
	return *v.values
}

func Generics[T any](fs *flag.FlagSet, name, usage string, parse func(string) (T, error), format func(T) string) *[]T {
	g := generics[T]{parse, format, new([]T)}
	fs.Var(&g, name, usage)
	return g.values
}

func GenericsVar[T any](fs *flag.FlagSet, p *[]T, name, usage string, parse func(string) (T, error), format func(T) string) {
	g := generics[T]{parse, format, p}
	fs.Var(&g, name, usage)
}
