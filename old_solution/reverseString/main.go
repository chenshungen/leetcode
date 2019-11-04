package main

import "fmt"

// https://leetcode.com/problems/reverse-string/
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
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(string(s))
}
