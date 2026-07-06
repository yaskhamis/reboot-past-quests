package piscine

func JumpOver(str string) string {
	result := ""

	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}

	result += "\n"

	return result
}
