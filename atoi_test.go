package clmconv_test

import (
	"math"
	"testing"

	"github.com/takuoki/clmconv"
)

func TestAtoi(t *testing.T) {
	cases := []struct {
		a        string
		expected int
		err      string
	}{
		// success
		{a: "A", expected: 0},
		{a: "B", expected: 1},
		{a: "Z", expected: 25},
		{a: "AA", expected: 26},
		{a: "ZZ", expected: 701},
		{a: "ABC", expected: 730},
		{a: "ABCDE", expected: 494264},
		{a: "FXSHRXX", expected: math.MaxInt32},
		{a: "CRPXNLSKVLJFHG", expected: math.MaxInt64 - 1},
		{a: "CRPXNLSKVLJFHH", expected: math.MaxInt64},
		{a: "a", expected: 0},
		{a: "AbCdE", expected: 494264},

		// error
		{a: "", err: "argument is empty string"},
		{a: "1", err: "must not contain non-alphabetic characters"},
		{a: "AB!", err: "must not contain non-alphabetic characters"},
	}

	for _, c := range cases {
		r, err := clmconv.Atoi(c.a)
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
