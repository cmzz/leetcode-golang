package main

func swapInt(nums []int, k int, j int)  {
	if nums[k] > nums[j] {
		t := nums[k]
		nums[k] = nums[j]
		nums[j] = t
	}
}

func quickSort(nums []int, left int, right int)  {
	p := nums[left]

	if left < right {

	}
}

func findKthLargest(nums []int, k int) int {
	if k >= 0 && k < len(nums) {
		quickSort(nums)
		return nums[k]
	}

	return 0
}

func main() {

}