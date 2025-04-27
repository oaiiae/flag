package values

import (
	"flag"
	"fmt"
	"strings"
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

// Generic declares a [flag.Value] implemented using the parse & format functions.
// The actual value type is T.
func Generic[T any](parse func(string) (T, error), format func(T) string) flag.Value {
	return &generic[T]{parse, format, new(T), false}
}

// GenericVar is like [Generic] but stores the value in p.
func GenericVar[T any](p *T, parse func(string) (T, error), format func(T) string) flag.Value {
	return &generic[T]{parse, format, p, true}
}

type genericList[T any] struct {
	parse  func(string) (T, error)
	format func(T) string
	values *[]T
}

func (v *genericList[T]) Set(s string) error {
	val, err := v.parse(s)
	if err != nil {
		return err
	}
	*v.values = append(*v.values, val)
	return nil
}

func (v *genericList[T]) String() string {
	if v.values == nil || len(*v.values) == 0 {
		return ""
	}
	a := make([]string, 0, len(*v.values))
	for i := range *v.values {
		a = append(a, v.format((*v.values)[i]))
	}
	return fmt.Sprint(a)
}

func (v *genericList[T]) Get() any {
	return *v.values
}

// GenericList declares a list-style [flag.Value] implemented using the parse & format functions.
// The actual value type is []T.
func GenericList[T any](parse func(string) (T, error), format func(T) string) flag.Value {
	return &genericList[T]{parse, format, new([]T)}
}

// GenericListVar is like [GenericList] but stores the values in p.
func GenericListVar[T any](p *[]T, parse func(string) (T, error), format func(T) string) flag.Value {
	return &genericList[T]{parse, format, p}
}

type genericSlice[T any] struct {
	sep    string
	parse  func(string) (T, error)
	format func(T) string
	values *[]T
}

func (v *genericSlice[T]) Set(s string) error {
	ss := strings.Split(s, v.sep)
	vs := make([]T, len(ss))
	for i, s := range ss {
		parsed, err := v.parse(s)
		if err != nil {
			return err
		}
		vs[i] = parsed
	}
	*v.values = vs
	return nil
}

func (v *genericSlice[T]) String() string {
	if v.values == nil || len(*v.values) == 0 {
		return ""
	}
	b := strings.Builder{}
	b.WriteString(v.format((*v.values)[0]))
	for _, val := range (*v.values)[1:] {
		b.WriteString(v.sep)
		b.WriteString(v.format(val))
	}
	return b.String()
}

func (v *genericSlice[T]) Get() any {
	return *v.values
}

// GenericSlice declares a slice-style [flag.Value] implemented using the parse & format functions.
// The input strings are split around sep before parsing.
// The actual value type is []T.
func GenericSlice[T any](sep string, parse func(string) (T, error), format func(T) string) flag.Value {
	return &genericSlice[T]{sep, parse, format, new([]T)}
}

// GenericSliceVar is like [GenericSlice] but stores the values in p.
func GenericSliceVar[T any](p *[]T, sep string, parse func(string) (T, error), format func(T) string) flag.Value {
	return &genericSlice[T]{sep, parse, format, p}
}
