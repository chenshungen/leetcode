package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0, len(nums))
	for i := 0; i <= len(nums); i++ {
		tmpRes := combinations(len(nums), i)
		for ti := 0; ti < len(tmpRes); ti++ {
			tmp := make([]int, 0, 0)
			for tj := 0; tj < len(tmpRes[ti]); tj++ {
				tmp = append(tmp, nums[tmpRes[ti][tj]-1])
			}

			if !isExist(tmp, res) {
				res = append(res, tmp)
			}
		}
	}
	return res
}

func isExist(ret []int, res [][]int) bool {
	sort.Sort(sort.IntSlice(ret))
	for _, tmp := range res {
		if len(tmp) != len(ret) {
			continue
		}

		sort.Sort(sort.IntSlice(tmp))
		count := 0
		for i := 0; i < len(ret); i++ {
			if ret[i] == tmp[i] {
				count++
			}
		}

		if count == len(ret) {
			return true
		}

	}

	return false
}

// C(n, k) = n!/k!(n-k)!
func binomial2(n, k int) int {
	if n < 0 || k < 0 {
		return 0
	}
	if n < k {
		return 0
	}
	return factorial(n) / (factorial(k) * factorial(n-k))
}

func factorial(n int) int {
	facVal := 1
	if n < 0 {
		return 1
	}

	for i := 1; i <= n; i++ {
		facVal *= i
	}
	return facVal

}

func combinations(n, k int) [][]int {
	combins := binomial2(n, k)
	data := make([][]int, combins)
	if len(data) == 0 {
		return data
	}
	data[0] = make([]int, k)
	for i := range data[0] {
		data[0][i] = i + 1
	}
	for i := 1; i < combins; i++ {
		next := make([]int, k)
		copy(next, data[i-1])
		nextCombination(next, n, k)
		data[i] = next
	}
	return data
}

// nextCombination generates the combination after s, overwriting the input value.
func nextCombination(s []int, n, k int) {
	for j := k - 1; j >= 0; j-- {
		if s[j] == n+j-k+1 {
			continue
		}
		s[j]++
		for l := j + 1; l < k; l++ {
			s[l] = s[j] + l - j
		}
		break
	}
}

func main() {
	// // fmt.Println(combine(4, 2))
	// // fmt.Println(combinations(4, 2))
	// nums := []int{1, 1, 2}
	// // fmt.Println(subsets(nums))
	// fmt.Println(permutations(nums))
	// nums = []int{1, 1, 2}
	// fmt.Println(permute(nums))

	nums := []int{4, 4, 4, 1, 4}
	fmt.Println(subsetsWithDup(nums))
}
