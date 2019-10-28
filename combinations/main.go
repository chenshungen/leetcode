package main

import (
	"fmt"
	"sort"
)

func combine(n int, k int) [][]int {
	combination := make([]int, k)
	res := [][]int{}

	var dfs func(int, int)
	dfs = func(idx, begin int) {
		if idx == k {
			temp := make([]int, k)
			copy(temp, combination)
			res = append(res, temp)
			return
		}

		for i := begin; i <= n+1-k+idx; i++ {
			combination[idx] = i
			dfs(idx+1, i+1)
		}
	}

	dfs(0, 1)

	return res
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0, len(nums))
	for i := 0; i <= len(nums); i++ {
		tmpRes := combine(len(nums), i)
		for ti := 0; ti < len(tmpRes); ti++ {
			tmp := make([]int, 0, 0)
			for tj := 0; tj < len(tmpRes[ti]); tj++ {
				tmp = append(tmp, nums[tmpRes[ti][tj]-1])
			}
			res = append(res, tmp)
		}
	}
	return res
}

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
	for _, tmp := range res {
		if len(tmp) != len(ret) {
			continue
		}

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

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func permute(arr []int) [][]int {
	arrLen := len(arr)
	if arrLen == 0 {
		return [][]int{}
	}

	sort.Sort(sort.IntSlice(arr))

	permuteNum := factorial(len(arr))
	res := make([][]int, 0, permuteNum)

	next := make([]int, len(arr))
	copy(next, arr)
	res = append(res, next)

	for {
		next = nextPermute(next)
		if next == nil {
			break
		}
		res = append(res, next)
	}

	return res
}

func nextPermute(arr []int) []int {
	var (
		j int
		k int
	)
	ret := make([]int, len(arr))
	copy(ret, arr)
	for j = len(ret) - 1; j > 0; j-- {
		if ret[j] > ret[j-1] {
			break
		}
	}

	if j <= 0 {
		return nil
	}

	j = j - 1

	for k = len(ret) - 1; k >= 0; k-- {
		if ret[k] > ret[j] {
			break
		}
	}

	ret[j], ret[k] = ret[k], ret[j]

	low := j + 1
	high := len(ret) - 1
	for low <= high {
		ret[low], ret[high] = ret[high], ret[low]
		low++
		high--
	}

	return ret
}

func nextPermutation(nums []int) {
	var (
		j int
		k int
	)

	for j = len(nums) - 1; j > 0; j-- {
		if nums[j] > nums[j-1] {
			break
		}
	}

	if j <= 0 {
		sort.Sort(sort.IntSlice(nums))
		return
	}

	j = j - 1

	for k = len(nums) - 1; k >= 0; k-- {
		if nums[k] > nums[j] {
			break
		}
	}

	nums[j], nums[k] = nums[k], nums[j]

	low := j + 1
	high := len(nums) - 1
	for low <= high {
		nums[low], nums[high] = nums[high], nums[low]
		low++
		high--
	}
}

func binomial(n, k int) int {
	if n < 0 || k < 0 {
		return 0
	}
	if n < k {
		return 0
	}
	// (n,k) = (n, n-k)
	if k > n/2 {
		k = n - k
	}
	b := 1
	for i := 1; i <= k; i++ {
		b = (n - k + i) * b / i
	}
	return b
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
