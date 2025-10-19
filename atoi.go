package clmconv

import (
	"errors"
	"math"
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

	// Length check for potential overflow
	// 14 characters is the practical limit for int64
	if len(s) > 14 {
		return 0, errors.New("integer overflow")
	}

	r := -c.offset
	for i, char := range s {
		var charValue int
		if 'A' <= char && char <= 'Z' {
			charValue = int(char) - runeOffsetUpper
		} else if 'a' <= char && char <= 'z' {
			charValue = int(char) - runeOffsetLower
		} else {
			return 0, errors.New("must not contain non-alphabetic characters")
		}

		power := pow26(len(s) - i - 1)

		// Check for multiplication overflow
		if charValue > 0 && power > math.MaxInt/charValue {
			return 0, errors.New("integer overflow")
		}

		product := charValue * power

		// Check for addition overflow
		if r > math.MaxInt-product {
			return 0, errors.New("integer overflow")
		}

		r += product
	}

	return r, nil
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
