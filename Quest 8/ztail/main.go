package main

import (
	"fmt"
	"os"
)

func atoi(s string) int {
	n := 0
	for _, c := range s {
		n = n*10 + int(c-'0')
	}
	return n
}

func main() {
	if len(os.Args) < 4 || os.Args[1] != "-c" {
		return
	}

	n := atoi(os.Args[2])
	hasError := false
	manyFiles := len(os.Args) > 4

	for i := 3; i < len(os.Args); i++ {
		data, err := os.ReadFile(os.Args[i])
		if err != nil {
			fmt.Println(err)
			hasError = true
			continue
		}

		if manyFiles {
			if i != 3 {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", os.Args[i])
		}

		start := 0
		if len(data) > n {
			start = len(data) - n
		}

		fmt.Print(string(data[start:]))
	}

	if hasError {
		os.Exit(1)
	}
}
