package piscine

func ToLower(s string) string {
	result := []rune(s) // convert string to rune slice
	for i, r := range result {
		if r >= 'A' && r <= 'Z' { // if uppercase
			result[i] = r + ('a' - 'A') // convert to lowercase
		}
	}
	return string(result) // convert back to string
}
