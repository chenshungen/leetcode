package leetcode

import (
	"math"
	"testing"
)

func myAtoi(s string) int {
	sl := len(s)
	if sl == 0 {
		return 0
	}

	// 符号位标记
	neg := 1
	var res []byte
	i := 0

	// 去除前面的空格
	for i < sl && s[i] == ' ' {
		i++
	}

	if i >= sl {
		return 0
	}

	// 获取正负号
	if s[i] == '-' {
		neg = -1
		i++
	} else if s[i] == '+' {
		neg = 1
		i++
	}

	// 跳过前置的0
	for i < sl && s[i] == '0' {
		i++
	}

	// 首位必须是数字
	if i < sl && !(s[i] >= '0' && s[i] <= '9') {
		return 0
	}

	// 开始迭代获取数字,从头往后找['0'-'9']之间的数
	for i < sl {
		// 数字
		if s[i] >= '0' && s[i] <= '9' {
			res = append(res, s[i])
		} else {
			break
		}
		i++
	}

	if len(res) > 10 { // 大于整数位数了
		if neg == -1 {
			return math.MinInt32
		} else if neg == 1 {
			return math.MaxInt32
		}
	}

	// 通过数字序列获得结果
	var ret int
	for i := 0; i < len(res); i++ {
		ret = 10*ret + int(res[i]-'0')
	}

	// 边界判定
	if neg == -1 && (neg*ret) < math.MinInt32 {
		return math.MinInt32
	} else if neg == 1 && (neg*ret) > math.MaxInt32 {
		return math.MaxInt32
	}

	return neg * ret
}

func TestMyAtoi(t *testing.T) {
	t.Log(myAtoi("9223372036854775808"))
}
