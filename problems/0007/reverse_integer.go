package problems

func reverse(x int) int {
	rev := 0
	for x != 0 {
		rev = rev*10 + x%10
		x /= 10
	}
	if rev > 1<<31-1 || rev < -(1<<31) {
		return 0
	}
	return rev
}
