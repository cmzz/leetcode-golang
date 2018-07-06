package main

import (
	"math"
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	l := 0
	h := make(map[uint8]int)
	left := 0
	newLeft := 0
	
	for i := 0; i < len(s); i++  {
		if _, ok := h[s[i]]; ok {
			newLeft = h[s[i]] + 1
			if newLeft > left{
				left = newLeft
			}
		}

		h[s[i]] = i
		fmt.Println(i, left, l, float64(i - left + 1))
		fmt.Printf("%v\n", h)

		l = int(math.Max(float64(l), float64(i - left + 1)))
	}

	return l
}

func main() {
	var a = "abba"
	fmt.Println("max: ", lengthOfLongestSubstring(a))
}