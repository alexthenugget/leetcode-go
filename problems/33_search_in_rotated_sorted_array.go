package problems

//Time complexity - O(log n), memory - O(1)

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		}
		if nums[start] <= nums[mid] {
			if target < nums[mid] && nums[start] <= target {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
