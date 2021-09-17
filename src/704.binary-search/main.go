package main

func search(nums []int, target int) int {
	nl := len(nums)
	if nl < 1 {
		return -1
	}

	left, right := 0, nl-1
	for left <= right {
		mid := left + (right-left)/2
		n := nums[mid]
		if n == target {
			return mid
		} else if target < n {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
