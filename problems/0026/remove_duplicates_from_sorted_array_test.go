package problems

import "testing"

func TestValidParentheses(t *testing.T) {
	cases := []struct {
		nums   []int
		result int
	}{
		{[]int{1, 1, 3}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}
	for _, c := range cases {
		if result := removeDuplicates(c.nums); result != c.result {
			t.Errorf("expected %v, got %v", c.result, result)
		}
	}
}
