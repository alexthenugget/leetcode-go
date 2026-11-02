package problems

//Time complexity - O(n+m), memory - O(1)

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	s1Array := [26]int{}
	s2Array := [26]int{}
	for i, _ := range s1 {
		s1Array[s1[i]-'a']++
		s2Array[s2[i]-'a']++
	}
	if s1Array == s2Array {
		return true
	}
	right := len(s1)
	for right < len(s2) {
		s2Array[s2[right-len(s1)]-'a']--
		s2Array[s2[right]-'a']++
		if s1Array == s2Array {
			return true
		}
		right++
	}
	return false
}
