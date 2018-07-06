package main

import (
	"fmt"
	"math"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int][]int)

	for i, v := range nums {
		if _, ok := m[v]; ok {
			m[v] = append(m[v], i)
		} else {
			m[v] = []int{i}
		}
	}

	fmt.Println(m)

	for _, v := range m {
		if len(v) > 1 {
			for i := 0; i < len(v) -1 ; i++  {
				for j := i + 1; j < len(v); j++ {
					if math.Abs(float64(v[i] - v[j])) <= float64(k) {
						return true
					}
				}
			}
		}
	}

	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{99, 99}, 2))
}