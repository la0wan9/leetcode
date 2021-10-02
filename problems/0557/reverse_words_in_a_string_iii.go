package problems

func reverseWords(s string) string {
	length := len(s)
	result := make([]byte, 0, length)
	for i := 0; i < length; {
		start := i
		for i < length && s[i] != ' ' {
			i++
		}
		for p := i - 1; p >= start; p-- {
			result = append(result, s[p])
		}
		for i < length && s[i] == ' ' {
			result = append(result, ' ')
			i++
		}
	}
	return string(result)
}
