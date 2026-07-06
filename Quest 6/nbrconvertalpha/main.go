package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := false
	start := 0

	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		start = 1
	}

	// Only process if there are arguments
	for i := start; i < len(args); i++ {
		arg := args[i]
		n := 0
		valid := true

		for _, r := range arg {
			if r < '0' || r > '9' {
				valid = false
				break
			}
			n = n*10 + int(r-'0')
		}

		if !valid || n < 1 || n > 26 {
			z01.PrintRune(' ')
			continue
		}

		if upper {
			z01.PrintRune(rune('A' + n - 1))
		} else {
			z01.PrintRune(rune('a' + n - 1))
		}
	}

	// Only print newline if there was at least one argument
	if len(args) > start {
		z01.PrintRune('\n')
	}
}
