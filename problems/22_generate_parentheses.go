package problems

//Time complexity - O(n * C_n), memory - O(n * C_n)
//C_n = (2*n)! / ((n+1)! * n!)

func generateParenthesis(n int) []string {
	var result []string
	generate(0, 0, n, "", &result)
	return result
}

func generate(open, close, n int, curr string, result *[]string) {
	if open == n && close == n {
		*result = append(*result, curr)
		return
	}
	if open < n {
		generate(open+1, close, n, curr+"(", result)
	}
	if close < open {
		generate(open, close+1, n, curr+")", result)
	}
}
