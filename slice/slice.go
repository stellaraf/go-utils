package slice

import "strings"

// Contains determines if a given slice contains a given item.
func Contains[T comparable](arr []T, item T) bool {
	for _, element := range arr {
		if element == item {
			return true
		}
	}
	return false
}

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

// Dedup creates a slice of unique elements from an input slice. Order is not guaranteed.
//
// # Example
//
//	Dedup([]string{"one", "two", "two"}) // []string{"one", "two"}
func Dedup[T comparable](a []T) []T {
	m := make(map[T]int, len(a))
	for i := 0; i < len(a); i++ {
		m[a[i]] = 0
	}
	s := make([]T, 0, len(a))
	for k := range m {
		s = append(s, k)
	}
	return s
}

// Chunk creates a slice of slices of a max chunk size.
// Given a size argument of 4:
//
//	[]int{1,2,3,4,5,6,7,9}
//	// Becomes:
//	[][]int{ {1,2,3,4}, {5,6,7,8}, {9} }
func Chunk[T any](items []T, size int) [][]T {
	batches := make([][]T, 0, (len(items)+size-1)/size)
	for size < len(items) {
		items, batches = items[size:], append(batches, items[0:size:size])
	}
	batches = append(batches, items)
	return batches
}

// HasAll determines if all elements of two slices are contained within both slices.
func HasAll[T comparable](s1 []T, s2 []T) bool {
	for _, el := range s1 {
		if !Contains(s2, el) {
			return false
		}
	}
	for _, el := range s2 {
		if !Contains(s1, el) {
			return false
		}
	}
	return true
}
