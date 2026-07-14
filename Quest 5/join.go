package piscine

func Join(strs []string, sep string) string {
	v := strs
	f := ""
	for i := 0; i < len(v); i++ {
		if i < len(v)-1 {
			f += v[i] + sep
		} else {
			f += v[i]
		}
	}
	return f
}
