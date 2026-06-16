package main

// Concat joins all parts into a single string.
// TODO: optimize to minimize allocations.
func Concat(parts []string) string {
	result := ""
	for _, p := range parts {
		result += p // naive: allocates on every iteration
	}
	return result
}
