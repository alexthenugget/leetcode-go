package problems

// Time complexity - O(n), memory - O(n)

import "strconv"

func summaryRanges(nums []int) []string {
	var result []string
	if len(nums) == 0 {
		return result
	}
	start, end := nums[0], nums[0]
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+1 != nums[i+1] {
			end = nums[i]
			if start == end {
				result = append(result, strconv.Itoa(start))
			} else {
				result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(end))
			}
			start = nums[i+1]
		}
	}
	end = nums[len(nums)-1]
	if start == end {
		result = append(result, strconv.Itoa(start))
	} else {
		result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(end))
	}
	return result
}
