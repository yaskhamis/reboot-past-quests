package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	str := ""
	for i := 0; i < len(a); i++ {
		str += string(a[i])
		if a[i] != a[len(a)-1] {
			str += "\n"
		}
	}
	for _, k := range str {
		z01.PrintRune(k)
	}
	z01.PrintRune('\n')
}
