package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	// Save original values to temporary variables
	tmpA := ***a
	tmpB := *b
	tmpC := *******c
	tmpD := ****d

	// Perform the swaps
	*******c = tmpA // a -> c
	****d = tmpC    // c -> d
	*b = tmpD       // d -> b
	***a = tmpB     // b -> a
}
