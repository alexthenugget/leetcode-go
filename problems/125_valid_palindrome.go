package problems

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)
	regex, _ := regexp.Compile(`[^A-Za-zА-Яа-я0-9]`)
	rs := regex.ReplaceAllString(s, "")
	runes := []rune(rs)
	length := len(runes)
	for i := 0; i < length-1; i++ {
		if runes[i] != runes[length-1-i] {
			return false
		}
	}
	return true
}
