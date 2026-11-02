package problems

//Time complexity - O(n*k), memory - O(n*k)

func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[[26]int][]string)
	for _, str := range strs {
		var key [26]int
		for _, r := range str {
			key[r-'a']++
		}
		hashMap[key] = append(hashMap[key], str)
	}
	var result [][]string
	for _, v := range hashMap {
		result = append(result, v)
	}
	return result
}
