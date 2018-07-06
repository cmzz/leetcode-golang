package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers) - 1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] + numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}

	return []int{0, 0}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}