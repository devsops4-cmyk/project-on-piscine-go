package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Step 1: count of digits
	counts := [10]int{}

	// Extract digits
	for n > 0 {
		d := n % 10
		counts[d]++
		n /= 10
	}

	// Step 2: print in ascending order
	for digit := 0; digit <= 9; digit++ {
		for i := 0; i < counts[digit]; i++ {
			z01.PrintRune(rune(digit + '0'))
		}
	}
}
