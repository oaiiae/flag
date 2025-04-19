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

func Time(fs *flag.FlagSet, name, usage string, layout string) *time.Time {
	return Generic(fs, name, usage, parseTime(layout), formatTime(layout))
}

func TimeVar(fs *flag.FlagSet, p *time.Time, name, usage string, layout string) {
	GenericVar(fs, p, name, usage, parseTime(layout), formatTime(layout))
}

func Times(fs *flag.FlagSet, name, usage string, layout string, split func(string) []string) *[]time.Time {
	return Generics(fs, name, usage, parseTime(layout), formatTime(layout), split)
}

func TimesVar(fs *flag.FlagSet, p *[]time.Time, name, usage string, layout string, split func(string) []string) {
	GenericsVar(fs, p, name, usage, parseTime(layout), formatTime(layout), split)
}
