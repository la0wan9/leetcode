package problems

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReverseInteger(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		result []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, c := range cases {
		if result := twoSum(c.nums, c.target); !cmp.Equal(result, c.result) {
			t.Errorf("expected %v, got %v", c.result, result)
		}
	}
}
