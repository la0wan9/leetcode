package problems

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		if prefix = lcp(prefix, strs[i]); prefix == "" {
			return ""
		}
	}
	return prefix
}

func lcp(x, y string) string {
	minLen := len(x)
	if l := len(y); l < minLen {
		minLen = l
	}
	index := 0
	for index < minLen && x[index] == y[index] {
		index++
	}
	return x[:index]
}
