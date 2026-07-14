package piscine

func AtoiBase(s string, base string) int {
	size := len(base)

	// Check if the base is valid
	if size < 2 {
		return 0
	}

	for i := 0; i < size; i++ {
		if base[i] == '+' || base[i] == '-' {
			return 0
		}
		for j := i + 1; j < size; j++ {
			if base[i] == base[j] {
				return 0
			}
		}
	}

	result := 0

	for i := 0; i < len(s); i++ {
		value := 0

		// Find the value of s[i] in the base
		for j := 0; j < size; j++ {
			if s[i] == base[j] {
				value = j
				break
			}
		}

		result = result*size + value
	}

	return result
}
