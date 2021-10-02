package problems

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	rev := 0
	for v := x; v != 0; v /= 10 {
		rev = rev*10 + v%10
	}
	return rev == x
}
