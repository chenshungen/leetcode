package main

import "fmt"

func uniquePaths(m int, n int) int {
	memo := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		memo[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			memo[i][j] = -1
		}
	}
	return uniquePathsHelper(m, n, &memo)
}

func uniquePathsHelper(m int, n int, memo *[][]int) int {
	if m == 1 {
		return 1
	}

	if n == 1 {
		return 1
	}

	if (*memo)[m][n] != -1 {
		return (*memo)[m][n]
	}
	(*memo)[m][n-1] = uniquePathsHelper(m, n-1, memo)
	(*memo)[m-1][n] = uniquePathsHelper(m-1, n, memo)
	return (*memo)[m][n-1] + (*memo)[m-1][n]
}

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

// 动态规划
func uniquePathsDP(m int, n int) int {
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

func testUniquePaths62() {
	fmt.Println(uniquePaths(10, 10))
	fmt.Println(uniquePathsDP(10, 10))
}
