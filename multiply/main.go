package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/multiply-strings/
func multiply(num1 string, num2 string) string {

	if len(num1) == 0 || len(num2) == 0 {
		return "0"
	}

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	steps := make([][]byte, 0, len(num2))
	for i := len(num2) - 1; i >= 0; i-- {
		var tmpRet int64
		var tmpAdd int64
		var step []byte
		for j := len(num1) - 1; j >= 0; j-- {
			tmpRet = int64(num2[i]-'0')*int64(num1[j]-'0') + tmpAdd
			if j == 0 {
				for tmpRet/10 > 0 {
					step = append(step, []byte(strconv.FormatInt(tmpRet%10, 10))...)
					tmpRet = tmpRet / 10
				}
				step = append(step, []byte(strconv.FormatInt(tmpRet%10, 10))...)
			} else {
				tmpAdd = tmpRet / 10
				step = append(step, []byte(strconv.FormatInt(tmpRet%10, 10))...)
			}
		}

		reverseString(step)
		//fmt.Println(string(step))
		if len(steps) > 0 {
			for i := 0; i < len(steps); i++ {
				step = append(step, '0')
			}
		}

		steps = append(steps, step)
	}

	lastLen := len(steps[len(steps)-1])
	for i := 0; i < len(steps)-1; i++ {
		step := steps[i]
		stepLen := len(step)
		for k := 0; k < lastLen-stepLen; k++ {
			step = append([]byte{'0'}, step...)
		}
		steps[i] = step
	}

	// for _, step := range steps {
	// 	fmt.Println(string(step))
	// }

	var tmp int64
	var res []byte
	stepLen := len(steps[0])

	for i := stepLen - 1; i >= 0; i-- {
		var sum int64
		for k := 0; k < len(steps); k++ {
			sum += int64(steps[k][i] - '0')
		}
		sum += tmp

		if sum >= 10 {
			tmp = sum / 10
		}

		if i != 0 {
			tmp = sum / 10
			res = append(res, strconv.FormatInt(sum%10, 10)...)
		} else {
			for sum/10 > 0 {
				res = append(res, strconv.FormatInt(sum%10, 10)...)
				sum = sum / 10
			}
			res = append(res, strconv.FormatInt(sum, 10)...)
		}
	}
	reverseString(res)
	//fmt.Println(string(res))

	return string(res)
}

func reverseString(s []byte) {
	low := 0
	high := len(s) - 1

	for low < high {
		s[low], s[high] = s[high], s[low]
		low++
		high--
	}
}

func main() {
	fmt.Println(multiply("15", "15"))
}
