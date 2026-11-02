package problems

//Time complexity - O(n), memory - O(n)

func isValid(s string) bool {
	stack := []rune{}
	var elem rune
	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			stack = append(stack, r)
			continue
		}
		if len(stack) != 0 {
			elem = stack[len(stack)-1]
		} else {
			elem = '0'
		}
		if elem == '(' && r == ')' {
			stack = stack[:len(stack)-1]
			continue
		}
		if elem == '[' && r == ']' {
			stack = stack[:len(stack)-1]
			continue
		}
		if elem == '{' && r == '}' {
			stack = stack[:len(stack)-1]
			continue
		}
		return false
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
