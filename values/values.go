// Package values provides a flexible way for defining [flag.Value].
//
// Aside of [RegistererFunc], there is 30 functions declaring various [flag.Value].
// Their names are matched by this regular expression:
//
//	(Generic|Basic|Stringer|Time|Duration)(List|Slice)?(Var)?
//
// If neither 'List' nor 'Slice' are present, then the value is parsed and
// stored to a variable. Multiple sets will effectively overwrite the value.
//
// If 'List' is present, the value may be set multiple times. Each successfully
// parsed string is actually appended to a slice.
//
// If 'Slice' is present, the flag may be invoked with a string formatted as
// multiple values joined with a separator. The [flag.Value] will split the
// input string before parsing the substrings into a slice as well. Calling the
// flag multiple times overwrites the full slice and does NOT append new values.
//
// If 'Var' is present, the function accepts another pointer parameter which
// will be used to store the parsed values.
//
// On top of 'Generic', which is at the root of the whole package, there is
// several variants to simplify standard usages, while keeping the flexibility
// brought by the generic implementation:
//
//   - 'Basic' for Go's basic types (int, uint64, float64, string, ...)
//   - 'Stringer' [fmt.Stringer] implementors ([*url.URL], [netip.Addr], ...)
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
