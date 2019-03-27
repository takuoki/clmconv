package clmconv_test

import (
	"math"
	"testing"

	"github.com/takuoki/clmconv"
)

func TestItoa(t *testing.T) {
	cases := []struct {
		i        int
		expected string
	}{
		{i: -1, expected: ""},
		{i: 0, expected: "A"},
		{i: 1, expected: "B"},
		{i: 25, expected: "Z"},
		{i: 26, expected: "AA"},
		{i: 701, expected: "ZZ"},
		{i: 730, expected: "ABC"},
		{i: 494264, expected: "ABCDE"},
		{i: math.MaxInt64, expected: "CRPXNLSKVLJFHH"},
	}

	for _, c := range cases {
		r := clmconv.Itoa(c.i)
		if r != c.expected {
			t.Errorf("value doesn't match: %s (expected %s)", r, c.expected)
		}
	}
}
