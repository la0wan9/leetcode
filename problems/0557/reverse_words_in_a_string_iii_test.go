package problems

import "testing"

func TestReverseWordsInAStringIII(t *testing.T) {
	cases := []struct {
		s      string
		result string
	}{
		{
			"Let's take LeetCode contest",
			"s'teL ekat edoCteeL tsetnoc",
		},
	}
	for _, c := range cases {
		if result := reverseWords(c.s); result != c.result {
			t.Errorf("expected %v, got %v", c.result, result)
		}
	}
}
