package clmconv

// ToAlphabet converts number to alphabet.
// If argument is zero or negative, returns empty string.
func ToAlphabet(i int) string {
	if i <= 0 {
		return ""
	}

	c := 0
	r := make([]rune, 0, 8)
	for i > 0 {
		mod := (i-1)%pow(26, c+1) + 1
		r = append([]rune{rune(mod/pow(26, c) + 64)}, r...)
		i = i - mod
		c++
	}
	return string(r)
}
