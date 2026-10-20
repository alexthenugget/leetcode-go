package problems

//Time complexity - O(m*n), memory - O(1)

func minPathSum(grid [][]int) int {
	for j := 0; j < len(grid[0]); j++ {
		for i := 0; i < len(grid); i++ {
			switch {
			case i == 0 && j == 0:
				continue
			case i == 0:
				grid[i][j] += grid[i][j-1]
			case j == 0:
				grid[i][j] += grid[i-1][j]
			default:
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
