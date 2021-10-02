package problems

func climbStairs(n int) int {
	x, y, z := 0, 0, 1
	for i := 0; i < n; i++ {
		x = y
		y = z
		z = x + y
	}
	return z
}
