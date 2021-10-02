package problems

func isValid(s string) bool {
	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || s[i] != m[stack[len(stack)-1]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
