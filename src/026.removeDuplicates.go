package main

import "fmt"

func removeDuplicates(nums []int) int {
	i := 0

	if len(nums) < 1 {
		return 0
	}

	for j := 1; j < len(nums) ; j++  {
		if nums[i] != nums[j] {
			i = i +1
			if i != j {
				nums[i] = nums[j]
			}
		}
	}
	
	return i + 1
}

func main()  {
	var arr []int
	arr = []int{1, 1, 2}
	fmt.Println(removeDuplicates(arr), arr)
}