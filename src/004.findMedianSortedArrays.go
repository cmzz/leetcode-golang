package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 中间值
	first := 0
	second := 0

	k := 0

	n1Len := len(nums1)
	n2Len := len(nums2)

	max := (n1Len + n2Len) / 2

	i1 := 0
	i2 := 0

	for {
		if k > max {
			break
		}

		second = first

		if i1 < n1Len && i2 < n2Len {
			if nums1[i1] < nums2[i2] {
				first = nums1[i1]
				i1++
			} else {
				first = nums2[i2]
				i2++
			}
		} else if i1 < n1Len {
			first = nums1[i1]
			i1++
		} else {
			first = nums2[i2]
			i2++
		}

		k++
	}

	if (n1Len + n2Len) % 2 == 0 {
		return float64(first + second) / 2
	} else {
		return float64(first)
	}
}

func main() {
	nums1 := []int{}
	nums2 := []int{1}

	fmt.Printf("%v", nums1)
	fmt.Printf("%v", nums2)

	fmt.Println("res", findMedianSortedArrays(nums1, nums2))
}