package series

func All(n int, s string) []string {
	var result []string

	for i := 0; i < len(s)-n+1; i++ {
		result = append(result, s[i:i+n])
	}
	return result
}

func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

func First(n int, s string) (string, bool) {
	if len(s) < n {
		return "", false
	}
	return s[0:n], true
}
