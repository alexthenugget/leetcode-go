package problems

//Time complexity - O(n), memory - O(1)

func characterReplacement(s string, k int) int {
	hashMap := make(map[byte]int)
	left := 0
	right := 0
	maxSymbol := 0
	maxLen := 0
	for right <= len(s)-1 {
		hashMap[s[right]]++
		maxSymbol = max(maxSymbol, hashMap[s[right]])
		for (right-left+1)-maxSymbol > k {
			hashMap[s[left]]--
			left++
		}
		maxLen = max(maxLen, right-left+1)
		right++
	}
	return maxLen
}
