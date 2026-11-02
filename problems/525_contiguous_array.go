package problems

//Time complexity - O(n), memory - O(n)

func findMaxLength(nums []int) int {
	currSum := 0
	hashMap := make(map[int]int)
	maxLen := 0
	hashMap[0] = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			currSum--
		} else {
			currSum++
		}
		if index, exists := hashMap[currSum]; exists {
			maxLen = max(i-index, maxLen)
		} else {
			hashMap[currSum] = i
		}
	}
	return maxLen
}
