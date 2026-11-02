package problems

//Time complexity - O(n), memory - O(n)

func subarraySum(nums []int, k int) int {
	hashMap := make(map[int]int)
	hashMap[0] = 1
	currSum := 0
	count := 0
	for _, num := range nums {
		currSum += num
		count += hashMap[currSum-k]
		hashMap[currSum]++
	}
	return count
}
