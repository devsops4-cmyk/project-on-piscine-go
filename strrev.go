package piscine

func StrRev(s string) string {
	runes := []rune(s) // convert string to rune slice
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i] // swap
	}

	return string(runes) // convert back to string
}
