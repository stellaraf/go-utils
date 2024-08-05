package mmap

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
