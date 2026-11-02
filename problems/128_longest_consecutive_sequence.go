package problems

//Time complexity - O(n), memory - O(n)

func longestConsecutive(nums []int) int {
	hashSet := make(map[int]struct{})
	for _, num := range nums {
		hashSet[num] = struct{}{}
	}
	maxLenSeq := 0
	for k, _ := range hashSet {
		lenSeq := 1
		delete(hashSet, k)
		prev := k - 1
		next := k + 1
		for {
			if _, exists := hashSet[prev]; exists {
				lenSeq++
				delete(hashSet, prev)
				prev--
				continue
			}
			break
		}
		for {
			if _, exists := hashSet[next]; exists {
				lenSeq++
				delete(hashSet, next)
				next++
				continue
			}
			break
		}
		maxLenSeq = max(maxLenSeq, lenSeq)
	}
	return maxLenSeq
}
