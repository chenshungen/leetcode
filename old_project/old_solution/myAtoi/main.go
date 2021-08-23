package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/string-to-integer-atoi/
func myAtoi(str string) int {
	strLen := len(str)
	if strLen <= 0 {
		return 0
	}
	var (
		isNeg bool
		re    int64
		i     int
		cnt   int
	)

	// 跳过开头所有空格
	for ; i < len(str); i++ {
		if str[i] != ' ' {
			break
		}
	}

	if i > len(str)-1 {
		return 0
	}

	// 正负数判定
	if str[i] == '-' {
		isNeg = true
		i++
	} else if str[i] == '+' {
		isNeg = false
		i++
	}

	// 跳过非数字
	for ; i < len(str); i++ {
		if str[i] != '0' {
			break
		}
	}

	if i > len(str)-1 {
		return 0
	}

	// 非数字开头的直接返回0
	if !(str[i] >= '0' && str[i] <= '9') {
		return 0
	}

	for ; i < len(str); i++ {
		if cnt > 10 { //大于整数所表示的最大位数 跳出循环
			break
		}
		if str[i] < '0' || str[i] > '9' {
			break
		}
		re = re*10 + int64((str[i] - '0'))
		cnt++
	}

	if re > math.MaxInt32 && !isNeg {
		return math.MaxInt32
	} else if isNeg && re*(-1) < int64(math.MinInt32) {
		return math.MinInt32
	}
	if isNeg {
		return int(re * (-1))
	}
	return int(re)

}

func main() {
	fmt.Println(myAtoi("  41313 fafa afa "))
}
