package problems

//Time complexity - O(log n), memory - O(1)

func findMin(nums []int) int {
	start := 0
	end := len(nums) - 1
	m := 5001
	for start <= end {
		mid := start + (end-start)/2
		m = min(m, nums[mid])
		if nums[mid] < nums[end] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return m
}
