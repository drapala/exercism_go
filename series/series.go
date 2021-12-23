package series

func All(n int, s string) []string {
	result := make([]string, 0)
	// Error handling
	if n > len(s) {
		return result // Empty slice
	}
	start := 0
	end := n
	for end <= len(s) {
		result = append(result, s[start:end])
		start++
		end++
	}
	return result
}

func UnsafeFirst(n int, s string) string {
	return All(n, s)[0]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	} else {
		return All(n, s)[0], true
	}
}