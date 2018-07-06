package main

import "fmt"

func plusOne(digits []int) []int {
	l := len(digits) - 1

	r := 0

	for l >= 0 {
		if digits[l] < 9 {
			digits[l] = digits[l] + 1
			r = 0
			break
		} else {
			digits[l] = 0
			r = 1
			l = l - 1
		}
	}

	if l < 0 {
		tmp := []int{1}
		for _, v := range digits {
			tmp = append(tmp, v)
		}

		digits = tmp
	} else {
		digits[l] = digits[l] + r
	}

	return digits
}

func main() {
	var arr []int

	arr = []int{9}
	fmt.Println(plusOne(arr))

	arr = []int{8, 9, 9, 9}
	fmt.Println(plusOne(arr))
}