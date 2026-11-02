package problems

//Time complexity - O(n), memory - O(1)

func longestSubarray(nums []int) int {
	left := 0
	right := 0
	maxLen := 0
	lastZero := -1
	for right < len(nums) {
		if nums[right] == 0 {
			if lastZero != -1 {
				left = lastZero
			}
			lastZero = right + 1
		}
		maxLen = max(maxLen, right-left)
		right++
	}
	return maxLen
}
