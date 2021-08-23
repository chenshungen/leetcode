package main

import "fmt"

func sortColors(nums []int) {
	tmp := new(int)
	skip := new(bool)
	i := 0
	for i < len(nums) {
		if nums[i] == 0 {
			slide2Right(nums[0:i+1], tmp, skip)
			if *skip {
				i++
			}
		} else if nums[i] == 2 {
			slide2Left(nums[i:], tmp, skip)
			if *skip {
				i++
			}
		} else {
			i++
		}
	}
}

func slide2Right(nums []int, tmp *int, skip *bool) {
	*skip = true
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			*skip = false
			break
		}
	}
	*tmp = nums[len(nums)-1]
	for i := len(nums) - 1; i >= 1; i-- {
		nums[i] = nums[i-1]
	}
	nums[0] = *tmp
}

func slide2Left(nums []int, tmp *int, skip *bool) {
	*skip = true
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			*skip = false
			break
		}
	}
	*tmp = nums[0]
	for i := 0; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
	nums[len(nums)-1] = *tmp
}

func sortColors2(nums []int) {
	count := make([]int, 3)
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}

	i := 0
	for color := 0; color < 3; color++ {
		for j := 0; j < count[color]; j++ {
			nums[i] = color
			i++
		}
	}
}

func main() {
	//nums := []int{1, 2, 1, 0, 1}
	nums := []int{2, 0, 2, 1, 1, 1, 2, 1, 0, 1, 1, 0}
	fmt.Println(nums)
	sortColors(nums)
	fmt.Println(nums)
	nums = []int{2, 0, 2, 1, 1, 1, 2, 1, 0, 1, 1, 0}
	sortColors2(nums)
	fmt.Println(nums)
}
