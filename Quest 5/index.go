package piscine

func Index(s string, toFind string) int {
	lenS := len(s)
	lenF := len(toFind)

	if lenF == 0 {
		return 0
	}
	if lenF > lenS {
		return -1
	}

	for i := 0; i <= lenS-lenF; i++ {
		match := true
		for j := 0; j < lenF; j++ {
			if s[i+j] != toFind[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}

	return -1
}
