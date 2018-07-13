package main

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	c := 0

	for _, i := range nums {
		if i == 1 {
			c++
			if c > max {
				max = c
			}
		} else {
			c = 0
		}
	}

	return max
}

func main() {

}