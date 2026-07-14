package main

import (
	"fmt"
	"os"
)

func help() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		help()
		return
	}

	s := ""
	insert := ""
	order := false

	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			help()
			return
		} else if arg == "-o" || arg == "--order" {
			order = true
		} else if len(arg) > 9 && arg[:9] == "--insert=" {
			insert = arg[9:]
		} else if len(arg) > 3 && arg[:3] == "-i=" {
			insert = arg[3:]
		} else {
			s = arg
		}
	}

	s += insert

	if order {
		r := []rune(s)
		for i := 0; i < len(r); i++ {
			for j := i + 1; j < len(r); j++ {
				if r[i] > r[j] {
					r[i], r[j] = r[j], r[i]
				}
			}
		}
		s = string(r)
	}

	fmt.Println(s)
}
