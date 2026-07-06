package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for a := 99; a >= 1; a-- {
		for b := a - 1; b >= 0; b-- {
			printTwoDigits(a)
			printRune(' ')
			printTwoDigits(b)

			if !(a == 1 && b == 0) {
				printRune(',')
				printRune(' ')
			}
		}
	}
}

func printTwoDigits(n int) {
	printRune(rune(n/10 + '0'))
	printRune(rune(n%10 + '0'))
}

func printRune(r rune) {
	z01.PrintRune(r)
}
