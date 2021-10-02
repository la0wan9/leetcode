package problems

import "testing"

func TestPalindromeNumber(t *testing.T) {
	cases := []struct {
		x      int
		result bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{-101, false},
	}
	for _, c := range cases {
		if result := isPalindrome(c.x); result != c.result {
			t.Errorf("expected: %v, got: %v", c.result, result)
		}
	}
}
