package main

import (
	"os"

	"github.com/01-edu/z01"
)

func vowel(r rune) bool {
	if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' {
		return true
	}
	return false
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		z01.PrintRune('\n')
		return
	}

	var v []rune

	for _, word := range args {
		for _, ch := range word {
			if vowel(ch) {
				v = append(v, ch)
			}
		}
	}

	j := len(v) - 1

	for i, word := range args {
		for _, ch := range word {
			if vowel(ch) {
				z01.PrintRune(v[j])
				j--
			} else {
				z01.PrintRune(ch)
			}
		}
		if i != len(args)-1 {
			z01.PrintRune(' ')
		}
	}

	z01.PrintRune('\n')
}
