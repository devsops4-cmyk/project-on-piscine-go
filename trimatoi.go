package piscine

func TrimAtoi(s string) int {
	sign := 1
	numStarted := false
	result := 0

	for _, r := range s {
		if r >= '0' && r <= '9' {
			result = result*10 + int(r-'0')
			numStarted = true
		} else if r == '-' && !numStarted {
			sign = -1
		} else if r == '+' && !numStarted {
			continue
		}
		// else: ignore other characters
	}

	return sign * result
}
