package piscine

import "github.com/01-edu/z01"

// q = queens
// Stores the row position of each queen.
// q[0] = queen position in column 0
var q [8]int

func EightQueens() {
	// Start solving from column 0
	s(0)
}

// s = solve
// c = current column
func s(c int) {
	// If all 8 queens are placed, print the solution
	if c == 8 {
		p()
		return
	}
	// r = row
	// Try every row from 1 to 8
	for r := 1; r <= 8; r++ {
		// Put queen in column c, row r
		q[c] = r
		// ok = okay / safe position
		if ok(c) {
			// Move to the next column
			s(c + 1)
		}
	}
}

// ok = check if the queen placement is safe
// c = current column
func ok(c int) bool {
	// i = previous column
	// Check queens already placed
	for i := 0; i < c; i++ {
		// Same row
		if q[i] == q[c] {
			return false
		}
		// Same diagonal
		if q[i]-i == q[c]-c {
			return false
		}
		// Same diagonal
		if q[i]+i == q[c]+c {
			return false
		}
	}
	return true
}

// p = print solution
func p() {
	// i = column
	for i := 0; i < 8; i++ {
		// Print queen positions
		z01.PrintRune(rune(q[i] + '0'))
	}
	z01.PrintRune('\n')
}
