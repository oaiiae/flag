package values

import "strings"

func Unsplit(s string) []string { return []string{s} }

func Split(sep string) func(s string) []string {
	return func(s string) []string { return strings.Split(s, sep) }
}
