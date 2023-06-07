package util

import "testing"

func TestRandomIntBetween(t *testing.T) {
	for i := 0; i < 1000; i++ {
		min := -10
		max := 10
		val := RandomIntBetween(min, max)

		if val < min || val > max {
			t.Errorf("RandomIntBetween(%v, %v) returned %v, which is out of the expected range", min, max, val)
		}
	}
}
