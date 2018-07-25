package main

import "fmt"

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	count := 0

	m := make(map[int]int)
	pairs := make(map[int][]int)

	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i] + k]; ok {
			if v != i {
				count++
				pairs[i] = []int{nums[i], nums[v]}
				delete(m, nums[i] + k)
			}
		}
	}

	fmt.Println(m, pairs)

	return count
}

func main() {
	fmt.Println(findPairs([]int{3, 1, 4, 1, 5}, 2))
}