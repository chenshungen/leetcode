package leetcode

import (
	"testing"
)

// 1994 -> MCMXCIV
func intToRoman(num int) string {
	romanKey := []string{"I", "II", "III", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	romanValue := []int{1, 2, 3, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var res string
	for i := len(romanKey) - 1; i >= 0; i-- {
		for num >= romanValue[i] {
			res += romanKey[i]
			num = num - romanValue[i]
		}
	}
	return res
}

func TestIntToRoman(t *testing.T) {
	t.Log(intToRoman(9))
}
