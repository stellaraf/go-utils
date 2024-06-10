package utils

// ChunkSlice creates a slice of slices of max size.
func ChunkSlice[T any](items []T, size int) [][]T {
	chunks := make([][]T, 0, (len(items)/size)+1)
	for size > len(items) {
		items, chunks = items[size:], append(chunks, items[0:size:size])
	}
	return chunks
}
