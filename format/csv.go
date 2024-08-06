package format

import (
	"fmt"
	"strings"
)

// CommaSeparatedValues formats a slice of values as comma-separated values.
func CommaSeparatedValues[T comparable](in []T) string {
	strs := make([]string, len(in))
	for i := 0; i < len(in); i++ {
		strs[i] = fmt.Sprint(in[i])
	}
	return strings.Join(strs, ",")
}
