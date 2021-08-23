package main

import (
	"fmt"
	"math"
)

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	var (
		re    float64
		isNeg bool
	)

	re = 1
	if x < 0 && n%2 != 0 {
		isNeg = true
	}

	if x == 1 || x == -1 {
		if isNeg {
			return -1
		}
		return 1
	}

	if n > 0 {
		for i := 0; i < n; i++ {
			re *= x
		}
	} else {
		for i := 0; i < -n; i++ {
			re *= x
		}
		re = 1 / re
	}

	if !isNeg && re > math.MaxFloat64 {
		return math.MaxFloat64
	}
	return re
}

func main() {
	fmt.Println(myPow(-13.62608, 3))
}
