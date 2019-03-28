package clmconv_test

import (
	"math"
	"testing"

	"github.com/takuoki/clmconv"
)

func TestAtoi(t *testing.T) {
	cases := []struct {
		str      string
		expected int
		err      string
	}{
		// success
		{str: "A", expected: 0},
		{str: "B", expected: 1},
		{str: "Z", expected: 25},
		{str: "AA", expected: 26},
		{str: "ZZ", expected: 701},
		{str: "ABC", expected: 730},
		{str: "ABCDE", expected: 494264},
		{str: "FXSHRXX", expected: math.MaxInt32},
		{str: "CRPXNLSKVLJFHG", expected: math.MaxInt64 - 1},
		{str: "CRPXNLSKVLJFHH", expected: math.MaxInt64},
		{str: "a", expected: 0},
		{str: "AbCdE", expected: 494264},

		// error
		{str: "", err: "argument is empty string"},
		{str: "1", err: "must not contain non-alphabetic characters"},
		{str: "AB!", err: "must not contain non-alphabetic characters"},
	}

	for _, c := range cases {
		r, err := clmconv.Atoi(c.str)
		if c.err == "" {
			if err != nil {
				t.Errorf("error must not occur: %s", err.Error())
			} else if r != c.expected {
				t.Errorf("value doesn't match: %d (expected %d)", r, c.expected)
			}
		} else {
			if err == nil {
				t.Errorf("error must occur: expected %s", c.err)
			} else if err.Error() != c.err {
				t.Errorf("error doesn't match: %s (expected %s)", err.Error(), c.err)
			}
		}
	}
}

func TestMustAtoi(t *testing.T) {
	// success
	func() {
		defer func() {
			if e := recover(); e != nil {
				t.Errorf("panic must not occur: %v", e)
			}
		}()
		clmconv.MustAtoi("ABC")
	}()
	// error
	func() {
		defer func() {
			if e := recover(); e == nil {
				t.Errorf("panic must occur")
			}
		}()
		clmconv.MustAtoi("AB!")
	}()
}
