package main

import "fmt"

func plusOne(digits []int) []int {
	add, ret := -1, 0
	for i := len(digits) - 1; i >= 0; i-- {
		if add == -1 {
			ret = digits[i] + 1
		} else {
			ret = digits[i] + add
		}

		if ret >= 10 {
			if i == 0 {
				digits[i] = ret % 10
				digits = append([]int{ret / 10}, digits...)
			} else {
				digits[i] = ret % 10
				add = ret / 10
			}
		} else {
			digits[i] = ret
			add = 0
		}

	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 9, 9, 9}))
}
