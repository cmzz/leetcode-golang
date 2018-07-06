package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int)  {
	for i := m + n - 1; m > 0 && n > 0 && i >= 0 ; i-- {
		if nums1[m - 1] > nums2[n - 1] {
			nums1[i] = nums1[m - 1]
			m--
		} else {
			nums1[i] = nums2[n - 1]
			n--
		}
	}

	if m == 0 {
		for ; n > 0;  {
			nums1[n - 1] = nums2[n - 1]
			n--
		}
	}
}

func main() {
	var nums1 []int
	var nums2 []int
	var m, n int

	nums1 = []int{1,2,3,0,0,0}
	nums2 = []int{2,5,6}
	m = 3
	n = 3

	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}