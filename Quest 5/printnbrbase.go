package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	size := len(base)

	// Check if the base is valid
	if size < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	for i := 0; i < size; i++ {
		if base[i] == '+' || base[i] == '-' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		for j := i + 1; j < size; j++ {
			if base[i] == base[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}

	var n uint

	if nbr < 0 {
		z01.PrintRune('-')
		n = uint(-(nbr + 1))
		n++
	} else {
		n = uint(nbr)
	}

	printBase(n, base)
}

func printBase(n uint, base string) {
	size := uint(len(base))

	if n >= size {
		printBase(n/size, base)
	}

	z01.PrintRune(rune(base[n%size]))
}
