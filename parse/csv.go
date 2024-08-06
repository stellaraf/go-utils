package parse

import (
	"regexp"
	"strconv"
	"strings"
)

// CommaSeparatedNumbers parses a comma-separated string of integers to a slice of integers.
func CommaSeparatedNumbers(raw string) []int {
	p := regexp.MustCompile(`[^0-9\,]+`)
	filtered := p.ReplaceAllString(raw, "")
	parts := strings.Split(filtered, ",")
	numbers := make([]int, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		n, err := strconv.Atoi(p)
		if err == nil && n != 0 {
			numbers = append(numbers, n)
		}
	}
	return numbers
}

// CommaSeparatedValues parses a comma-separated string of values into a slice of trimmed strings.
func CommaSeparatedValues(value string) []string {
	split := strings.Split(value, ",")
	values := make([]string, len(split))
	for i := 0; i < len(values); i++ {
		values[i] = strings.TrimSpace(split[i])
	}
	return values
}
