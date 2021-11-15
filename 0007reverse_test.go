package leetcode

import (
	"math"
	"testing"
)

func reverse(x int) int {
	// 是否为负数
	neg := 1
	if x < 0 {
		neg = -1
	}
	x = neg * x
	// x = 123
	res := make([]int, 0, 10)
	for x != 0 {
		res = append(res, x%10)
		x = x / 10
	}

	var ret int
	for i := 0; i < len(res); i++ {
		ret = 10*ret + res[i]
	}

	if neg == -1 && math.Abs(float64(ret)) > math.Abs(float64(math.MinInt32)) {
		return 0
	}

	if neg == 1 && ret > int(math.MaxInt32) {
		return 0
	}

	return neg * ret
}

func TestReverse(t *testing.T) {
	t.Log(reverse(-4999999999999999))
}
