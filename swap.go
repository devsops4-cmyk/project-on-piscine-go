package piscine

func Swap(a *int, b *int) {
	temp := *a // store the value of a
	*a = *b    // set a to the value of b
	*b = temp  // set b to the original value of a
}
