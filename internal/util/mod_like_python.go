package util

func ModLikePython(n, m int) int {
	if m < 0 {
		m = -m
	}
	if n < 0 {
		return ((n % m) + m) % m
	}
	return n % m
}
