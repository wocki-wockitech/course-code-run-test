package main

import "strings"

// Concat joins all parts into a single string.
// Uses strings.Builder with pre-calculated capacity for minimal allocations.
func Concat(parts []string) string {
	total := 0
	for _, p := range parts {
		total += len(p)
	}
	var b strings.Builder
	b.Grow(total)
	for _, p := range parts {
		b.WriteString(p)
	}
	return b.String()
}
