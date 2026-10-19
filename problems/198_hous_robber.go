package problems

//Time complexity - O(n), memory - O(1)

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	i := 2
	prevPrev := nums[0]
	prev := max(nums[0], nums[1])
	for i < len(nums) {
		curr := max(prev, prevPrev+nums[i])
		prevPrev, prev = prev, curr
		i++
	}
	return prev
}
