package problems

//Time complexity - O(n), memory - O(1)

func minSubArrayLen(target int, nums []int) int {
	left := 0
	right := 0
	sum := 0
	minLen := len(nums) + 1
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			minLen = min(minLen, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}
	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}
