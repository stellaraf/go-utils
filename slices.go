package utils

import "strings"

// FilterPartialDuplicates takes a slice of strings and filters out items that are a partial match
// for other items in the slice. For example:
//
//	[]string{"thing1", "thing that has thing1", "thing2", "thing that has thing2", "thing3"}
//	// Becomes:
//	[]string{"thing1", "thing2", "thing3"}
func FilterPartialDuplicates(in []string) []string {
	out := make([]string, 0, len(in))
	// Iterate over each item in the input slice.
	for _, i := range in {
		// Check if the current item contains any other items in the slice.
		contains := false
		for _, otherItem := range in {
			if i != otherItem && strings.Contains(i, otherItem) {
				contains = true
				break
			}
		}
		// If the item doesn't contain any other items, add it to the result slice.
		if !contains {
			out = append(out, i)
		}
	}
	return out
}
