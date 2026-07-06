package piscine

func ConcatParams(args []string) string {
	r := ""
	for i, c := range args {
		r = r + c
		if i != len(args)-1 {
			r += "\n"
		}
	}
	return r
}
