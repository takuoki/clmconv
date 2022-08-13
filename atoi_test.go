package clmconv_test

import (
	"math"
	"testing"

	"github.com/takuoki/clmconv"
)

func TestConverter_Atoi(t *testing.T) {
	dc := clmconv.New()
	cc := clmconv.New(clmconv.WithStartFromOne())

	cases := []struct {
		conv     *clmconv.Converter
		a        string
		expected int
		err      string
	}{
		// success (default converter)
		{conv: dc, a: "A", expected: 0},
		{conv: dc, a: "B", expected: 1},
		{conv: dc, a: "Z", expected: 25},
		{conv: dc, a: "AA", expected: 26},
		{conv: dc, a: "ZZ", expected: 701},
		{conv: dc, a: "ABC", expected: 730},
		{conv: dc, a: "ABCDE", expected: 494264},
		{conv: dc, a: "FXSHRXX", expected: math.MaxInt32},
		{conv: dc, a: "CRPXNLSKVLJFHG", expected: math.MaxInt64 - 1},
		{conv: dc, a: "CRPXNLSKVLJFHH", expected: math.MaxInt64},
		{conv: dc, a: "a", expected: 0},
		{conv: dc, a: "AbCdE", expected: 494264},

		// success (custom converter)
		{conv: cc, a: "A", expected: 1},
		{conv: cc, a: "B", expected: 2},
		{conv: cc, a: "Z", expected: 26},
		{conv: cc, a: "AA", expected: 27},
		{conv: cc, a: "ZZ", expected: 702},
		{conv: cc, a: "ABC", expected: 731},
		{conv: cc, a: "ABCDE", expected: 494265},
		{conv: cc, a: "FXSHRXW", expected: math.MaxInt32},
		{conv: cc, a: "CRPXNLSKVLJFHF", expected: math.MaxInt64 - 1},
		{conv: cc, a: "CRPXNLSKVLJFHG", expected: math.MaxInt64},
		{conv: cc, a: "a", expected: 1},
		{conv: cc, a: "AbCdE", expected: 494265},

		// error
		{conv: dc, a: "", err: "argument is empty string"},
		{conv: dc, a: "1", err: "must not contain non-alphabetic characters"},
		{conv: dc, a: "AB!", err: "must not contain non-alphabetic characters"},
	}

	for _, c := range cases {
		r, err := c.conv.Atoi(c.a)
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

func TestConverter_MustAtoi(t *testing.T) {
	dc := clmconv.New()

	t.Run("success", func(t *testing.T) {
		defer func() {
			if e := recover(); e != nil {
				t.Errorf("panic must not occur: %v", e)
			}
		}()
		dc.MustAtoi("ABC")
	})
	t.Run("panic", func(t *testing.T) {
		defer func() {
			if e := recover(); e == nil {
				t.Errorf("panic must occur")
			}
		}()
		dc.MustAtoi("AB!")
	})
}
