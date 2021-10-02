package problems

import "testing"

func TestValidParentheses(t *testing.T) {
	cases := []struct {
		s      string
		result bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}
	for _, c := range cases {
		if result := isValid(c.s); result != c.result {
			t.Errorf("expected %v, got %v", c.result, result)
		}
	}
}
