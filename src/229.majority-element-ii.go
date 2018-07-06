package main

import "fmt"

func majorityElement(nums []int) []int {
	tmp := make(map[int]int)
	tmp2 := []int{}

	for _, v := range nums {
		if _, ok := tmp[v]; ok {
			tmp[v] = tmp[v] + 1
		} else {
			tmp[v] = 1
		}
	}

	for k, v := range tmp {
		if v > len(nums) / 3 {
			tmp2 = append(tmp2, k)
		}
	}

	return tmp2
}

func main() {
	fmt.Println(majorityElement([]int{3,2,3}))
	fmt.Println(majorityElement([]int{1,1,1,3,3,2,2,2}))
}