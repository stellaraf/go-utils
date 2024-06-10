package utils

// ChuckSlice creates a slice of slices of a max chunk size.
// Given a size argument of 4:
//
//	[]int{1,2,3,4,5,6,7,9}
//	// Becomes:
//	[][]int{ {1,2,3,4}, {5,6,7,8}, {9} }
func ChunkSlice[T any](items []T, size int) [][]T {
	batches := make([][]T, 0, (len(items)+size-1)/size)
	for size < len(items) {
		items, batches = items[size:], append(batches, items[0:size:size])
	}
	batches = append(batches, items)
	return batches
}
