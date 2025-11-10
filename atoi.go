package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	sign := 1
	start := 0

	// Check for sign
	if s[0] == '+' {
		start = 1
	} else if s[0] == '-' {
		sign = -1
		start = 1
	}

	result := 0
	for i := start; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' { // Non-digit character
			return 0
		}
		result = result*10 + int(c-'0')
	}

	return result * sign
}
