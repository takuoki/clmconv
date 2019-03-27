package clmconv_test

import (
	"testing"

	"github.com/takuoki/clmconv"
)

func TestToNum(t *testing.T) {
	cases := []struct {
		str      string
		expected int
		err      string
	}{
		// success
		{str: "", expected: 0},
		{str: "A", expected: 1},
		{str: "B", expected: 2},
		{str: "Z", expected: 26},
		{str: "AA", expected: 27},
		{str: "ZZ", expected: 702},
		{str: "ABC", expected: 731},
		{str: "ABCDE", expected: 494265},
		{str: "a", expected: 1},
		{str: "AbCdE", expected: 494265},

		// error
		{str: "1", err: "must not contain non-alphabetic characters"},
		{str: "AB!", err: "must not contain non-alphabetic characters"},
	}

	for _, c := range cases {
		r, err := clmconv.ToNum(c.str)
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
