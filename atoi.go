package clmconv

import (
	"errors"
)

// Atoi converts alphabet to number.
func Atoi(s string) (int, error) {
	return defaultConverter.Atoi(s)
}

// Atoi converts alphabet to number with converter.
func (c *Converter) Atoi(s string) (int, error) {
	if s == "" {
		return 0, errors.New("argument is empty string")
	}
	var r int
	for i, char := range s {
		if 'A' <= char && char <= 'Z' {
			r += (int(char) - runeOffsetUpper) * pow26(len(s)-i-1)
		} else if 'a' <= char && char <= 'z' {
			r += (int(char) - runeOffsetLower) * pow26(len(s)-i-1)
		} else {
			return 0, errors.New("must not contain non-alphabetic characters")
		}
	}
	return r - c.offset, nil
}

// MustAtoi converts alphabet to number.
// In case of invalid argument, throws panic.
func MustAtoi(s string) int {
	return defaultConverter.MustAtoi(s)
}

// MustAtoi converts alphabet to number with converter.
// In case of invalid argument, throws panic.
func (c *Converter) MustAtoi(s string) int {
	r, err := c.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return r
}
