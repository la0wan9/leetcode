package problems

import "testing"

func TestReverseInteger(t *testing.T) {
	cases := []struct {
		x      int
		result int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{0, 0},
	}
	for _, c := range cases {
		if result := reverse(c.x); result != c.result {
			t.Errorf("expected %v, got %v", c.result, result)
		}
	}
}
