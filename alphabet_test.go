package clmconv_test

import (
	"testing"

	"github.com/takuoki/clmconv"
)

func TestToAlphabet(t *testing.T) {
	cases := []struct {
		i        int
		expected string
	}{
		{i: 0, expected: ""},
		{i: -1, expected: ""},
		{i: 1, expected: "A"},
		{i: 2, expected: "B"},
		{i: 26, expected: "Z"},
		{i: 27, expected: "AA"},
		{i: 702, expected: "ZZ"},
		{i: 731, expected: "ABC"},
		{i: 494265, expected: "ABCDE"},
	}

	for _, c := range cases {
		r := clmconv.ToAlphabet(c.i)
		if r != c.expected {
			t.Errorf("value doesn't match: %s (expected %s)", r, c.expected)
		}
	}
}
