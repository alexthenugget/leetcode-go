package problems

//Time complexity - O(m*n), memory - O(m*n)

func numIslands(grid [][]byte) int {
	n := len(grid[0])
	m := len(grid)
	visited := make([][]bool, m)
	for k := 0; k < m; k++ {
		visited[k] = make([]bool, n)
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				if !visited[i][j] {
					count++
					dfs(grid, &visited, i, j)
				}
			}
		}
	}
	return count
}

func dfs(grid [][]byte, visited *[][]bool, i int, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}
	if grid[i][j] == '1' {
		if (*visited)[i][j] {
			return
		}

		(*visited)[i][j] = true
	}
	dfs(grid, visited, i+1, j)
	dfs(grid, visited, i, j+1)
	dfs(grid, visited, i-1, j)
	dfs(grid, visited, i, j-1)
	return
}
