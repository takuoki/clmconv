package clmconv

import (
	"errors"
)

// Atoi converts alphabet to number.
func Atoi(s string) (int, error) {
	if s == "" {
		return 0, errors.New("argument is empty string")
	}
	var r int
	for i, c := range s {
		if 'A' <= c && c <= 'Z' {
			r += (int(c) - 64) * pow26(len(s)-i-1)
		} else if 'a' <= c && c <= 'z' {
			r += (int(c) - 96) * pow26(len(s)-i-1)
		} else {
			return 0, errors.New("must not contain non-alphabetic characters")
		}
	}
	return r - 1, nil
}

// MustAtoi converts alphabet to number.
// In case of invalid argument, throws panic.
func MustAtoi(s string) int {
	r, err := Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return r
}
