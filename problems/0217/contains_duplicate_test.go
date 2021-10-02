package problems

import "testing"

func TestContainsDuplicate(t *testing.T) {
	cases := []struct {
		nums   []int
		result bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}
	for _, c := range cases {
		if result := containsDuplicate(c.nums); result != c.result {
			t.Errorf("expected %v, got %v", c.result, result)
		}
	}
}
