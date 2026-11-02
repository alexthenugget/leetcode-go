package problems

//Time complexity - O(n), memory - O(1)

func lengthOfLongestSubstring(s string) int {
	left := 0
	right := 0
	currLen := 0
	maxLen := 0
	currSet := make(map[byte]struct{})
	for right <= len(s)-1 {
		if _, exists := currSet[s[right]]; exists {
			maxLen = max(maxLen, currLen)
			for s[left] != s[right] {
				delete(currSet, s[left])
				left++
				currLen--
			}
			left++
			right++
		} else {
			currSet[s[right]] = struct{}{}
			right++
			currLen++
		}
	}
	maxLen = max(maxLen, currLen)
	return maxLen
}
