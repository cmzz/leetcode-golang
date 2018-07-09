package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (n+1) / 2

	for _, i := range nums {
		sum = sum - i
	}

	return sum
}

func main()  {
	fmt.Println(missingNumber([]int{0,1,2,3,5}))
}