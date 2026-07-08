package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	comb := make([]int, n)
	printComb(0, 0, n, comb)
	z01.PrintRune('\n')
}

func printComb(pos int, start int, n int, comb []int) {
	if pos == n {
		for i := 0; i < n; i++ {
			z01.PrintRune(rune(comb[i] + '0'))
		}

		if comb[0] != 10-n {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}

	for i := start; i <= 9; i++ {
		comb[pos] = i
		printComb(pos+1, i+1, n, comb)
	}
}
