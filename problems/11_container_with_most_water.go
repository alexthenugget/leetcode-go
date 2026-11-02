package problems

//Time complexity - O(n), memory - O(1)

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxSum := 0
	for left < right {
		currSum := min(height[left], height[right]) * (right - left)
		maxSum = max(maxSum, currSum)
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return maxSum
}
