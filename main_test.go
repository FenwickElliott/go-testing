package main

import "testing"

func TestSum(t *testing.T) {
	res := Sum(9, 8)
	if res != 9+8 {
		t.Errorf("Sum failed. Got: %d, Expected: %d.", res, 17)
	}
}
