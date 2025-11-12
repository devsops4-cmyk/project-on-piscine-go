package piscine

func Split(s, sep string) []string {
	if sep == "" {
		return []string{s} // edge case: empty separator returns the whole string
	}

	// First, count how many separators are in s
	count := 0
	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			count++
			i += len(sep) - 1 // skip past the separator
		}
	}

	// Preallocate slice
	result := make([]string, 0, count+1)

	start := 0
	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			result = append(result, s[start:i])
			start = i + len(sep)
			i = start - 1 // move i to end of separator
		}
	}
	// append the last substring after the final separator
	result = append(result, s[start:])

	return result
}
