package main

import "fmt"

func majorityElement(nums []int) int {
	tmp := make(map[int]int)

	for _, v := range nums {
		if _, ok := tmp[v]; ok {
			tmp[v] = tmp[v] + 1
		} else {
			tmp[v] = 1
		}
	}

	max := -1
	maxk := -1
	for k, v := range tmp {
		if v > len(nums) / 2 {
			if v > max {
				maxk = k
			}
		}
	}

	return maxk
}

func main() {
	fmt.Println(majorityElement([]int{3,2,3}))
	fmt.Println(majorityElement([]int{2,2,1,1,1,2,2}))
}