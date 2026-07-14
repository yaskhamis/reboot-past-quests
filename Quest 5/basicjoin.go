package piscine

func BasicJoin(elems []string) string {
	v := elems
	f := ""
	for i := 0; i < len(v); i++ {
		f += v[i]
	}
	return f
}
