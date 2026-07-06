package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 || nb == 3 {
		return true
	}
	if nb%2 == 0 || nb%3 == 0 {
		return false
	}

	// check divisors from 5 to sqrt(nb) skipping even numbers
	for i := 5; i*i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}

	return true
}
