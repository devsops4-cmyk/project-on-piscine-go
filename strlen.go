package piscine

func StrLen(s string) int {
	count := 0
	for range s { // iterates over runes
		count++
	}
	return count
}
