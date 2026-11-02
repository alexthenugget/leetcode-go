package problems

//Time complexity - O(n), memory - O(n)

func firstUniqChar(s string) int {
	st := []rune(s)
	m := make(map[rune]int)
	for _, r := range st {
		m[r]++
	}

	for i, r := range st {
		if m[r] == 1 {
			return i
		}
	}
	return -1
}
