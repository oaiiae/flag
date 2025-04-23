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

func formatDuration(d time.Duration) string { return d.String() }

func Duration() flag.Value {
	return Generic(time.ParseDuration, formatDuration)
}

func DurationVar(p *time.Duration) flag.Value {
	return GenericVar(p, time.ParseDuration, formatDuration)
}

func DurationList() flag.Value {
	return GenericList(time.ParseDuration, formatDuration)
}

func DurationListVar(p *[]time.Duration) flag.Value {
	return GenericListVar(p, time.ParseDuration, formatDuration)
}

func DurationSlice(sep string) flag.Value {
	return GenericSlice(sep, time.ParseDuration, formatDuration)
}

func DurationSliceVar(p *[]time.Duration, sep string) flag.Value {
	return GenericSliceVar(p, sep, time.ParseDuration, formatDuration)
}
