package problems

//Time complexity - O(n*log m), memory - O(1)

func minEatingSpeed(piles []int, h int) int {
	maxInt := 0
	for _, pile := range piles {
		maxInt = max(maxInt, pile)
	}
	result := maxInt + 1
	minInt := 1
	for minInt <= maxInt {
		mid := minInt + (maxInt-minInt)/2
		sumHours := 0
		for _, pile := range piles {
			sumHours += (pile + mid - 1) / mid
		}
		if sumHours > h {
			minInt = mid + 1
		} else {
			result = min(result, mid)
			maxInt = mid - 1
		}
	}
	return result
}
