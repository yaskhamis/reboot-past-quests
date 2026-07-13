package piscine

func FindNextPrimel(nb int) int {
	if nb < 2 {
		nb = 2
	}
	for i := nb; ; i++ {
		isprime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isprime = false
				break
			}
		}
		if isprime {
			return i
		}
	}
}
