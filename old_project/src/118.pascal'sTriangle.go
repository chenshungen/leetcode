package main

import "fmt"

func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for i := 1; i <= numRows; i++ {
		triangle[i-1] = make([]int, i)
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 || j == len(triangle[i])-1 {
				triangle[i][j] = 1
			} else {
				triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
			}
		}
	}

	return triangle
}

func getRow(rowIndex int) []int {
	res := make([]int, 1, rowIndex+1)
	res[0] = 1
	if rowIndex == 0 {
		return res
	}

	for i := 0; i < rowIndex; i++ {
		res = append(res, 1)
		for j := len(res) - 2; j > 0; j-- {
			res[j] += res[j-1]
		}
	}

	return res
}

func testGeneratePascalTriangle118() {
	fmt.Println(generate(5))
	fmt.Println(getRow(3))
}
