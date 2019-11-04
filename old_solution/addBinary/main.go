package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	aSlice := []byte(a)
	bSlice := []byte(b)
	aLen, bLen := len(a), len(b)
	// 对齐
	if aLen > bLen {
		for i := 0; i < aLen-bLen; i++ {
			bSlice = append([]byte{'0'}, bSlice...)
		}
	} else if aLen < bLen {
		for i := 0; i < bLen-aLen; i++ {
			aSlice = append([]byte{'0'}, aSlice...)
		}
	}

	var tmpRet int
	var tmpAdd int
	res := make([]byte, 0, len(aSlice))
	for i := len(aSlice) - 1; i >= 0; i-- {
		tmpRet = addBin(aSlice[i], bSlice[i]) + tmpAdd
		if i != 0 {
			tmpAdd = tmpRet / 2
			res = append(res, strconv.FormatInt(int64(tmpRet%2), 10)...)
		} else {
			if tmpRet/2 > 0 {
				res = append(res, strconv.FormatInt(int64(tmpRet%2), 10)...)
				tmpRet = tmpRet / 2
			}
			res = append(res, strconv.FormatInt(int64(tmpRet), 10)...)
		}
	}

	reverseString(res)
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

func addBin(a, b byte) int {
	return int(a-'0') + int(b-'0')
}

func main() {
	fmt.Println(addBinary("1111111111", "1111111111"))
}
