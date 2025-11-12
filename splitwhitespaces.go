package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	wordStart := -1

	for i, char := range s {
		if char != ' ' && char != '\t' && char != '\n' {
			if wordStart == -1 {
				wordStart = i // start of a word
			}
		} else {
			if wordStart != -1 {
				words = append(words, s[wordStart:i]) // end of a word
				wordStart = -1
			}
		}
	}

	// add the last word if the string didn't end with a separator
	if wordStart != -1 {
		words = append(words, s[wordStart:])
	}

	return words
}
