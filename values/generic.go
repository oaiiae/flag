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

func Generic[T any](parse func(string) (T, error), format func(T) string) flag.Value {
	return &generic[T]{parse, format, new(T), false}
}

func GenericVar[T any](p *T, parse func(string) (T, error), format func(T) string) flag.Value {
	return &generic[T]{parse, format, p, true}
}

type generics[T any] struct {
	split  func(string) []string
	parse  func(string) (T, error)
	format func(T) string
	values *[]T
}

func (v *generics[T]) Set(s string) error {
	for _, s := range v.split(s) {
		val, err := v.parse(s)
		if err != nil {
			return err
		}
		*v.values = append(*v.values, val)
	}
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

func Generics[T any](parse func(string) (T, error), format func(T) string, split func(string) []string) flag.Value {
	return &generics[T]{split, parse, format, new([]T)}
}

func GenericsVar[T any](p *[]T, parse func(string) (T, error), format func(T) string, split func(string) []string) flag.Value {
	return &generics[T]{split, parse, format, p}
}
