package search_insert_position

// searchInsert 如果给定的数，在数组中，那么就能找到，返回中间索引即可
// 如果不在数组中，一定会找不到 mid
// 		最终，left 和 right 应该是相邻的2个位置
// 返回 right 的位置
func searchInsert(nums []int, target int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}

	l, h := 0, n-1
	if target > nums[h] {
		return h + 1 // n
	}

	if target < nums[0] {
		return 0
	}

	for l <= h-1 {
		mid := l + (h-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}

	return h
}
