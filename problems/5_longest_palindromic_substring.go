package problems

//Time complexity - O(n^2), memory - O(1)

func longestPalindrome(s string) string {
	var maxString string
	for i := 0; i < len(s); i++ {
		maxString1 := palindrome(i, i, s)
		maxString2 := palindrome(i, i+1, s)
		if len(maxString1) > len(maxString) {
			maxString = maxString1
		}
		if len(maxString2) > len(maxString) {
			maxString = maxString2
		}
	}
	return maxString
}

func palindrome(left int, right int, s string) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
