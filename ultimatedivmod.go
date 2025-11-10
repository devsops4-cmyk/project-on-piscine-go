package piscine

func UltimateDivMod(a *int, b *int) {
	div := *a / *b // divide the values a and b point to
	mod := *a % *b // find the remainder
	*a = div       // store the division result in a
	*b = mod       // store the remainder in b
}
