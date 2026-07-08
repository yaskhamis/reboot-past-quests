package piscine

func StrRev(s string) string {
	arr := []rune(s)
	char := ""
	for i := len(arr) - 1; i >= 0; i-- {
		char += string(arr[i])
	}
	return char
}
