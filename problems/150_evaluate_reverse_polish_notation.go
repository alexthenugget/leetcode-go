package problems

import (
	"strconv"
)

//Time complexity - O(n), memory - O(n)

func evalRPN(tokens []string) int {
	var stack []int
	for _, elem := range tokens {
		num, err := strconv.Atoi(elem)
		if err == nil {
			stack = append(stack, num)
		} else {
			num1 := stack[len(stack)-1]
			num2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			result := 0
			switch elem {
			case "+":
				result = num2 + num1
			case "-":
				result = num2 - num1
			case "/":
				result = num2 / num1
			default:
				result = num2 * num1
			}
			stack = append(stack, result)
		}
	}
	return stack[0]
}
