package main

import "github.com/01-edu/z01"

type point struct {
	x int

	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	for _, char := range []rune{'x', ' ', '=', ' ', 4 + '0', 2 + '0', ',', ' ', 'y', ' ', '=', ' ', 2 + '0', 1 + '0', '\n'} {
		z01.PrintRune(char)
	}
	setPoint(points)
}
