package main

import "fmt"

func reverseArr(a []int) {
	for i := 0; i < len(a) / 2 ; i++ {
		t := a[i]
		n := len(a) - 1 - i
		a[i] = a[n]
		a[n] = t
	}
}

func rotate(nums []int, k int)  {
	if k > len(nums) {
		k = k % len(nums)
	}

	var (
		tmp1 []int
		tmp2 []int
	)

	tmp1 = append(nums[len(nums) - k:])
	tmp2 = append(nums[0:len(nums) - k])

	reverseArr(tmp1)
	reverseArr(tmp2)
	nums = append(tmp2, tmp1[0:]...)
	reverseArr(nums)

	fmt.Println(tmp1, tmp2, nums)
}

func main() {
	var a []int
	a = []int{1,2,3,4,5,6,7}
	rotate(a, 3)
	fmt.Println(a)

	a = []int{1,2,3,4,5,6}
	rotate(a, 2)
	fmt.Println(a)

	//a = []int{-1}
	//rotate(a, 2)
	//fmt.Println(a)
	//
	//a = []int{1, 2}
	//rotate(a, 3)
	//fmt.Println(a)
}