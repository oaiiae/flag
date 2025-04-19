package values

import (
	"flag"
	"time"
)

func parseTime(layout string) func(s string) (time.Time, error) {
	return func(s string) (time.Time, error) { return time.Parse(layout, s) }
}

func formatTime(layout string) func(t time.Time) string {
	return func(t time.Time) string { return t.Format(layout) }
}

func Time(layout string) flag.Value {
	return Generic(parseTime(layout), formatTime(layout))
}

func TimeVar(p *time.Time, layout string) flag.Value {
	return GenericVar(p, parseTime(layout), formatTime(layout))
}

func Times(layout string, split func(string) []string) flag.Value {
	return Generics(parseTime(layout), formatTime(layout), split)
}

func TimesVar(p *[]time.Time, layout string, split func(string) []string) flag.Value {
	return GenericsVar(p, parseTime(layout), formatTime(layout), split)
}
