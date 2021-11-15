package leetcode

import "testing"

func romanToInt(s string) int {
	romanLookup := map[string]int{
		"I":   1,
		"II":  2,
		"III": 3,
		"IV":  4,
		"V":   5,
		"IX":  9,
		"X":   10,
		"XL":  40,
		"L":   50,
		"XC":  90,
		"C":   100,
		"CD":  400,
		"D":   500,
		"CM":  900,
		"M":   1000,
	}
	var (
		res int
		v   int
		ok  bool
	)

	sl := len(s)
	for sl > 0 {
		if sl >= 2 {
			v, ok = romanLookup[s[:2]]
			if !ok {
				v, ok = romanLookup[s[:1]]
				s = s[1:]
			} else {
				s = s[2:]
			}
		} else if sl >= 1 {
			v, ok = romanLookup[s[:1]]
			s = s[1:]
		}

		if ok {
			res += v
		}

		sl = len(s)
	}

	return res
}

func TestRomanToInt(t *testing.T) {
	t.Log(romanToInt("MCMXCIV"))
}
