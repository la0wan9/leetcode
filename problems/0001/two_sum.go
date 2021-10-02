package problems

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if another, ok := m[target-num]; ok {
			return []int{another, i}
		}
		m[num] = i
	}
	return nil
}
