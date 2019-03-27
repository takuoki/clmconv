package clmconv

import (
	"errors"
)

// ToNum converts alphabet to number.
// If argument is empty string, returns zero.
func ToNum(s string) (int, error) {
	length := len(s)
	r := 0
	for i, c := range s {
		if 'A' <= c && c <= 'Z' {
			r += (int(c) - 64) * pow(26, length-i-1)
		} else if 'a' <= c && c <= 'z' {
			r += (int(c) - 96) * pow(26, length-i-1)
		} else {
			return 0, errors.New("must not contain non-alphabetic characters")
		}
	}
	return r, nil
}

// MustToNum converts alphabet to number.
// In case of invalid argument, throws panic.
func MustToNum(s string) int {
	r, err := ToNum(s)
	if err != nil {
		panic(err.Error())
	}
	return r
}

func pow(x, y int) int {
	r := 1
	for i := 0; i < y; i++ {
		r = r * x
	}
	return r
}
