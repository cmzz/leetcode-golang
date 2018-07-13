package main

import (
	"fmt"
	"math"
)

func findDisappearedNumbers(nums []int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		index := int64(math.Abs(float64(nums[i])) - 1)

		nums[index] = -(int(math.Abs(float64(nums[index]))))
	}

	for i, v := range nums{
		if v > 0 {
			result = append(result, i + 1)
		}
	}

	return result
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{1,1,4,2}))
}