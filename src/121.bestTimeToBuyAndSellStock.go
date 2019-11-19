package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}

func maxProfit2(prices []int) int {
	var num = 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			num += prices[i+1] - prices[i]
		}
	}
	return num

}

func maxProfit3(prices []int) int {
	size := len(prices)
	if size <= 1 {
		return 0
	}

	profits := []int{}
	temp := 0
	for i := 1; i < size; i++ {
		diff := prices[i] - prices[i-1]

		if temp*diff >= 0 {
			temp += diff
			continue
		}

		profits = append(profits, temp)
		temp = diff
	}
	profits = append(profits, temp)

	res := 0
	for i := 0; i < len(profits); i++ {
		temp = maxInSlice(profits[:i]) + maxInSlice(profits[i:])
		if res < temp {
			res = temp
		}
	}

	return res
}

func maxInSlice(ps []int) int {
	max, tmp := 0, 0

	for _, p := range ps {
		if tmp < 0 {
			tmp = 0
		}

		tmp += p

		if max < tmp {
			max = tmp
		}
	}

	return max
}
func testMaxProfit121_123() {
	//prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	//prices := []int{1, 2}
	prices := []int{7, 6, 5, 4, 3, 2, 1}
	fmt.Println(maxProfit(prices))
	fmt.Println(maxProfit2(prices))
	fmt.Println(maxProfit3(prices))
}
