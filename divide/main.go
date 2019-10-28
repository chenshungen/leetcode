package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	var (
		count int
		sum   int
		da    int
		db    int
		isNeg bool
	)

	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		isNeg = true
	}

	if dividend < 0 {
		da = -dividend
	} else {
		da = dividend
	}
	if divisor < 0 {
		db = -divisor
	} else {
		db = divisor
	}

	for {
		sum += db
		if sum <= da {
			count++
		} else {
			break
		}
	}

	if count > math.MaxInt32 && !isNeg {
		return math.MaxInt32
	} else if isNeg && int64(count*(-1)) < int64(math.MinInt32) {
		return math.MinInt32
	}
	if isNeg {
		return int(count * (-1))
	}
	return count
}

func main() {
	fmt.Println(divide(-2147483648, -1))
}
