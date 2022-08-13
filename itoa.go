package clmconv

// Itoa converts number to alphabet.
// If argument is negative, returns empty string.
func Itoa(i int) string {
	return defaultConverter.Itoa(i)
}

// Itoa converts number to alphabet with converter.
// If argument is negative, returns empty string.
func (c *Converter) Itoa(i int) string {
	i = i + c.offset - 1
	if i < 0 {
		return ""
	}
	var r []rune
	for char := 0; ; char++ {
		mod := i%pow26(char+1) + 1
		r = append([]rune{rune(mod/pow26(char) + c.runeOffset)}, r...)
		i -= mod
		if i <= 0 {
			break
		}
	}
	return string(r)
}
