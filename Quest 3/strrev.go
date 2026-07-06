package piscine

func StrRev(s string) string {
	runes := []rune(s)         // convert string to slice of runes
	n := len(runes)            // get length
	for i := 0; i < n/2; i++ { // swap from ends toward center
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes) // convert back to string
}
