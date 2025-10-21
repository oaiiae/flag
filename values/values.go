// Package values provides implementations of [flag.Value] and primitives to register them.
//
// Aside of [RegistererFunc], there is 30 functions declaring various [flag.Value].
// Their names are matched by this regular expression:
//
//	(Generic|Basic|Stringer|Time|Duration)(List|Slice)?(Var)?
//
// If neither 'List' nor 'Slice' are present, then the value is parsed and
// stored to a variable. Multiple sets will overwrite the value.
//
// If 'List' is present, the value is then a slice and can be set multiple times.
//
// If 'Slice' is present, the value is also a slice, but all its values are set
// at once every time the flag is invoked. The [flag.Value] will split the input
// string and parse the substrings.
//
// If 'Var' is present, the function accepts another pointer parameter which
// will be used to store the parsed values.
//
// On top of 'Generic', several [flag.Value] variants are implemented to simplify
// common use-cases:
//
//   - 'Basic' for Go's basic types (int, uint64, float64, string, ...)
//   - 'Stringer' for object imlementing [fmt.Stringer] ([*url.URL], [netip.Addr], ...)
//   - 'Time' takes a layout for use in [time.Time.Format] and [time.Parse]
//   - 'Duration' for [time.Duration] values
//
// The values shall then be registered using [flag.FlagSet.Var].
package values

import (
	_ "flag"      // for documentation links
	_ "fmt"       // for documentation links
	_ "net/netip" // for documentation links
	_ "net/url"   // for documentation links
	_ "time"      // for documentation links
)
