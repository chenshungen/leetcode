package main

import "fmt"

func lengthOfLastWord(s string) int {
	strLen := len(s)
	if strLen <= 0 {
		return 0
	}

	low := 0
	high := strLen - 1
	for low <= high && s[low] == ' ' {
		low++
	}
	for low <= high && s[high] == ' ' {
		high--
	}

	var word string
	for low <= high {
		if s[low] == ' ' {
			for low <= high && s[low] == ' ' {
				low++
			}
			word = ""
		} else {
			word += string(s[low])
			low++
		}
	}

	if word != "" {
		return len(word)
	}
	return 0
}
func main() {
	fmt.Println(lengthOfLastWord("Hello World Golang"))
}
