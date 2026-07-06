package piscine

func Capitalize(s string) string {
	result := ""
	newWord := true

	for _, r := range s {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			if newWord {
				// capitalize first letter if it's a letter
				if r >= 'a' && r <= 'z' {
					r -= 32
				}
				newWord = false
			} else {
				// lowercase letters
				if r >= 'A' && r <= 'Z' {
					r += 32
				}
			}
		} else {
			newWord = true // next alphanumeric character starts a new word
		}
		result += string(r)
	}

	return result
}
