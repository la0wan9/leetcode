package problems

func reverseWords(s string) string {
	length := len(s)
	rev := []byte{}
	for i := 0; i < length; {
		start := i
		for i < length && s[i] != ' ' {
			i++
		}
		for p := i - 1; p >= start; p-- {
			rev = append(rev, s[p])
		}
		for i < length && s[i] == ' ' {
			rev = append(rev, ' ')
			i++
		}
	}
	return string(rev)
}
