package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if len(nums) < 1 {
		return 0
	}

	for i, v := range nums {
		if v == target {
			return i
		} else if v > target {
			return i
		}
	}

	return len(nums)
}

func main() {
	var arr []int

	arr = []int{1,3,5,6}
	fmt.Println(searchInsert(arr, 5))
}