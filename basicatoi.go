package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, c := range s {
		// convert rune '0'..'9' to integer 0..9
		digit := int(c - '0')
		result = result*10 + digit
	}
	return result
}
