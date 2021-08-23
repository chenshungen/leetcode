package main

import "fmt"

func climbStairs(n int) int {
	if n < 2 {
		return n
	}
	res := make([]int, n+1)
	res[0] = 1
	res[1] = 1
	for i := 2; i <= n; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res[n]
}

func main() {
	fmt.Println(climbStairs(4))
}
