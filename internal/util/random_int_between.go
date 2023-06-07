package util

import "math/rand"

func RandomIntBetween(min, max int) int {
	if min > max {
		min, max = max, min
	}

	return rand.Intn(max-min+1) + min
}
