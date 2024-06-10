package utils

// Set creates a slice of unique elements from an input slice. Order is not guaranteed.
//
// # Example
//
//	Set([]string{"one", "two", "two"}) // []string{"one", "two"}
func Set[T comparable](a []T) []T {
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
