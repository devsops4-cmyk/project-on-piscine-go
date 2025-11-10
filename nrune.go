package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	if n <= 0 || n > len(runes) {
		return 0
	}
	return runes[n-1] // subtract 1 to convert 1-based to 0-based index
}
