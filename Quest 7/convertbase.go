package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Step 1: Convert baseFrom to a base-10 integer
	val := 0
	lenFrom := len(baseFrom)
	for i := 0; i < len(nbr); i++ {
		// Find the value of the current character in baseFrom
		index := 0
		for j := 0; j < lenFrom; j++ {
			if baseFrom[j] == nbr[i] {
				index = j
				break
			}
		}
		val = val*lenFrom + index
	}
	// Edge case: if the number is 0
	if val == 0 {
		return string(baseTo[0])
	}
	// Step 2: Convert the base-10 integer to baseTo
	result := ""
	lenTo := len(baseTo)
	for val > 0 {
		remainder := val % lenTo
		result = string(baseTo[remainder]) + result
		val /= lenTo
	}
	return result
}
