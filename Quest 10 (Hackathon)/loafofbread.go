package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}

	runes := []rune(str)

	// Count non-space characters
	count := 0
	for _, r := range runes {
		if r != ' ' {
			count++
		}
	}

	// Empty or only spaces
	if count == 0 {
		return "\n"
	}

	// Less than 5 characters
	if count < 5 {
		return "Invalid Output\n"
	}

	result := ""
	group := []rune{}
	index := 0
	first := true

	for index < len(runes) {
		for len(group) < 5 && index < len(runes) {
			if runes[index] != ' ' {
				group = append(group, runes[index])
			}
			index++
		}

		if len(group) == 0 {
			break
		}

		if !first {
			result += " "
		}
		first = false

		result += string(group)
		group = []rune{}

		// Skip one character
		if index < len(runes) {
			index++
		}
	}

	// Add remaining characters
	for index < len(runes) {
		if runes[index] != ' ' {
			result += " " + string(runes[index:])
			break
		}
		index++
	}

	return result + "\n"
}
