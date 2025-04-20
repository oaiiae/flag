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

func TimeList(layout string) flag.Value {
	return GenericList(parseTime(layout), formatTime(layout))
}

func TimeListVar(p *[]time.Time, layout string) flag.Value {
	return GenericListVar(p, parseTime(layout), formatTime(layout))
}

func TimeSlice(sep string, layout string) flag.Value {
	return GenericSlice(sep, parseTime(layout), formatTime(layout))
}

func TimeSliceVar(p *[]time.Time, sep string, layout string) flag.Value {
	return GenericSliceVar(p, sep, parseTime(layout), formatTime(layout))
}
