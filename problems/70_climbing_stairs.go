package problems

//Time complexity - O(n), memory - O(n)

func climbStairs(n int) int {
	cache := make(map[int]int)
	return steps(n, 1, cache) + steps(n, 2, cache)
}

func steps(target int, count int, cache map[int]int) int {
	if val, ok := cache[count]; ok {
		return val
	}
	if count == target {
		return 1
	}
	if count > target {
		return 0
	}
	cache[count] = steps(target, count+1, cache) + steps(target, count+2, cache)
	return cache[count]
}
