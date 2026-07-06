package piscine

func SplitWhiteSpaces(s string) []string {
	words := []string{}
	word := ""
	for _, c := range s {
		if c == ' ' || c == '\t' || c == '\n' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(c)
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words
}
