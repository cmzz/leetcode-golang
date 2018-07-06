package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) < 1 {
		return 0
	}

	i := 0

	for j := 0; j < len(nums) ; j++  {
		if nums[j] != val {
			nums[i] = nums[j]
			i = i + 1
		}
	}

	return i
}

func main()  {
	var arr []int

	arr = []int {0,1,2,2,3,0,4,2}
	fmt.Println(removeElement(arr, 2), arr)
}