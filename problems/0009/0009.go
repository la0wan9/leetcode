package problems

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	rev := 0
	for t := x; t != 0; t /= 10 {
		rev = rev*10 + t%10
	}
	return x == rev
}
