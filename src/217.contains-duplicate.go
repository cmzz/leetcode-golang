package main

func containsDuplicate1(nums []int) bool {
	for i := 0 ; i < len(nums) - 1; i++  {
		for j := i + 1; j < len(nums) ; j++  {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

func containsDuplicate(nums []int) bool {
	e := make(map[int]int)

	for _, v := range nums {
		if _, ok := e[v]; ok {
			return true
		} else {
			e[v] = v
		}
	}

	return false
}

func main() {

}