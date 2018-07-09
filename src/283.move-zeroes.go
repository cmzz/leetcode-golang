package main

import "fmt"

func swap(nums []int, i int , j int)  {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}

func move(nums []int, s int, p int)  {
	fmt.Println("+",nums, s, p)
	for ; s < len(nums); s++ {
		fmt.Println(p, s, s-p, nums[s], "=>", nums[s-p])
		t := nums[s - p]
		nums[s - p] = nums[s]
		nums[s] = t
	}
	fmt.Println("-", nums)
}

func moveZeroes(nums []int)  {
	p := 0

	for i := 0; i < len(nums);i++ {
		if nums[i] == 0 {
			p++
		}

		if nums[i] != 0 {
			if p > 0 && i > 0 {
				move(nums, i, p)
				i = i - p
				p = 0
			}
		}

		fmt.Println(i, nums)
	}
}

func main() {
	arr := []int{4,2,4,0,0,3,0,5,1,0}

	moveZeroes(arr)
	fmt.Println(arr)
}