package problems

//Time complexity - O(n), memory - O(1)

func maxDistToClosest(seats []int) int {
	maxDist := 0
	left := 0
	right := 0
	for right < len(seats) {
		if seats[left] == 0 && seats[right] == 1 {
			maxDist = max(maxDist, right-left)
			left = right
		}
		if seats[left] == 1 && seats[right] == 1 {
			maxDist = max(maxDist, (right-left)/2)
			left = right
		}
		right++
	}
	maxDist = max(maxDist, right-left-1)
	return maxDist
}
