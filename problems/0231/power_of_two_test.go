package problems

import "testing"

func TestContainsDuplicate(t *testing.T) {
	cases := []struct {
		n      int
		result bool
	}{
		{1, true},
		{16, true},
		{3, false},
		{4, true},
		{5, false},
	}
	for _, c := range cases {
		if result := isPowerOfTwo(c.n); result != c.result {
			t.Errorf("expected %v, got %v", c.result, result)
		}
	}
}
