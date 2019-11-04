package main

import "fmt"

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	iMap := make(map[int]bool)
	jMap := make(map[int]bool)
	row, col := len(matrix), len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				iMap[i] = true
				jMap[j] = true
			}
		}
	}

	for ti := range iMap {
		for tj := 0; tj < col; tj++ {
			matrix[ti][tj] = 0
		}
	}

	for tj := range jMap {
		for ti := 0; ti < row; ti++ {
			matrix[ti][tj] = 0
		}
	}
}

// setZeroes2 常数空间解决方案
func setZeroes2(mat [][]int) {
	m, n := len(mat), len(mat[0])
	col0 := 1

	/**
	 * 从上往下，从左往右 扫描矩阵
	 * 利用 mat[i][0] = 0 表示，第 i 行中含有 0
	 * 利用 mat[0][j] = 0 表示，第 j 列中含有 0
	 * 特别地，
	 * mat[0][0] = 0 仅表示，第 0 行中含有 0
	 * Col0 = 0 表示，第 0 列中含有 0
	 * NOTICE: 循环的顺序很重要
	 * 需要保证 mat[i][0] 和 mat[0][j] 被标记后，不再做为别的标记的依据
	 * 要不然的话，标记有可能会被污染
	 */
	for i := 0; i < m; i++ {
		if mat[i][0] == 0 {
			col0 = 0
		}
		for j := 1; j < n; j++ {
			if mat[i][j] == 0 {
				mat[i][0] = 0
				mat[0][j] = 0
			}
		}
	}

	/**
	 * 从下往上，从右往左 扫描矩阵
	 * 并根据前面的标记修改 mat[i][j] 的值
	 * NOTICE: 循环的顺序很重要
	 * 需要保证 mat[i][0] 和 mat[0][j] 被修改后，不再做为别的修改的标记
	 * 要不然的话，标记有可能会被污染
	 */
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 1; j-- {
			if mat[i][0] == 0 || mat[0][j] == 0 {
				mat[i][j] = 0
			}
		}
		if col0 == 0 {
			mat[i][0] = 0
		}
	}

}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	if len(matrix[0]) == 0 {
		return false
	}

	row, col := len(matrix), len(matrix[0])
	if target < matrix[0][0] || target > matrix[row-1][col-1] {
		return false
	}

	nums := make([]int, 0, row*col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			nums = append(nums, matrix[i][j])
		}
	}

	return binSearch(nums, target)
}

func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}

	row, col := len(matrix), len(matrix[0])
	if target < matrix[0][0] || target > matrix[row-1][col-1] {
		return false
	}

	targetIndex := row - 1
	for i := 1; i < row; i++ {
		if target < matrix[i][0] {
			targetIndex = i - 1
			break
		}
	}

	return binSearch(matrix[targetIndex], target)
}

func binSearch(nums []int, target int) bool {
	low := 0
	mid := 0
	high := len(nums) - 1
	for low <= high {
		mid = (high + low) >> 1
		if target < nums[mid] {
			high = mid - 1
		} else if target > nums[mid] {
			low = mid + 1
		} else if target == nums[mid] {
			return true
		}
	}
	return false
}

func main() {
	//matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	//matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	// matrix := [][]int{{0, 1, 2, 3}, {3, 5, 6, 7}, {8, 9, 10, 11}}
	// fmt.Println(matrix)
	// setZeroes(matrix)
	// fmt.Println(matrix)
	//matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	matrix := [][]int{{-10}, {-7}, {-5}}
	fmt.Println(searchMatrix(matrix, -10))
	fmt.Println(searchMatrix2(matrix, -10))
}
