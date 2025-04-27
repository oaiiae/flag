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

// Time declares a [flag.Value] for a single [time.Time] value, using layout for parsing and formatting.
// The actual value type is [time.Time].
func Time(layout string) flag.Value {
	return Generic(parseTime(layout), formatTime(layout))
}

// TimeVar is like [Time] but stores the value in p.
func TimeVar(p *time.Time, layout string) flag.Value {
	return GenericVar(p, parseTime(layout), formatTime(layout))
}

// TimeList declares a list-style [flag.Value] for multiple [time.Time] values, using layout for parsing and formatting.
// The actual value type is [][time.Time].
func TimeList(layout string) flag.Value {
	return GenericList(parseTime(layout), formatTime(layout))
}

// TimeListVar is like [TimeList] but stores the values in p.
func TimeListVar(p *[]time.Time, layout string) flag.Value {
	return GenericListVar(p, parseTime(layout), formatTime(layout))
}

// TimeSlice declares a slice-style [flag.Value] for multiple [time.Time] values, using layout for parsing and formatting.
// The input strings are split around sep before parsing.
// The actual value type is [][time.Time].
func TimeSlice(sep string, layout string) flag.Value {
	return GenericSlice(sep, parseTime(layout), formatTime(layout))
}

// TimeSliceVar is like [TimeSlice] but stores the values in p.
func TimeSliceVar(p *[]time.Time, sep string, layout string) flag.Value {
	return GenericSliceVar(p, sep, parseTime(layout), formatTime(layout))
}

func formatDuration(d time.Duration) string { return d.String() }

// Duration declares a [flag.Value] for a single [time.Duration] value.
// The actual value type is [time.Duration].
func Duration() flag.Value {
	return Generic(time.ParseDuration, formatDuration)
}

// DurationVar is like [Duration] but stores the value in p.
func DurationVar(p *time.Duration) flag.Value {
	return GenericVar(p, time.ParseDuration, formatDuration)
}

// DurationList declares a list-style [flag.Value] for multiple [time.Duration] values.
// The actual value type is [][time.Duration].
func DurationList() flag.Value {
	return GenericList(time.ParseDuration, formatDuration)
}

// DurationListVar is like [DurationList] but stores the values in p.
func DurationListVar(p *[]time.Duration) flag.Value {
	return GenericListVar(p, time.ParseDuration, formatDuration)
}

// DurationSlice declares a slice-style [flag.Value] for multiple [time.Duration] values.
// The input strings are split around sep before parsing.
// The actual value type is [][time.Duration].
func DurationSlice(sep string) flag.Value {
	return GenericSlice(sep, time.ParseDuration, formatDuration)
}

// DurationSliceVar is like [DurationSlice] but stores the values in p.
func DurationSliceVar(p *[]time.Duration, sep string) flag.Value {
	return GenericSliceVar(p, sep, time.ParseDuration, formatDuration)
}
