package main

import (
	"fmt"
	"math"
)

func thirdMax(nums []int) int {
	var (
		fn int
		sn int
		tn int
	)
	fn = math.MinInt64
	sn = math.MinInt64
	tn = math.MinInt64

	for i := 0; i < len(nums) ; i++  {
		t := nums[i]

		if t != fn && t != sn && t > tn {
			if t > fn {
				tn = sn
				sn = fn
				fn = t
			} else if t > sn {
				tn = sn
				sn = t
			} else if t > tn {
				tn = t
			}
		}
	}

	if tn != math.MinInt64 {
		return tn
	}

	return fn
}

func main() {

	//fmt.Println(thirdMax([]int{2,3,1}))
	//
	//fmt.Println(thirdMax([]int{2, 2, 3, 1}))
	//fmt.Println(thirdMax([]int{1, 1, 2}))
	fmt.Println(thirdMax([]int{1,-2147483648,2}))
}