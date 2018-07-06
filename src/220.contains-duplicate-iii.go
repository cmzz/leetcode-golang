package main

import (
	"math"
	"fmt"
)

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	m := make(map[int][]int)

	for i, v := range nums {
		if _, ok := m[v]; ok {
			m[v] = append(m[v], i)
		} else {
			m[v] = []int{i}
		}
	}

	fmt.Println(m)

	for v1, ilist1 := range m {
		for v2, ilist2 := range m {
			if math.Abs(float64(v1 - v2)) <= float64(t) {
				for i1 := 0; i1 < len(ilist1); i1++ {
					for i2 := 0; i2 < len(ilist2); i2++ {
						if ilist1[i1] != ilist2[i2] && math.Abs(float64(ilist1[i1] - ilist2[i2])) <= float64(k) {
							return true
						}
					}
				}
			}
		}
	}
	
	return false
}

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1,2,3,1}, 3, 0))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1,0,1,1}, 1, 2))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1,5,9,1,5,9}, 2, 3))
	fmt.Println(containsNearbyAlmostDuplicate([]int{2, 1}, 1, 1))
}