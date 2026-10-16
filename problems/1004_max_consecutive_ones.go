package problems

//Time complexity - O(n), memory - O(1)

func longestOnes(nums []int, k int) int {
	maxCount := 0
	left := 0
	zeroCount := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}
		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		maxCount = max(right-left+1, maxCount)
	}
	return maxCount
}
