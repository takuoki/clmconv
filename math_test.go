package clmconv_test

import (
	"testing"

	"github.com/takuoki/clmconv"
)

// Usually we only use small numbers.
var testInts = []int{0, 1, 2, 3, 4, 5}

func BenchmarkPow26(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range testInts {
			clmconv.Pow26(v)
		}
	}
}

func BenchmarkPow26Inner(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range testInts {
			pow26(v)
		}
	}
}

func BenchmarkPow26NormalInner(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range testInts {
			pow26Normal(v)
		}
	}
}

func BenchmarkPow26WithMemoInner(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range testInts {
			pow26WithMemo(v)
		}
	}
}

var pow26tab = [...]int{1, 26, 676, 17576, 456976, 11881376}

func pow26(n int) int {
	if 0 <= n && n <= 5 {
		return pow26tab[n]
	}
	return pow26(n-1) * 26
}

func pow26Normal(n int) int {
	if n == 0 {
		return 1
	}
	return pow26Normal(n-1) * 26
}

var m = map[int]int{}

func pow26WithMemo(n int) int {
	if n == 0 {
		return 1
	}
	if v, ok := m[n]; ok {
		return v
	}
	r := pow26WithMemo(n-1) * 26
	m[n] = r
	return r
}
