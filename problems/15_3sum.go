package problems

//Time complexity - O(n^2), memory - O(n^2)

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	index := 0
	left := index + 1
	right := len(nums) - 1
	var result [][]int
	for index < len(nums)-1 {
		if index > 0 && nums[index] == nums[index-1] {
			index++
			left = index + 1
			continue
		}
		for left < right {
			if right < len(nums)-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			if left > index+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			sum := nums[index] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[index], nums[left], nums[right]})
				left++
				right--
			} else {
				if sum > 0 {
					right--
				} else {
					left++
				}
			}
		}
		index++
		left = index + 1
		right = len(nums) - 1
	}
	return result
}
