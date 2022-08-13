package clmconv_test

import (
	"math"
	"testing"

	"github.com/takuoki/clmconv"
)

func TestConverter_Itoa(t *testing.T) {
	dc := clmconv.New()
	cc := clmconv.New(clmconv.WithStartFromOne(), clmconv.WithLowercase())

	cases := []struct {
		conv     *clmconv.Converter
		i        int
		expected string
	}{
		// success (default converter)
		{conv: dc, i: -1, expected: ""},
		{conv: dc, i: 0, expected: "A"},
		{conv: dc, i: 1, expected: "B"},
		{conv: dc, i: 25, expected: "Z"},
		{conv: dc, i: 26, expected: "AA"},
		{conv: dc, i: 701, expected: "ZZ"},
		{conv: dc, i: 730, expected: "ABC"},
		{conv: dc, i: 494264, expected: "ABCDE"},
		{conv: dc, i: math.MaxInt32, expected: "FXSHRXX"},
		{conv: dc, i: math.MaxInt64 - 1, expected: "CRPXNLSKVLJFHG"},
		{conv: dc, i: math.MaxInt64, expected: "CRPXNLSKVLJFHH"},

		// success (custom converter)
		{conv: cc, i: 0, expected: ""},
		{conv: cc, i: 1, expected: "a"},
		{conv: cc, i: 2, expected: "b"},
		{conv: cc, i: 26, expected: "z"},
		{conv: cc, i: 27, expected: "aa"},
		{conv: cc, i: 702, expected: "zz"},
		{conv: cc, i: 731, expected: "abc"},
		{conv: cc, i: 494265, expected: "abcde"},
		{conv: cc, i: math.MaxInt32, expected: "fxshrxw"},
		{conv: cc, i: math.MaxInt64 - 1, expected: "crpxnlskvljfhf"},
		{conv: cc, i: math.MaxInt64, expected: "crpxnlskvljfhg"},
	}

	for _, c := range cases {
		r := c.conv.Itoa(c.i)
		if r != c.expected {
			t.Errorf("value doesn't match: %s (expected %s)", r, c.expected)
		}
	}
}
