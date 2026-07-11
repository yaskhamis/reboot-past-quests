package piscine

// FindNextPrime returns the first prime number
// that is greater than or equal to nb.
func FindNextPrime(nb int) int {
	// If the number is 2 or smaller,
	// the first prime is always 2.
	if nb <= 2 {
		return 2
	}
	// If the number is even,
	// move to the next odd number.
	if nb%2 == 0 {
		nb++
	}
	// Keep checking until we find a prime number.
	for !isPrime(nb) {
		// Skip to the next odd number.
		// We add 2 because even numbers
		// (except 2) are never prime.
		nb += 2
	}
	// Return the prime number we found.
	return nb
}

// isPrime checks if a number is prime.
func isPrime(n int) bool {
	// Numbers smaller than 2 are not prime.
	if n < 2 {
		return false
	}
	// 2 is the only even prime number.
	if n == 2 {
		return true
	}
	// If the number is even,
	// it cannot be prime.
	if n%2 == 0 {
		return false
	}
	// Check only odd numbers starting from 3.
	// Stop when i*i is greater than n.
	for i := 3; i*i <= n; i += 2 {
		// If n is divisible by i,
		// it is not prime.
		if n%i == 0 {
			return false
		}
	}
	// No divisors were found,
	// so the number is prime.
	return true
}
