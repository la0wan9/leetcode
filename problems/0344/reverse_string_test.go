package problems

import "testing"

func TestContainsDuplicate(t *testing.T) {
	cases := []struct {
		s      []byte
		result []byte
	}{
		{
			[]byte{'h', 'e', 'l', 'l', 'o'},
			[]byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			[]byte{'H', 'a', 'n', 'n', 'a', 'h'},
			[]byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}
	for _, c := range cases {
		result := c.s
		reverseString(c.s)
		if string(result) != string(c.result) {
			t.Errorf("expected %v, got %v", c.result, result)
		}
	}
}
