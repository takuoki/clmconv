package clmconv

var pow26tab = [...]int{1, 26, 676, 17576, 456976, 11881376}

func pow26(n int) int {
	if 0 <= n && n <= 5 {
		return pow26tab[n]
	}
	return pow26(n-1) * 26
}
