package problems

import "testing"

func TestClimbingStairs(t *testing.T) {
	cases := []struct {
		n      int
		result int
	}{
		{2, 2},
		{3, 3},
	}
	for _, c := range cases {
		if result := climbStairs(c.n); result != c.result {
			t.Errorf("expected %v, got %v", c.result, result)
		}
	}
}
