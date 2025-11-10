package piscine

func LastRune(s string) rune {
	runes := []rune(s) // convert string to slice of runes
	if len(runes) == 0 {
		return 0 // handle empty string
	}
	return runes[len(runes)-1] // return the last rune
}
