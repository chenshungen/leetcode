package main

import "fmt"

func trailingZeroes(n int) int {
	res := 0

	for n >= 5 {
		n /= 5
		res += n
	}

	return res
}

func testTrailingZeroes172() {
	fmt.Println(trailingZeroes(30))
}
