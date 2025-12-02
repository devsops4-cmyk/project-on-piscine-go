package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	const EvenMsg = "I have an even number of arguments"
	const OddMsg = "I have an odd number of arguments"

	argCount := len(os.Args[1:])

	if isEven(argCount) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
