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
	g := generic[time.Time]{parseTime(layout), formatTime(layout), new(time.Time), false}
	fs.Var(&g, name, usage)
	return g.value
}

func TimeVar(fs *flag.FlagSet, p *time.Time, name, usage string, layout string) {
	g := generic[time.Time]{parseTime(layout), formatTime(layout), p, true}
	fs.Var(&g, name, usage)
}

func Times(fs *flag.FlagSet, name, usage string, layout string) *[]time.Time {
	g := generics[time.Time]{parseTime(layout), formatTime(layout), new([]time.Time)}
	fs.Var(&g, name, usage)
	return g.values
}

func TimesVar(fs *flag.FlagSet, p *[]time.Time, name, usage string, layout string) {
	g := generics[time.Time]{parseTime(layout), formatTime(layout), p}
	fs.Var(&g, name, usage)
}
