package piscine

func Split(s, sep string) []string {
	str := ""
	out := []string{}
	sepLen := len(sep)

	// If separator is empty, just return the string in a slice (or handle as needed)
	if sepLen == 0 {
		return []string{s}
	}

	for i := 0; i < len(s); {
		// Look ahead: does the current position match the separator?
		if i+sepLen <= len(s) && s[i:i+sepLen] == sep {
			if str != "" { // Only append if we have accumulated characters
				out = append(out, str)
				str = ""
			}
			i += sepLen // Skip the entire separator
		} else {
			str += string(s[i]) // Accumulate the character just like your original code
			i++
		}
	}

	// Append the final remaining string if it's not empty
	if str != "" {
		out = append(out, str)
	}

	return out
}
