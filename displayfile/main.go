package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		if len(os.Args) < 2 {
			fmt.Println("File name missing")
		} else {
			fmt.Println("Too many arguments")
		}
		return
	}

	content, _ := os.ReadFile(os.Args[1])
	fmt.Print(string(content))
}
