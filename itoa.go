package clmconv

// Itoa converts number to alphabet.
// If argument is negative, returns empty string.
func Itoa(i int) string {
	return defaultConverter.Itoa(i)
}

// Itoa converts number to alphabet with converter.
// If argument is negative, returns empty string.
func (c *Converter) Itoa(i int) string {
	// Adjust input based on converter offset
	adjustedInput := i + c.offset - 1
	if adjustedInput < 0 {
		return ""
	}

	var result []rune
	current := adjustedInput

	// Convert to bijective base-26 (Excel-style column naming)
	for {
		// Calculate position in current digit (0-25)
		remainder := current % 26

		// Convert to rune (A-Z or a-z based on runeOffset)
		char := rune(remainder + c.runeOffset + 1)
		result = append([]rune{char}, result...)

		// Move to next higher digit
		current = current / 26

		// For bijective base-26, we need to adjust when moving to next digit
		if current > 0 {
			current--
		} else {
			break
		}
	}

	return string(result)
}
