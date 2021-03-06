package clmconv

// Itoa converts number to alphabet.
// If argument is negative, returns empty string.
func Itoa(i int) string {
	if i < 0 {
		return ""
	}
	var r []rune
	for c := 0; ; c++ {
		mod := i%pow26(c+1) + 1
		r = append([]rune{rune(mod/pow26(c) + 64)}, r...)
		i -= mod
		if i <= 0 {
			break
		}
	}
	return string(r)
}
