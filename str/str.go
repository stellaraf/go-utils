package str

// Truncate truncates an input string to a specific length with a trailing '...' when necessary.
func Truncate(str string, length int) string {
	if length <= 0 {
		return ""
	}
	if length > 3 {
		length -= 3
	}
	count := 0
	truncated := ""
	for _, char := range str {
		truncated += string(char)
		count++
		if count >= length {
			break
		}
	}
	if len(truncated) == length {
		truncated += "..."
	}
	return truncated
}

// RemoveEmpty removes empty strings from a string slice.
func RemoveEmpty(strs []string) []string {
	out := make([]string, 0, len(strs))
	for _, s := range strs {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
