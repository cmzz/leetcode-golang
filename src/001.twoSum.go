package main

import "fmt"

/**
给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。

示例:

给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9

所以返回 [0, 1]
 */

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j:= i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i,j}
			}			
		}
	}

	return []int{}
}

/**
 * 有问题
 */
func twoSum2(nums []int, target int) []int {
	var numMap = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = i
	}

	for i, v := range numMap {
		if _, ok := numMap[target - i]; ok {
			return []int{v, numMap[target - i]}
		}
	}
	
	return []int{}
}

func main() {
	var ret = twoSum2([]int{11,2,15,7}, 9)
	fmt.Printf("%v\n", ret)
}
