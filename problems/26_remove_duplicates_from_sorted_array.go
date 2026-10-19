package problems

//Time complexity - O(n), memory - O(1)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	left := 0
	right := 1
	count := 1
	for right < len(nums) {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
			count++
		}
		right++
	}
	return count
}
