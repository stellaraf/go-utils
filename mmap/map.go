package mmap

import "sort"

// HasKey determines if a given map has an entry for a given key.
func HasKey[K comparable, V any](m map[K]V, key K) bool {
	_, hasKey := m[key]
	return hasKey
}

// AssertValue gets an item from a map and asserts its type.
//
// If either the key does not exist in the map or the value is not of the specified type,
// an empty interface and false is returned.
func AssertValue[V any, K comparable](m map[K]any, key K) (V, bool) {
	if val, ok := m[key]; ok {
		if a, ok := val.(V); ok {
			return a, true
		}
	}
	var o V
	return o, false
}

// Sort sorts a map of string keys alphabetically by its keys.
func Sort[T comparable](m map[string]T) (r map[string]T) {
	r = make(map[string]T, len(m))
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		r[k] = m[k]
	}
	return
}

// Merge merges multiple maps into a single map.
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	final := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			final[k] = v
		}
	}
	return final
}
