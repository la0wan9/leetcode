package problems

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	cases := []struct {
		strs   []string
		result string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}
	for _, c := range cases {
		if result := longestCommonPrefix(c.strs); result != c.result {
			t.Errorf("expected %v, got %v", c.result, result)
		}
	}
}
