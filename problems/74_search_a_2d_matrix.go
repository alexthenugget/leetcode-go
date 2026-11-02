package problems

//Time complexity - O(log(m*n)), memory - O(1)

func searchMatrix(matrix [][]int, target int) bool {
	rowLen := len(matrix[0])
	start := 0
	end := len(matrix)*rowLen - 1
	for start <= end {
		mid := start + (end-start)/2
		i := mid / rowLen
		j := mid % rowLen
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false
}
