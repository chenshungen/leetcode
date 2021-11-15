package leetcode

import "testing"

func letterCombinations(digits string) []string {

	lookup := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	var res []string
	for i := range digits {
		dc, ok := lookup[digits[i]]
		if !ok {
			return []string{}
		}

		if len(res) == 0 {
			res = append(res, dc...)
		} else {
			var tmpRes []string
			for _, s1 := range res {
				for _, s2 := range dc {
					tmpRes = append(tmpRes, s1+s2)
				}
			}
			res = tmpRes
		}
	}
	return res
}

func TestLetterCombinations(t *testing.T) {
	t.Log(letterCombinations("2467"))
}
