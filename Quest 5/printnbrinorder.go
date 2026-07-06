package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Count digits
	digits := [10]int{}
	for n > 0 {
		d := n % 10
		digits[d]++
		n /= 10
	}

	// Print in ascending order
	for i := 0; i <= 9; i++ {
		for digits[i] > 0 {
			z01.PrintRune(rune('0' + i))
			digits[i]--
		}
	}
}
