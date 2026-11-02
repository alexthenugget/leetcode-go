package problems

//Time complexity - O(n), memory - O(n)

func dailyTemperatures(temperatures []int) []int {
	stack := []int{0}
	result := make([]int, len(temperatures))
	for i := 1; i < len(temperatures); i++ {
		curr := temperatures[i]
		for len(stack) > 0 {
			index := stack[len(stack)-1]
			if temperatures[index] >= curr {
				break
			}
			stack = stack[:len(stack)-1]
			result[index] = i - index
		}
		stack = append(stack, i)
	}
	return result
}
