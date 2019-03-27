package clmconv

var m = map[int]int{}

func pow26(i int) int {
	if i == 0 {
		return 1
	}
	if v, ok := m[i]; ok {
		return v
	}
	r := pow26(i-1) * 26
	m[i] = r
	return r
}
