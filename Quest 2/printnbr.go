package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		printNegative(n)
		return
	}

	printPositive(n)
}

func printPositive(n int) {
	if n >= 10 {
		printPositive(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}

func printNegative(n int) {
	if n <= -10 {
		printNegative(n / 10)
	}
	z01.PrintRune(rune(-(n % 10) + '0'))
}
