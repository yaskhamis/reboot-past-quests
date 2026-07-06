package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)
	strx := "x = "
	stry := ", y = "
	pp(strx)
	pd(52)
	pd(50)
	pp(stry)
	pd(50)
	pd(49)
	pd(10)
}

func pp(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func pd(s rune) {
	z01.PrintRune(s)
}
