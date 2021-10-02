package problems

import "testing"

func TestClimbingStairs(t *testing.T) {
	cases := []struct {
		nums   []int
		result int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
	}
	for _, c := range cases {
		if result := singleNumber(c.nums); result != c.result {
			t.Errorf("expected %v, got %v", c.result, result)
		}
	}
}
