package main

import (
	"fmt"
	"os"
	"strconv"
)

func printLastBytes(filename string, count int) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if count > len(data) {
		count = len(data)
	}

	fmt.Print(string(data[len(data)-count:]))
	return nil
}

func main() {
	if len(os.Args) < 3 || os.Args[1] != "-c" {
		fmt.Println("Usage: go run . -c <count> <file1> [file2 ...]")
		os.Exit(1)
	}

	// Parse byte count
	count, err := strconv.Atoi(os.Args[2])
	if err != nil || count < 0 {
		fmt.Println("ERROR: invalid byte count")
		os.Exit(1)
	}

	files := os.Args[3:]
	exitCode := 0

	for i, f := range files {
		// Print separator if multiple files
		if len(files) > 1 {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", f)
		}

		err := printLastBytes(f, count)
		if err != nil {
			fmt.Println(err)
			exitCode = 1
		}
	}

	os.Exit(exitCode)
}
