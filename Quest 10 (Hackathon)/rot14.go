package piscine

func Rot14(s string) string {
	result := ""
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			c = 'a' + (c-'a'+14)%26
		} else if c >= 'A' && c <= 'Z' {
			c = 'A' + (c-'A'+14)%26
		}
		result += string(c)
	}
	return result
}
