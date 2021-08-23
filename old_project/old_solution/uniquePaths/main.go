package main

import (
	"fmt"
)

// 递归耗时太长(当维度越来越大时)
func uniquePaths2(m int, n int) int {
	if m == 1 {
		return 1
	}

	if n == 1 {
		return 1
	}
	return uniquePaths2(m, n-1) + uniquePaths2(m-1, n)
}

// 暴力求解
func uniquePaths3(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	indexs := make([][]int, 0, m)
	count := 0
	i, j := 0, 0
	for i < m {

		if j < n-1 {
			for j < n {
				// 记录每个可以向下转的节点
				if i+1 < m {
					indexs = append(indexs, []int{i, j})
				}
				j++
			}

			if j == n {
				j = n - 1
			}
		}

		// 需要向下走
		if i < m-1 {
			if len(indexs) > 0 {
				indexs = indexs[:len(indexs)-1]
			}
			for i < m {
				i++
			}

			if i == m {
				i = m - 1
			}

		}

		if i == m-1 && j == n-1 {
			count++
		}

		if len(indexs) == 0 {
			break
		}

		lastTurnPoint := indexs[len(indexs)-1]

		i, j = lastTurnPoint[0]+1, lastTurnPoint[1]

		indexs = indexs[:len(indexs)-1]
	}

	return count
}

func uniquePaths(m int, n int) int {
	// path[i][j] 代表了，到达 (i,j) 格子的不同路径数目
	path := [100][100]int{}

	for i := 0; i < m; i++ {
		// 到达第 0 列的格子，只有一条路径
		path[i][0] = 1
	}

	for j := 0; j < n; j++ {
		// 到达第 0 行的格子，只有一条路径
		path[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 到达 (i,j) 格子的路径数目，等于
			// 到达 上方格子 和 左边格子 路径数之和
			path[i][j] = path[i-1][j] + path[i][j-1]
		}
	}

	return path[m-1][n-1]
}

func minPathSum(grid [][]int) int {
	// 已经默认 m 和 n 大于 0 了
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < m; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}

	for j := 1; j < n; j++ {
		dp[0][j] = grid[0][j] + dp[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}

	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	if obstacleGrid[0][0] == 0 {
		dp[0][0] = 1
	}

	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = dp[i-1][0]
		}
	}

	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
func main() {
	fmt.Println(uniquePaths3(10, 10))
	fmt.Println(uniquePaths2(10, 10))
	fmt.Println(uniquePaths(10, 10))
}
