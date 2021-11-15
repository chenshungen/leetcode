package leetcode

import (
	"math"
	"testing"
)

func divide(dividend, divisor int) int {
	if dividend == math.MinInt32 { // 考虑被除数为最小值的情况
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 { // 考虑除数为最小值的情况
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	if dividend == 0 { // 考虑被除数为 0 的情况
		return 0
	}

	// 一般情况，使用二分查找
	// 将所有的正数取相反数，这样就只需要考虑一种情况
	rev := false
	if dividend < 0 {
		dividend = -dividend
		rev = !rev
	}
	if divisor < 0 {
		divisor = -divisor
		rev = !rev
	}

	candidates := []int{divisor}
	for y := divisor; y <= dividend; { // 注意溢出
		y += y
		candidates = append(candidates, y)
	}

	ans := 0
	for i := len(candidates) - 1; i >= 0; i-- {
		if dividend > 0 && candidates[i] <= dividend {
			ans |= 1 << i
			dividend -= candidates[i]
		}
	}
	if rev {
		return -ans
	}
	return ans
}

func TestDivide(t *testing.T) {
	t.Log(divide(100, 5))
}
