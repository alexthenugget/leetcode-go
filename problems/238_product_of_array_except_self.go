package problems

//Time complexity - O(n), memory - O(n)

func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	suffix := 1
	output[0] = 1
	for i := 1; i < len(nums); i++ {
		output[i] = output[i-1] * nums[i-1]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		output[i] = output[i] * suffix
		suffix *= nums[i]
	}
	return output
}
