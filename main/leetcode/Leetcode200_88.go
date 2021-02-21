package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	visited := make([][]bool, len(grid))
	count := 0
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	for i := range visited {
		for j := range visited[0] {
			if !visited[i][j] && grid[i][j] == '1' {
				count++
				dfs(i, j, visited, grid)
			}
		}
	}
	return count
}

func dfs(i int, j int, visited [][]bool, grid [][]byte) {
	if !visited[i][j] && grid[i][j] == '1' {
		visited[i][j] = true
		dfs(i+1, j, visited, grid)
		dfs(i, j+1, visited, grid)
		dfs(i-1, j, visited, grid)
		dfs(i, j-1, visited, grid)
	}
}
