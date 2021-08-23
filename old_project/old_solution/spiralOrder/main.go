package main

import (
	"container/ring"
	"fmt"
)

//https://leetcode.com/problems/spiral-matrix/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	dir := ring.New(4)
	for _, d := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		dir.Value = d
		dir = dir.Next()
	}

	rows := len(matrix)
	cols := len(matrix[0])
	res := make([]int, 0, rows*cols)

	visited := make([][]int, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]int, cols)
	}
	i, j := 0, -1
	for !allVisited(visited) {
		dirValue := dir.Value.([]int)
		i, j = i+dirValue[0], j+dirValue[1]

		if (i >= 0 && i < rows) &&
			(j >= 0 && j < cols) &&
			visited[i][j] == 1 {
			i, j = i-dirValue[0], j-dirValue[1]
			dir = dir.Next()
		}

		if i < 0 {
			i = 0
			dir = dir.Next()
		}

		if j < 0 {
			j = 0
			dir = dir.Next()
		}

		if i >= rows {
			i--
			dir = dir.Next()
		}

		if j >= cols {
			j--
			dir = dir.Next()
		}

		if visited[i][j] == 0 {
			visited[i][j] = 1
			res = append(res, matrix[i][j])
		}
	}

	return res
}

func allVisited(visited [][]int) bool {
	full := true
	for i := 0; i < len(visited); i++ {
		for j := 0; j < len(visited[i]); j++ {
			if visited[i][j] == 0 {
				full = false
			}
		}
	}
	return full
}

//https://leetcode.com/problems/spiral-matrix-ii/
func generateMatrix(n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}

	dir := ring.New(4)
	for _, d := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		dir.Value = d
		dir = dir.Next()
	}

	iter := 1
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	visited := make([][]int, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]int, n)
	}
	i, j := 0, -1
	for !allVisited(visited) {
		dirValue := dir.Value.([]int)
		i, j = i+dirValue[0], j+dirValue[1]

		if (i >= 0 && i < n) &&
			(j >= 0 && j < n) &&
			visited[i][j] == 1 {
			i, j = i-dirValue[0], j-dirValue[1]
			dir = dir.Next()
		}

		if i < 0 {
			i = 0
			dir = dir.Next()
		}

		if j < 0 {
			j = 0
			dir = dir.Next()
		}

		if i >= n {
			i--
			dir = dir.Next()
		}

		if j >= n {
			j--
			dir = dir.Next()
		}

		if visited[i][j] == 0 {
			visited[i][j] = 1
			matrix[i][j] = iter
			iter++
		}
	}

	return matrix
}

func rotate(m [][]int) {
	n := len(m)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := m[i][j]
			// 左边的行 等于 右边的列
			m[i][j] = m[n-j-1][i]
			m[n-j-1][i] = m[n-i-1][n-j-1]
			m[n-i-1][n-j-1] = m[j][n-i-1]
			m[j][n-i-1] = temp
		}
	}
}

func main() {
	//matrix := generateMatrix(3)
	//fmt.Println(matrix)
	//fmt.Println(spiralOrder(matrix))
	//fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}))
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	fmt.Println(matrix)
	rotate(matrix)
	fmt.Println(matrix)
}
