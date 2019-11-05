package main

func twoSum(nums []int, target int) []int {
	var idxs []int
	mMap := make(map[int]int)
	nlen := len(nums)
	for i := 0; i < nlen; i++ {
		if _, ok := mMap[nums[i]]; !ok {
			mMap[target-nums[i]] = i
		} else {
			idxs = append(idxs, mMap[nums[i]], i)
			break
		}
	}
	return idxs
}
