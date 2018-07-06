package main

import "fmt"

func reserveNumber(x int) int {
	r := x % 10

	for true {
		x = x/10
		if x == 0 {
			break
		}

		r = r * 10 + x % 10
	}

	return r
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	y := reserveNumber(x)
	if x == y {
		return true
	}

	return false
}

func main()  {
	n := 34
	fmt.Println(n, isPalindrome(n))

	n = -100
	fmt.Println(n, isPalindrome(n))
}