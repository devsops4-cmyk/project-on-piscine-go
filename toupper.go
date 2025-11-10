package piscine

func ToUpper(s string) string {
	result := []rune(s) // convert string to rune slice
	for i, r := range result {
		if r >= 'a' && r <= 'z' { // if lowercase
			result[i] = r - ('a' - 'A') // convert to uppercase
		}
	}
	return string(result) // convert back to string
}
